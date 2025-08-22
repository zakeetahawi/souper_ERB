package customers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"elkhawaga-erp/internal/db"
	"elkhawaga-erp/internal/db/repo"

	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		db: db.GetDB(),
	}
}

// CreateCustomer - إنشاء عميل جديد
func (s *CustomerService) CreateCustomer(req *CustomerCreateRequest) (*CustomerResponse, error) {
	// التحقق من صحة البيانات
	if err := s.validateCustomerData(req); err != nil {
		return nil, err
	}

	// التحقق من عدم تكرار رقم الهاتف
	if err := s.checkPhoneUniqueness(req.PhonePrimary, 0); err != nil {
		return nil, err
	}

	// إنشاء كود العميل
	customerCode, err := s.generateCustomerCode(req.BranchID)
	if err != nil {
		return nil, err
	}

	// إنشاء العميل
	customer := ConvertToRepoModel(req)
	customer.CustomerCode = customerCode

	if err := s.db.Create(customer).Error; err != nil {
		return nil, err
	}

	// جلب العميل مع البيانات المرتبطة
	return s.GetCustomerByID(customer.ID)
}

// GetCustomerByID - جلب عميل حسب المعرف
func (s *CustomerService) GetCustomerByID(id uint) (*CustomerResponse, error) {
	var customer repo.Customer

	err := s.db.Preload("Branch").
		Preload("CustomerType").
		Preload("CustomerClassification").
		Where("id = ?", id).
		First(&customer).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	response := ConvertToResponse(&customer)

	// إضافة اسم المحافظة
	if governorate, err := s.getGovernorateName(customer.GovernorateCode); err == nil {
		response.GovernorateName = governorate
	}

	return response, nil
}

// UpdateCustomer - تحديث بيانات العميل
func (s *CustomerService) UpdateCustomer(id uint, req *CustomerUpdateRequest) (*CustomerResponse, error) {
	// التحقق من وجود العميل
	var existingCustomer repo.Customer
	if err := s.db.Where("id = ?", id).First(&existingCustomer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	// التحقق من صحة البيانات
	if err := s.validateCustomerData(&CustomerCreateRequest{
		Name:                     req.Name,
		PhonePrimary:             req.PhonePrimary,
		PhoneSecondary:           req.PhoneSecondary,
		GovernorateCode:          req.GovernorateCode,
		DistrictName:             req.DistrictName,
		CustomerTypeID:           req.CustomerTypeID,
		CustomerClassificationID: req.CustomerClassificationID,
		BranchID:                 req.BranchID,
		Interests:                req.Interests,
	}); err != nil {
		return nil, err
	}

	// التحقق من عدم تكرار رقم الهاتف (باستثناء العميل الحالي)
	if err := s.checkPhoneUniqueness(req.PhonePrimary, id); err != nil {
		return nil, err
	}

	// تحديث البيانات
	updates := map[string]interface{}{
		"name":                       req.Name,
		"phone_primary":              req.PhonePrimary,
		"phone_secondary":            req.PhoneSecondary,
		"governorate_code":           req.GovernorateCode,
		"district_name":              req.DistrictName,
		"customer_type_id":           req.CustomerTypeID,
		"customer_classification_id": req.CustomerClassificationID,
		"branch_id":                  req.BranchID,
		"interests":                  req.Interests,
	}

	if err := s.db.Model(&existingCustomer).Updates(updates).Error; err != nil {
		return nil, err
	}

	return s.GetCustomerByID(id)
}

// DeleteCustomer - حذف عميل (حذف ناعم)
func (s *CustomerService) DeleteCustomer(id uint) error {
	var customer repo.Customer
	if err := s.db.Where("id = ?", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer not found")
		}
		return err
	}

	return s.db.Delete(&customer).Error
}

