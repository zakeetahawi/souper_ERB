package repo

import (
	"time"

	"gorm.io/gorm"
)

// Branch - الفروع
type Branch struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Code      string         `json:"code" gorm:"uniqueIndex;not null;size:8"`
	NameAr    string         `json:"name_ar" gorm:"not null"`
	NameEn    string         `json:"name_en"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// User - المستخدمون
type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username" gorm:"uniqueIndex;not null;size:64"`
	PasswordHash string         `json:"-" gorm:"not null"`
	BranchID     *uint          `json:"branch_id"`
	Branch       *Branch        `json:"branch,omitempty"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// العلاقات
	Roles []Role `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

// UserSession - جلسات المستخدمين
type UserSession struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	UserID            uint           `json:"user_id" gorm:"not null"`
	User              User           `json:"user,omitempty"`
	DeviceFingerprint string         `json:"device_fingerprint" gorm:"not null"`
	RefreshToken      string         `json:"refresh_token" gorm:"uniqueIndex;not null"`
	IP                string         `json:"ip"`
	UserAgent         string         `json:"user_agent"`
	IsActive          bool           `json:"is_active" gorm:"default:true"`
	RevokedAt         *time.Time     `json:"revoked_at"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// Role - الأدوار
type Role struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"uniqueIndex;not null;size:64"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// العلاقات
	Users       []User       `json:"users,omitempty" gorm:"many2many:user_roles;"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
}

// Permission - الصلاحيات
type Permission struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ModuleKey string         `json:"module_key" gorm:"not null;size:64"`
	Action    string         `json:"action" gorm:"not null;size:64"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// العلاقات
	Roles []Role `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
}

// UserRole - علاقة المستخدمين بالأدوار
type UserRole struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	RoleID uint `json:"role_id" gorm:"primaryKey"`
}

// RolePermission - علاقة الأدوار بالصلاحيات
type RolePermission struct {
	RoleID       uint `json:"role_id" gorm:"primaryKey"`
	PermissionID uint `json:"permission_id" gorm:"primaryKey"`
}

// FeatureFlag - أعلام الميزات
type FeatureFlag struct {
	ModuleKey string `json:"module_key" gorm:"primaryKey;size:64"`
	Enabled   bool   `json:"enabled" gorm:"not null;default:true"`
}

// RefCustomerType - أنواع العملاء المرجعية
type RefCustomerType struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	NameAr string `json:"name_ar" gorm:"uniqueIndex;not null"`
	NameEn string `json:"name_en"`
}

// RefCustomerClassification - تصنيفات العملاء المرجعية
type RefCustomerClassification struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	NameAr string `json:"name_ar" gorm:"uniqueIndex;not null"`
	NameEn string `json:"name_en"`
}

// Customer - العملاء
type Customer struct {
	ID                       uint                       `json:"id" gorm:"primaryKey"`
	CustomerCode             string                     `json:"customer_code" gorm:"uniqueIndex;not null;size:32"`
	Name                     string                     `json:"name" gorm:"not null"`
	PhonePrimary             string                     `json:"phone_primary" gorm:"uniqueIndex;not null;size:20"`
	PhoneSecondary           *string                    `json:"phone_secondary" gorm:"size:20"`
	GovernorateCode          string                     `json:"governorate_code" gorm:"not null;size:16"`
	DistrictName             string                     `json:"district_name" gorm:"not null"`
	CustomerTypeID           *uint                      `json:"customer_type_id"`
	CustomerType             *RefCustomerType           `json:"customer_type,omitempty"`
	CustomerClassificationID *uint                      `json:"customer_classification_id"`
	CustomerClassification   *RefCustomerClassification `json:"customer_classification,omitempty"`
	BranchID                 uint                       `json:"branch_id" gorm:"not null"`
	Branch                   *Branch                    `json:"branch,omitempty"`
	Interests                []string                   `json:"interests" gorm:"type:text[];default:'{}'"`
	CreatedAt                time.Time                  `json:"created_at"`
	UpdatedAt                time.Time                  `json:"updated_at"`
	DeletedAt                gorm.DeletedAt             `json:"deleted_at,omitempty" gorm:"index"`
}

// Governorate - المحافظات
type Governorate struct {
	Code      string   `json:"code" gorm:"primaryKey;size:16"`
	NameAr    string   `json:"name_ar" gorm:"not null"`
	NameEn    string   `json:"name_en"`
	Districts []string `json:"districts" gorm:"type:text[]"`
}

// TableName - تحديد اسم الجدول للمحافظات
func (Governorate) TableName() string {
	return "governorates"
}
