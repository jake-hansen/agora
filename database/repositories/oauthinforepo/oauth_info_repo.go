package oauthinforepo

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type OAuthInfoRepo struct {
	DB	*gorm.DB
}

func (o *OAuthInfoRepo) Create(oauthToken *domain.OAuthInfo) (uint, error) {
	if err := o.DB.Create(&oauthToken).Error; err != nil {
		return 0, fmt.Errorf("error creating OAuthInfo: %w", err)
	}
	return oauthToken.ID, nil
}

func (o *OAuthInfoRepo) GetAll() ([]*domain.OAuthInfo, error) {
	var oauthTokens []*domain.OAuthInfo

	if err := o.DB.Find(&oauthTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all OAuthInfos: %w", err)
	}
	return oauthTokens, nil
}

func (o *OAuthInfoRepo) GetByID(ID uint) (*domain.OAuthInfo, error) {
	oauthToken := new(domain.OAuthInfo)
	if err := o.DB.First(oauthToken, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving OAuthInfo with id %d: %w", ID, err)
	}
	return oauthToken, nil
}

func (o *OAuthInfoRepo) GetAllByMeetingProviderId(providerID uint) ([]*domain.OAuthInfo, error) {
	var oauthTokens []*domain.OAuthInfo
	if err := o.DB.Where("meeting_provider_id = ?", providerID).Find(&oauthTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving OAuthInfos by MeetingPlatform ID %d: %w", providerID, err)
	}
	return oauthTokens, nil
}

func (o *OAuthInfoRepo) GetAllByUserID(userID uint) ([]*domain.OAuthInfo, error) {
	var oauthTokens []*domain.OAuthInfo
	if err := o.DB.Where("user_id = ?", userID).Find(&oauthTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving OAuthInfos by User ID %d: %w", userID, err)
	}
	return oauthTokens, nil
}

func (o *OAuthInfoRepo) Update(oauthToken *domain.OAuthInfo) error {
	if err := o.DB.Model(oauthToken).Updates(domain.OAuthInfo{
		UserID:            oauthToken.UserID,
		MeetingPlatformID: oauthToken.MeetingPlatformID,
		AccessToken:       oauthToken.AccessToken,
		RefreshToken:      oauthToken.RefreshToken,
	}).Error; err != nil {
		return fmt.Errorf("error updating OAuthInfo with ID %d: %w", oauthToken.ID, err)
	}
	return nil
}

func (o *OAuthInfoRepo) Delete(ID uint) error {
	if err := o.DB.Delete(&domain.OAuthInfo{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting OAuthInfo with id %d: %w", ID, err)
	}
	return nil
}

func (o *OAuthInfoRepo) GetByUserIDAndMeetingPlatformID(userID uint, meetingPlatformID uint) (*domain.OAuthInfo, error) {
	oauthToken := new(domain.OAuthInfo)

	if err := o.DB.First(oauthToken,"user_id = ? AND meeting_platform_id = ?", userID, meetingPlatformID).Error; err != nil {
		return nil, fmt.Errorf("error finding OAuthInfo for user id %d with meeting platform id %d: %w", userID, meetingPlatformID, err)
	}
	return oauthToken, nil
}


