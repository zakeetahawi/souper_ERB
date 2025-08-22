package customers

import (
	"time"

	"elkhawaga-erp/internal/db/repo"
)

// CustomerCreateRequest - طلب إنشاء عميل جديد
type CustomerCreateRequest struct {
	Name                     string   `json:"name" validate:"required"`
	PhonePrimary             string   `json:"phone_primary" validate:"required"`
	PhoneSecondary           *string  `json:"phone_secondary"`
	GovernorateCode          string   `json:"governorate_code" validate:"required"`
	DistrictName             string   `json:"district_name" validate:"required"`
	CustomerTypeID           *uint    `json:"customer_type_id"`
	CustomerClassificationID *uint    `json:"customer_classification_id"`
	BranchID                 uint     `json:"branch_id" validate:"required"`
	Interests                []string `json:"interests"`
}

// CustomerUpdateRequest - طلب تحديث عميل
type CustomerUpdateRequest struct {
	Name                     string   `json:"name" validate:"required"`
	PhonePrimary             string   `json:"phone_primary" validate:"required"`
	PhoneSecondary           *string  `json:"phone_secondary"`
	GovernorateCode          string   `json:"governorate_code" validate:"required"`
	DistrictName             string   `json:"district_name" validate:"required"`
	CustomerTypeID           *uint    `json:"customer_type_id"`
	CustomerClassificationID *uint    `json:"customer_classification_id"`
	BranchID                 uint     `json:"branch_id" validate:"required"`
	Interests                []string `json:"interests"`
}

// CustomerResponse - استجابة بيانات العميل
type CustomerResponse struct {
	ID                     uint                            `json:"id"`
	CustomerCode           string                          `json:"customer_code"`
	Name                   string                          `json:"name"`
	PhonePrimary           string                          `json:"phone_primary"`
	PhoneSecondary         *string                         `json:"phone_secondary"`
	GovernorateCode        string                          `json:"governorate_code"`
	GovernorateName        string                          `json:"governorate_name"`
	DistrictName           string                          `json:"district_name"`
	CustomerType           *CustomerTypeResponse           `json:"customer_type,omitempty"`
	CustomerClassification *CustomerClassificationResponse `json:"customer_classification,omitempty"`
	Branch                 *BranchResponse                 `json:"branch,omitempty"`
	Interests              []string                        `json:"interests"`
	CreatedAt              time.Time                       `json:"created_at"`
	UpdatedAt              time.Time                       `json:"updated_at"`
}

// CustomerTypeResponse - استجابة نوع العميل
type CustomerTypeResponse struct {
	ID     uint   `json:"id"`
	NameAr string `json:"name_ar"`
	NameEn string `json:"name_en"`
}

// CustomerClassificationResponse - استجابة تصنيف العميل
type CustomerClassificationResponse struct {
	ID     uint   `json:"id"`
	NameAr string `json:"name_ar"`
	NameEn string `json:"name_en"`
}

// BranchResponse - استجابة الفرع
type BranchResponse struct {
	ID     uint   `json:"id"`
	Code   string `json:"code"`
	NameAr string `json:"name_ar"`
}

// CustomerListResponse - استجابة قائمة العملاء
type CustomerListResponse struct {
	Customers  []CustomerResponse `json:"customers"`
	Pagination PaginationResponse `json:"pagination"`
}

// PaginationResponse - استجابة التصفح
type PaginationResponse struct {
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// CustomerFilter - فلتر العملاء
type CustomerFilter struct {
	Search         string `json:"search"`
	Governorate    string `json:"governorate"`
	District       string `json:"district"`
	Type           *uint  `json:"type"`
	Classification *uint  `json:"classification"`
	Branch         *uint  `json:"branch"`
	Page           int    `json:"page"`
	Size           int    `json:"size"`
	Sort           string `json:"sort"`
	Order          string `json:"order"`
}

// GovernorateResponse - استجابة المحافظة
type GovernorateResponse struct {
	Code   string `json:"code"`
	NameAr string `json:"name_ar"`
	NameEn string `json:"name_en"`
}

// DistrictResponse - استجابة المنطقة
type DistrictResponse struct {
	Code   string `json:"code"`
	NameAr string `json:"name_ar"`
	NameEn string `json:"name_en"`
}

// ConvertToResponse - تحويل نموذج قاعدة البيانات إلى استجابة
func ConvertToResponse(customer *repo.Customer) *CustomerResponse {
	response := &CustomerResponse{
		ID:              customer.ID,
		CustomerCode:    customer.CustomerCode,
		Name:            customer.Name,
		PhonePrimary:    customer.PhonePrimary,
		PhoneSecondary:  customer.PhoneSecondary,
		GovernorateCode: customer.GovernorateCode,
		DistrictName:    customer.DistrictName,
		Interests:       customer.Interests,
		CreatedAt:       customer.CreatedAt,
		UpdatedAt:       customer.UpdatedAt,
	}

	// إضافة بيانات الفرع
	if customer.Branch != nil {
		response.Branch = &BranchResponse{
			ID:     customer.Branch.ID,
			Code:   customer.Branch.Code,
			NameAr: customer.Branch.NameAr,
		}
	}

	// إضافة بيانات نوع العميل
	if customer.CustomerType != nil {
		response.CustomerType = &CustomerTypeResponse{
			ID:     customer.CustomerType.ID,
			NameAr: customer.CustomerType.NameAr,
			NameEn: customer.CustomerType.NameEn,
		}
	}

	// إضافة بيانات تصنيف العميل
	if customer.CustomerClassification != nil {
		response.CustomerClassification = &CustomerClassificationResponse{
			ID:     customer.CustomerClassification.ID,
			NameAr: customer.CustomerClassification.NameAr,
			NameEn: customer.CustomerClassification.NameEn,
		}
	}

	return response
}

// ConvertToRepoModel - تحويل طلب الإنشاء إلى نموذج قاعدة البيانات
func ConvertToRepoModel(req *CustomerCreateRequest) *repo.Customer {
	return &repo.Customer{
		Name:                     req.Name,
		PhonePrimary:             req.PhonePrimary,
		PhoneSecondary:           req.PhoneSecondary,
		GovernorateCode:          req.GovernorateCode,
		DistrictName:             req.DistrictName,
		CustomerTypeID:           req.CustomerTypeID,
		CustomerClassificationID: req.CustomerClassificationID,
		BranchID:                 req.BranchID,
		Interests:                req.Interests,
	}
}
