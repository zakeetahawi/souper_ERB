package modules

import (
	"elkhawaga-erp/internal/auth"
	"elkhawaga-erp/internal/db"
	"elkhawaga-erp/internal/db/repo"

	"gorm.io/gorm"
)

type ModuleRegistry struct {
	db          *gorm.DB
	rbacService *auth.RBACService
	modules     map[string]*Module
}

type Module struct {
	Key         string
	NameAr      string
	NameEn      string
	Description string
	Enabled     bool
	Permissions []string
}

func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		db:          db.GetDB(),
		rbacService: auth.NewRBACService(),
		modules:     make(map[string]*Module),
	}
}

// RegisterModule - تسجيل موديول جديد
func (r *ModuleRegistry) RegisterModule(module *Module) error {
	r.modules[module.Key] = module

	// تسجيل الصلاحيات في قاعدة البيانات
	if err := r.rbacService.RegisterModulePermissions(module.Key, module.Permissions); err != nil {
		return err
	}

	// إنشاء أو تحديث Feature Flag
	featureFlag := &repo.FeatureFlag{
		ModuleKey: module.Key,
		Enabled:   module.Enabled,
	}

	var existingFlag repo.FeatureFlag
	err := r.db.Where("module_key = ?", module.Key).First(&existingFlag).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// إنشاء Feature Flag جديد
			return r.db.Create(featureFlag).Error
		}
		return err
	}

	// تحديث Feature Flag الموجود
	return r.db.Model(&existingFlag).Update("enabled", module.Enabled).Error
}

// GetModule - جلب موديول محدد
func (r *ModuleRegistry) GetModule(key string) (*Module, bool) {
	module, exists := r.modules[key]
	return module, exists
}

// GetAllModules - جلب جميع الموديولات
func (r *ModuleRegistry) GetAllModules() map[string]*Module {
	return r.modules
}

// IsModuleEnabled - التحقق من تفعيل الموديول
func (r *ModuleRegistry) IsModuleEnabled(key string) bool {
	var featureFlag repo.FeatureFlag
	err := r.db.Where("module_key = ?", key).First(&featureFlag).Error
	if err != nil {
		return false
	}
	return featureFlag.Enabled
}

// EnableModule - تفعيل موديول
func (r *ModuleRegistry) EnableModule(key string) error {
	return r.db.Model(&repo.FeatureFlag{}).
		Where("module_key = ?", key).
		Update("enabled", true).Error
}

// DisableModule - تعطيل موديول
func (r *ModuleRegistry) DisableModule(key string) error {
	return r.db.Model(&repo.FeatureFlag{}).
		Where("module_key = ?", key).
		Update("enabled", false).Error
}

// GetUserEnabledModules - جلب الموديولات المفعلة للمستخدم
func (r *ModuleRegistry) GetUserEnabledModules(userID uint) ([]string, error) {
	var enabledModules []string

	// جلب الموديولات التي لديه صلاحيات عليها
	userModules, err := r.rbacService.GetUserModules(userID)
	if err != nil {
		return nil, err
	}

	// فلترة الموديولات المفعلة فقط
	for _, moduleKey := range userModules {
		if r.IsModuleEnabled(moduleKey) {
			enabledModules = append(enabledModules, moduleKey)
		}
	}

	return enabledModules, nil
}

// InitializeDefaultModules - تهيئة الموديولات الافتراضية
func (r *ModuleRegistry) InitializeDefaultModules() error {
	defaultModules := []*Module{
		{
			Key:         "customers",
			NameAr:      "العملاء",
			NameEn:      "Customers",
			Description: "إدارة العملاء والبيانات الأساسية",
			Enabled:     true,
			Permissions: []string{"read", "create", "update", "delete", "export", "import"},
		},
		{
			Key:         "orders",
			NameAr:      "الطلبات",
			NameEn:      "Orders",
			Description: "إدارة الطلبات والمبيعات",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "approve", "cancel"},
		},
		{
			Key:         "inventory",
			NameAr:      "المخزون",
			NameEn:      "Inventory",
			Description: "إدارة المخزون والمنتجات",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "adjust", "transfer"},
		},
		{
			Key:         "field-survey",
			NameAr:      "المسح الميداني",
			NameEn:      "Field Survey",
			Description: "إدارة المسح الميداني والاستطلاعات",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "assign", "complete"},
		},
		{
			Key:         "factory",
			NameAr:      "المصنع",
			NameEn:      "Factory",
			Description: "إدارة العمليات الإنتاجية",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "schedule", "monitor"},
		},
		{
			Key:         "installations",
			NameAr:      "التركيبات",
			NameEn:      "Installations",
			Description: "إدارة التركيبات والخدمات",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "schedule", "complete"},
		},
		{
			Key:         "data-sync",
			NameAr:      "مزامنة البيانات",
			NameEn:      "Data Sync",
			Description: "مزامنة البيانات مع الأنظمة الخارجية",
			Enabled:     false,
			Permissions: []string{"read", "sync", "configure"},
		},
		{
			Key:         "maintenance",
			NameAr:      "الصيانة",
			NameEn:      "Maintenance",
			Description: "إدارة الصيانة والخدمات",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "schedule", "complete"},
		},
		{
			Key:         "development",
			NameAr:      "التطوير",
			NameEn:      "Development",
			Description: "إدارة مشاريع التطوير",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "assign", "track"},
		},
		{
			Key:         "display-tuning",
			NameAr:      "ضبط العرض",
			NameEn:      "Display Tuning",
			Description: "ضبط وإعدادات العرض",
			Enabled:     false,
			Permissions: []string{"read", "configure", "optimize"},
		},
		{
			Key:         "sales",
			NameAr:      "المبيعات",
			NameEn:      "Sales",
			Description: "إدارة المبيعات والتسويق",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "analyze", "report"},
		},
		{
			Key:         "marketing",
			NameAr:      "التسويق",
			NameEn:      "Marketing",
			Description: "إدارة الحملات التسويقية",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "launch", "track"},
		},
		{
			Key:         "projects",
			NameAr:      "المشاريع",
			NameEn:      "Projects",
			Description: "إدارة المشاريع والمهام",
			Enabled:     false,
			Permissions: []string{"read", "create", "update", "delete", "assign", "track"},
		},
	}

	for _, module := range defaultModules {
		if err := r.RegisterModule(module); err != nil {
			return err
		}
	}

	return nil
}