// ListCustomers - جلب قائمة العملاء
func (s *CustomerService) ListCustomers(filter *CustomerFilter) (*CustomerListResponse, error) {
	var customers []repo.Customer
	var total int64

	query := s.db.Model(&repo.Customer{}).
		Preload("Branch").
		Preload("CustomerType").
		Preload("CustomerClassification")

	// تطبيق الفلاتر
	query = s.applyFilters(query, filter)

	// حساب العدد الإجمالي
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// تطبيق الترتيب والتصفح
	query = s.applySortingAndPagination(query, filter)

	// جلب البيانات
	if err := query.Find(&customers).Error; err != nil {
		return nil, err
	}

	// تحويل إلى استجابة
	responses := make([]CustomerResponse, len(customers))
	for i, customer := range customers {
		response := ConvertToResponse(&customer)

		// إضافة اسم المحافظة
		if governorate, err := s.getGovernorateName(customer.GovernorateCode); err == nil {
			response.GovernorateName = governorate
		}

		responses[i] = *response
	}

	// حساب معلومات التصفح
	totalPages := int((total + int64(filter.Size) - 1) / int64(filter.Size))

	return &CustomerListResponse{
		Customers: responses,
		Pagination: PaginationResponse{
			Page:       filter.Page,
			Size:       filter.Size,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

// GetGovernorates - جلب قائمة المحافظات
func (s *CustomerService) GetGovernorates() ([]GovernorateResponse, error) {
	var governorates []repo.Governorate
	if err := s.db.Find(&governorates).Error; err != nil {
		return nil, err
	}

	responses := make([]GovernorateResponse, len(governorates))
	for i, gov := range governorates {
		responses[i] = GovernorateResponse{
			Code:   gov.Code,
			NameAr: gov.NameAr,
			NameEn: gov.NameEn,
		}
	}

	return responses, nil
}

// GetDistricts - جلب قائمة المناطق حسب المحافظة
func (s *CustomerService) GetDistricts(governorateCode string) ([]DistrictResponse, error) {
	var governorate repo.Governorate
	if err := s.db.Where("code = ?", governorateCode).First(&governorate).Error; err != nil {
		return nil, err
	}

	responses := make([]DistrictResponse, len(governorate.Districts))
	for i, district := range governorate.Districts {
		responses[i] = DistrictResponse{
			Code:   strings.ToUpper(district),
			NameAr: district,
			NameEn: district, // يمكن تحسين هذا لاحقاً
		}
	}

	return responses, nil
}

// GetCustomerTypes - جلب أنواع العملاء
func (s *CustomerService) GetCustomerTypes() ([]CustomerTypeResponse, error) {
	var types []repo.RefCustomerType
	if err := s.db.Find(&types).Error; err != nil {
		return nil, err
	}

	responses := make([]CustomerTypeResponse, len(types))
	for i, t := range types {
		responses[i] = CustomerTypeResponse{
			ID:     t.ID,
			NameAr: t.NameAr,
			NameEn: t.NameEn,
		}
	}

	return responses, nil
}

// GetCustomerClassifications - جلب تصنيفات العملاء
func (s *CustomerService) GetCustomerClassifications() ([]CustomerClassificationResponse, error) {
	var classifications []repo.RefCustomerClassification
	if err := s.db.Find(&classifications).Error; err != nil {
		return nil, err
	}

	responses := make([]CustomerClassificationResponse, len(classifications))
	for i, c := range classifications {
		responses[i] = CustomerClassificationResponse{
			ID:     c.ID,
			NameAr: c.NameAr,
			NameEn: c.NameEn,
		}
	}

	return responses, nil
}

// validateCustomerData - التحقق من صحة بيانات العميل
func (s *CustomerService) validateCustomerData(req *CustomerCreateRequest) error {
	// التحقق من صحة رقم الهاتف
	if err := s.validatePhoneNumber(req.PhonePrimary); err != nil {
		return err
	}

	if req.PhoneSecondary != nil {
		if err := s.validatePhoneNumber(*req.PhoneSecondary); err != nil {
			return err
		}
	}

	// التحقق من وجود المحافظة
	var governorate repo.Governorate
	if err := s.db.Where("code = ?", req.GovernorateCode).First(&governorate).Error; err != nil {
		return errors.New("invalid governorate")
	}

	// التحقق من وجود المنطقة في المحافظة
	districtExists := false
	for _, district := range governorate.Districts {
		if district == req.DistrictName {
			districtExists = true
			break
		}
	}
	if !districtExists {
		return errors.New("invalid district for the selected governorate")
	}

	// التحقق من وجود الفرع
	var branch repo.Branch
	if err := s.db.Where("id = ?", req.BranchID).First(&branch).Error; err != nil {
		return errors.New("invalid branch")
	}

	// التحقق من وجود نوع العميل إذا تم تحديده
	if req.CustomerTypeID != nil {
		var customerType repo.RefCustomerType
		if err := s.db.Where("id = ?", *req.CustomerTypeID).First(&customerType).Error; err != nil {
			return errors.New("invalid customer type")
		}
	}

	// التحقق من وجود تصنيف العميل إذا تم تحديده
	if req.CustomerClassificationID != nil {
		var classification repo.RefCustomerClassification
		if err := s.db.Where("id = ?", *req.CustomerClassificationID).First(&classification).Error; err != nil {
			return errors.New("invalid customer classification")
		}
	}

	return nil
}

// validatePhoneNumber - التحقق من صحة رقم الهاتف
func (s *CustomerService) validatePhoneNumber(phone string) error {
	// تنظيف الرقم
	phone = strings.TrimSpace(phone)

	// إزالة المسافات والشرطات
	phone = regexp.MustCompile(`[\s\-\(\)]`).ReplaceAllString(phone, "")

	// التحقق من الصيغة المصرية
	egyptPattern := regexp.MustCompile(`^(\+20|20|0)?1[0125][0-9]{8}$`)
	if !egyptPattern.MatchString(phone) {
		return errors.New("invalid Egyptian phone number format")
	}

	return nil
}

// checkPhoneUniqueness - التحقق من عدم تكرار رقم الهاتف
func (s *CustomerService) checkPhoneUniqueness(phone string, excludeID uint) error {
	var count int64
	query := s.db.Model(&repo.Customer{}).Where("phone_primary = ?", phone)

	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	if err := query.Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("phone number already exists")
	}

	return nil
}

// generateCustomerCode - إنشاء كود العميل
func (s *CustomerService) generateCustomerCode(branchID uint) (string, error) {
	// جلب كود الفرع
	var branch repo.Branch
	if err := s.db.Where("id = ?", branchID).First(&branch).Error; err != nil {
		return "", err
	}

	// جلب آخر رقم تسلسلي للفرع
	var lastCustomer repo.Customer
	err := s.db.Where("branch_id = ?", branchID).
		Order("customer_code DESC").
		First(&lastCustomer).Error

	var nextNumber int
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			nextNumber = 1
		} else {
			return "", err
		}
	} else {
		// استخراج الرقم من الكود السابق
		parts := strings.Split(lastCustomer.CustomerCode, "-")
		if len(parts) == 2 {
			if num, err := strconv.Atoi(parts[1]); err == nil {
				nextNumber = num + 1
			} else {
				nextNumber = 1
			}
		} else {
			nextNumber = 1
		}
	}

	return fmt.Sprintf("%s-%04d", branch.Code, nextNumber), nil
}

