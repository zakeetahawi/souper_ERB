package customers

import (
	"strconv"

	"elkhawaga-erp/internal/auth"

	"github.com/gofiber/fiber/v2"
)

type CustomerAPI struct {
	service *CustomerService
	auth    *auth.RBACService
}

func NewCustomerAPI() *CustomerAPI {
	return &CustomerAPI{
		service: NewCustomerService(),
		auth:    auth.NewRBACService(),
	}
}

// RegisterRoutes - تسجيل مسارات API
func (api *CustomerAPI) RegisterRoutes(app *fiber.App) {
	customers := app.Group("/api/customers")

	customers.Get("/", api.ListCustomers)
	customers.Get("/:id", api.GetCustomer)
	customers.Post("/", api.CreateCustomer)
	customers.Put("/:id", api.UpdateCustomer)
	customers.Delete("/:id", api.DeleteCustomer)

	// مسارات مرجعية
	geo := app.Group("/api/geo")
	geo.Get("/governorates", api.GetGovernorates)
	geo.Get("/districts", api.GetDistricts)

	refs := app.Group("/api/refs")
	refs.Get("/types", api.GetCustomerTypes)
	refs.Get("/classifications", api.GetCustomerClassifications)
}

// ListCustomers - جلب قائمة العملاء
func (api *CustomerAPI) ListCustomers(c *fiber.Ctx) error {
	// التحقق من الصلاحيات
	userID := c.Locals("user_id").(uint)
	hasPermission, err := api.auth.CheckPermission(userID, "customers", "read")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to check permissions",
		})
	}
	if !hasPermission {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "insufficient permissions",
		})
	}

	// تحليل المعاملات
	filter := &CustomerFilter{
		Page:  1,
		Size:  20,
		Sort:  "created_at",
		Order: "desc",
	}

	// معاملات التصفح
	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		filter.Page = page
	}
	if size, err := strconv.Atoi(c.Query("size")); err == nil && size > 0 && size <= 100 {
		filter.Size = size
	}

	// معاملات البحث والفلترة
	filter.Search = c.Query("search")
	filter.Governorate = c.Query("governorate")
	filter.District = c.Query("district")
	filter.Sort = c.Query("sort", "created_at")
	filter.Order = c.Query("order", "desc")

	// معاملات الفلترة الرقمية
	if typeID, err := strconv.ParseUint(c.Query("type"), 10, 32); err == nil {
		uintTypeID := uint(typeID)
		filter.Type = &uintTypeID
	}
	if classificationID, err := strconv.ParseUint(c.Query("classification"), 10, 32); err == nil {
		uintClassificationID := uint(classificationID)
		filter.Classification = &uintClassificationID
	}
	if branchID, err := strconv.ParseUint(c.Query("branch"), 10, 32); err == nil {
		uintBranchID := uint(branchID)
		filter.Branch = &uintBranchID
	}

	// جلب البيانات
	result, err := api.service.ListCustomers(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

// GetCustomer - جلب عميل محدد
func (api *CustomerAPI) GetCustomer(c *fiber.Ctx) error {
	// التحقق من الصلاحيات
	userID := c.Locals("user_id").(uint)
	hasPermission, err := api.auth.CheckPermission(userID, "customers", "read")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to check permissions",
		})
	}
	if !hasPermission {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "insufficient permissions",
		})
	}

	// تحليل معرف العميل
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid customer ID",
		})
	}

	// جلب العميل
	customer, err := api.service.GetCustomerByID(uint(id))
	if err != nil {
		if err.Error() == "customer not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"error":   "customer not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

// CreateCustomer - إنشاء عميل جديد
func (api *CustomerAPI) CreateCustomer(c *fiber.Ctx) error {
	// التحقق من الصلاحيات
	userID := c.Locals("user_id").(uint)
	hasPermission, err := api.auth.CheckPermission(userID, "customers", "create")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to check permissions",
		})
	}
	if !hasPermission {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "insufficient permissions",
		})
	}

	// تحليل البيانات
	var req CustomerCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid request body",
		})
	}

	// إنشاء العميل
	customer, err := api.service.CreateCustomer(&req)
	if err != nil {
		if err.Error() == "phone number already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"success": false,
				"error":   "phone number already exists",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":            customer.ID,
			"customer_code": customer.CustomerCode,
			"message":       "تم إنشاء العميل بنجاح",
		},
	})
}

// UpdateCustomer - تحديث عميل
func (api *CustomerAPI) UpdateCustomer(c *fiber.Ctx) error {
	// التحقق من الصلاحيات
	userID := c.Locals("user_id").(uint)
	hasPermission, err := api.auth.CheckPermission(userID, "customers", "update")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to check permissions",
		})
	}
	if !hasPermission {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "insufficient permissions",
		})
	}

	// تحليل معرف العميل
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid customer ID",
		})
	}

	// تحليل البيانات
	var req CustomerUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid request body",
		})
	}

	// تحديث العميل
	_, err = api.service.UpdateCustomer(uint(id), &req)
	if err != nil {
		if err.Error() == "customer not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"error":   "customer not found",
			})
		}
		if err.Error() == "phone number already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"success": false,
				"error":   "phone number already exists",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"message": "تم تحديث العميل بنجاح",
		},
	})
}

// DeleteCustomer - حذف عميل
func (api *CustomerAPI) DeleteCustomer(c *fiber.Ctx) error {
	// التحقق من الصلاحيات
	userID := c.Locals("user_id").(uint)
	hasPermission, err := api.auth.CheckPermission(userID, "customers", "delete")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to check permissions",
		})
	}
	if !hasPermission {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "insufficient permissions",
		})
	}

	// تحليل معرف العميل
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid customer ID",
		})
	}

	// حذف العميل
	err = api.service.DeleteCustomer(uint(id))
	if err != nil {
		if err.Error() == "customer not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"error":   "customer not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"message": "تم حذف العميل بنجاح",
		},
	})
}

// GetGovernorates - جلب قائمة المحافظات
func (api *CustomerAPI) GetGovernorates(c *fiber.Ctx) error {
	governorates, err := api.service.GetGovernorates()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    governorates,
	})
}

// GetDistricts - جلب قائمة المناطق
func (api *CustomerAPI) GetDistricts(c *fiber.Ctx) error {
	governorateCode := c.Query("g")
	if governorateCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "governorate code is required",
		})
	}

	districts, err := api.service.GetDistricts(governorateCode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    districts,
	})
}

// GetCustomerTypes - جلب أنواع العملاء
func (api *CustomerAPI) GetCustomerTypes(c *fiber.Ctx) error {
	types, err := api.service.GetCustomerTypes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    types,
	})
}

// GetCustomerClassifications - جلب تصنيفات العملاء
func (api *CustomerAPI) GetCustomerClassifications(c *fiber.Ctx) error {
	classifications, err := api.service.GetCustomerClassifications()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    classifications,
	})
}
