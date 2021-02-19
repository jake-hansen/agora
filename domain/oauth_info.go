package domain

import "gorm.io/gorm"

type OAuthInfo struct {
	gorm.Model
	UserID			  uint
	MeetingProviderID uint
	AccessToken		  string
	RefreshToken 	  string
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

type OAuthTokenService interface {

}

func (O OAuthInfo) TableName() string {
	return "oauth_info"
}

