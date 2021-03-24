package domain

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// OAuthInfo represents OAuth token information.
type OAuthInfo struct {
	gorm.Model
	UserID            uint
	MeetingPlatformID uint
	AccessToken       string
	RefreshToken      string
	TokenType         string
	Expiry            time.Time
}

// OAuthInfoRepository manages storing and retrieving OAuthInfos.
type OAuthInfoRepository interface {
	Create(oauthToken *OAuthInfo) (uint, error)
	GetAll() ([]*OAuthInfo, error)
	GetByID(ID uint) (*OAuthInfo, error)
	GetAllByMeetingPlatformID(ID uint) ([]*OAuthInfo, error)
	GetAllByUserID(userID uint) ([]*OAuthInfo, error)
	GetByUserIDAndMeetingPlatformID(userID uint, meetingPlatformID uint) (*OAuthInfo, error)
	Update(oauthToken *OAuthInfo) error
	Delete(ID uint) error
}

// OAuthInfoService manages the creation and retrieval of OAuthInfos for a User.
type OAuthInfoService interface {
	CreateOAuthInfo(ctx context.Context, authorization string, userID uint, platform *MeetingPlatform) error
	GetOAuthInfo(userID uint, platform *MeetingPlatform) (*OAuthInfo, error)
}

// TableName returns the table name that is used in the database for OAuthInfos.
func (O OAuthInfo) TableName() string {
	return "oauth_info"
}

// TokenExistsError is an error that occurs when an OAuthInfo already exists for a User and MeetingPlatform
// combination, but a new OAuthInfo is trying to be created.
type TokenExistsError struct {
	UserID   uint
	Platform string
}

func (t TokenExistsError) Error() string {
	return fmt.Sprintf("OAuth tokens already exist for user with id %d for platform %s", t.UserID, t.Platform)
}