// applyFilters - تطبيق الفلاتر على الاستعلام
func (s *CustomerService) applyFilters(query *gorm.DB, filter *CustomerFilter) *gorm.DB {
	if filter.Search != "" {
		searchTerm := "%" + filter.Search + "%"
		query = query.Where("name ILIKE ? OR phone_primary ILIKE ? OR customer_code ILIKE ?",
			searchTerm, searchTerm, searchTerm)
	}

	if filter.Governorate != "" {
		query = query.Where("governorate_code = ?", filter.Governorate)
	}

	if filter.District != "" {
		query = query.Where("district_name = ?", filter.District)
	}

	if filter.Type != nil {
		query = query.Where("customer_type_id = ?", *filter.Type)
	}

	if filter.Classification != nil {
		query = query.Where("customer_classification_id = ?", *filter.Classification)
	}

	if filter.Branch != nil {
		query = query.Where("branch_id = ?", *filter.Branch)
	}

	return query
}

// applySortingAndPagination - تطبيق الترتيب والتصفح
func (s *CustomerService) applySortingAndPagination(query *gorm.DB, filter *CustomerFilter) *gorm.DB {
	// الترتيب
	sortField := filter.Sort
	if sortField == "" {
		sortField = "created_at"
	}

	order := filter.Order
	if order == "" {
		order = "desc"
	}

	query = query.Order(fmt.Sprintf("%s %s", sortField, order))

	// التصفح
	offset := (filter.Page - 1) * filter.Size
	query = query.Offset(offset).Limit(filter.Size)

	return query
}

// getGovernorateName - جلب اسم المحافظة
func (s *CustomerService) getGovernorateName(code string) (string, error) {
	var governorate repo.Governorate
	err := s.db.Where("code = ?", code).First(&governorate).Error
	if err != nil {
		return "", err
	}
	return governorate.NameAr, nil
}
