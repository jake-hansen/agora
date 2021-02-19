package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type OAuthInfo struct {
	gorm.Model
	UserID			  uint
	MeetingProviderID uint
	AccessToken		  string
	RefreshToken 	  string
	TokenType		  string
	Expiry			  time.Time
}

type OAuthInfoRepository interface {
	Create(oauthToken *OAuthInfo) (uint, error)
	GetAll() ([]*OAuthInfo, error)
	GetByID(ID uint) (*OAuthInfo, error)
	GetAllByMeetingProviderId(providerID uint) ([]*OAuthInfo, error)
	GetAllByUserID(userID uint) ([]*OAuthInfo, error)
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

