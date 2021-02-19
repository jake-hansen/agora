package domain

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type OAuthInfo struct {
	gorm.Model
	UserID            uint
	MeetingPlatformID uint
	AccessToken       string
	RefreshToken      string
	TokenType         string
	Expiry            time.Time
}

type OAuthInfoRepository interface {
	Create(oauthToken *OAuthInfo) (uint, error)
	GetAll() ([]*OAuthInfo, error)
	GetByID(ID uint) (*OAuthInfo, error)
	GetAllByMeetingProviderId(providerID uint) ([]*OAuthInfo, error)
	GetAllByUserID(userID uint) ([]*OAuthInfo, error)
	GetByUserIDAndMeetingPlatformID(userID uint, meetingPlatformID uint) (*OAuthInfo, error)
	Update(oauthToken *OAuthInfo) error
	Delete(ID uint) error
}

type OAuthInfoService interface {
	CreateOAuthInfo(ctx context.Context, authorization string, userID uint, platform *MeetingPlatform) error
	GetOAuthInfo(userID uint, platform *MeetingPlatform) (*OAuthInfo, error)
}

func (O OAuthInfo) TableName() string {
	return "oauth_info"
}

type TokenExistsError struct {
	UserID		uint
	Platform	string
}

func (t TokenExistsError) Error() string {
	return fmt.Sprintf("OAuth tokens already exist for user with id %d for platform %s", t.UserID, t.Platform)
}

