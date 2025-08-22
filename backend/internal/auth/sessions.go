package auth

import (
	"errors"
	"time"

	"elkhawaga-erp/internal/config"
	"elkhawaga-erp/internal/db"
	"elkhawaga-erp/internal/db/repo"

	"gorm.io/gorm"
)

type SessionService struct {
	db *gorm.DB
}

func NewSessionService() *SessionService {
	return &SessionService{
		db: db.GetDB(),
	}
}

// CreateSession - إنشاء جلسة جديدة للمستخدم
func (s *SessionService) CreateSession(userID uint, deviceFingerprint, refreshToken, ip, userAgent string) error {
	// إذا كان Single Device Login مفعل، نلغي الجلسات السابقة
	if config.Get().Features.SingleDeviceLogin {
		if err := s.RevokeUserSessions(userID); err != nil {
			return err
		}
	}

	session := &repo.UserSession{
		UserID:            userID,
		DeviceFingerprint: deviceFingerprint,
		RefreshToken:      refreshToken,
		IP:                ip,
		UserAgent:         userAgent,
		IsActive:          true,
	}

	return s.db.Create(session).Error
}

// ValidateSession - التحقق من صحة الجلسة
func (s *SessionService) ValidateSession(userID uint, refreshToken string) (*repo.UserSession, error) {
	var session repo.UserSession

	err := s.db.Where("user_id = ? AND refresh_token = ? AND is_active = ?",
		userID, refreshToken, true).First(&session).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found or inactive")
		}
		return nil, err
	}

	// التحقق من انتهاء صلاحية الجلسة
	if session.RevokedAt != nil {
		return nil, errors.New("session has been revoked")
	}

	return &session, nil
}

// RevokeSession - إلغاء جلسة محددة
func (s *SessionService) RevokeSession(sessionID uint) error {
	now := time.Now()
	return s.db.Model(&repo.UserSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"is_active":  false,
			"revoked_at": &now,
		}).Error
}

// RevokeUserSessions - إلغاء جميع جلسات المستخدم
func (s *SessionService) RevokeUserSessions(userID uint) error {
	now := time.Now()
	return s.db.Model(&repo.UserSession{}).
		Where("user_id = ? AND is_active = ?", userID, true).
		Updates(map[string]interface{}{
			"is_active":  false,
			"revoked_at": &now,
		}).Error
}

// RevokeSessionByDevice - إلغاء جلسة حسب معرف الجهاز
func (s *SessionService) RevokeSessionByDevice(userID uint, deviceFingerprint string) error {
	now := time.Now()
	return s.db.Model(&repo.UserSession{}).
		Where("user_id = ? AND device_fingerprint = ? AND is_active = ?",
			userID, deviceFingerprint, true).
		Updates(map[string]interface{}{
			"is_active":  false,
			"revoked_at": &now,
		}).Error
}

// GetUserActiveSessions - جلب الجلسات النشطة للمستخدم
func (s *SessionService) GetUserActiveSessions(userID uint) ([]repo.UserSession, error) {
	var sessions []repo.UserSession
	err := s.db.Where("user_id = ? AND is_active = ?", userID, true).Find(&sessions).Error
	return sessions, err
}

// CleanupExpiredSessions - تنظيف الجلسات المنتهية الصلاحية
func (s *SessionService) CleanupExpiredSessions() error {
	cfg := config.Get()
	expiryTime := time.Now().Add(-cfg.JWT.RefreshExpires)

	return s.db.Model(&repo.UserSession{}).
		Where("created_at < ? AND is_active = ?", expiryTime, true).
		Updates(map[string]interface{}{
			"is_active":  false,
			"revoked_at": time.Now(),
		}).Error
}
