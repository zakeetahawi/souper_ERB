package auth

import (
	"errors"

	"elkhawaga-erp/internal/db"
	"elkhawaga-erp/internal/db/repo"

	"gorm.io/gorm"
)

type RBACService struct {
	db *gorm.DB
}

func NewRBACService() *RBACService {
	return &RBACService{
		db: db.GetDB(),
	}
}

// CheckPermission - التحقق من صلاحية المستخدم
func (r *RBACService) CheckPermission(userID uint, moduleKey, action string) (bool, error) {
	var count int64

	err := r.db.Table("users").
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN role_permissions ON user_roles.role_id = role_permissions.role_id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("users.id = ? AND permissions.module_key = ? AND permissions.action = ?",
			userID, moduleKey, action).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetUserPermissions - جلب جميع صلاحيات المستخدم
func (r *RBACService) GetUserPermissions(userID uint) ([]repo.Permission, error) {
	var permissions []repo.Permission

	err := r.db.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&permissions).Error

	return permissions, err
}

// GetUserRoles - جلب أدوار المستخدم
func (r *RBACService) GetUserRoles(userID uint) ([]repo.Role, error) {
	var user repo.User
	err := r.db.Preload("Roles").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user.Roles, nil
}

// AssignRoleToUser - تعيين دور للمستخدم
func (r *RBACService) AssignRoleToUser(userID, roleID uint) error {
	userRole := repo.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	return r.db.Create(&userRole).Error
}

// RemoveRoleFromUser - إزالة دور من المستخدم
func (r *RBACService) RemoveRoleFromUser(userID, roleID uint) error {
	return r.db.Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&repo.UserRole{}).Error
}

// CreateRole - إنشاء دور جديد
func (r *RBACService) CreateRole(name string) (*repo.Role, error) {
	role := &repo.Role{
		Name: name,
	}

	err := r.db.Create(role).Error
	return role, err
}

// CreatePermission - إنشاء صلاحية جديدة
func (r *RBACService) CreatePermission(moduleKey, action string) (*repo.Permission, error) {
	permission := &repo.Permission{
		ModuleKey: moduleKey,
		Action:    action,
	}

	err := r.db.Create(permission).Error
	return permission, err
}

// AssignPermissionToRole - تعيين صلاحية للدور
func (r *RBACService) AssignPermissionToRole(roleID, permissionID uint) error {
	rolePermission := repo.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
	}

	return r.db.Create(&rolePermission).Error
}

// RemovePermissionFromRole - إزالة صلاحية من الدور
func (r *RBACService) RemovePermissionFromRole(roleID, permissionID uint) error {
	return r.db.Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Delete(&repo.RolePermission{}).Error
}

// GetRolePermissions - جلب صلاحيات الدور
func (r *RBACService) GetRolePermissions(roleID uint) ([]repo.Permission, error) {
	var role repo.Role
	err := r.db.Preload("Permissions").Where("id = ?", roleID).First(&role).Error
	if err != nil {
		return nil, err
	}
	return role.Permissions, nil
}

// RegisterModulePermissions - تسجيل صلاحيات الموديول
func (r *RBACService) RegisterModulePermissions(moduleKey string, actions []string) error {
	for _, action := range actions {
		// التحقق من وجود الصلاحية
		var existingPermission repo.Permission
		err := r.db.Where("module_key = ? AND action = ?", moduleKey, action).
			First(&existingPermission).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// إنشاء الصلاحية الجديدة
				permission := &repo.Permission{
					ModuleKey: moduleKey,
					Action:    action,
				}
				if err := r.db.Create(permission).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}

// GetUserModules - جلب الموديولات المتاحة للمستخدم
func (r *RBACService) GetUserModules(userID uint) ([]string, error) {
	var modules []string

	err := r.db.Table("permissions").
		Select("DISTINCT permissions.module_key").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Pluck("permissions.module_key", &modules).Error

	return modules, err
}
