package oauthinforepo

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

// OAuthInfoRepo is a repository that holds information about OAuthInfos
// backed by a database.
type OAuthInfoRepo struct {
	DB	*gorm.DB
}

// Create creates the given OAuthInfo in the database.
func (o *OAuthInfoRepo) Create(oauthToken *domain.OAuthInfo) (uint, error) {
	if err := o.DB.Create(&oauthToken).Error; err != nil {
		return 0, fmt.Errorf("error creating OAuthInfo: %w", err)
	}
	return oauthToken.ID, nil
}

// GetAll gets all of the OAuthInfos in the database.
func (o *OAuthInfoRepo) GetAll() ([]*domain.OAuthInfo, error) {
	var oauthTokens []*domain.OAuthInfo

	if err := o.DB.Find(&oauthTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all OAuthInfos: %w", err)
	}
	return oauthTokens, nil
}

// GetByID retrieves an OAuthInfo based on the given ID.
func (o *OAuthInfoRepo) GetByID(ID uint) (*domain.OAuthInfo, error) {
	oauthToken := new(domain.OAuthInfo)
	if err := o.DB.First(oauthToken, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving OAuthInfo with id %d: %w", ID, err)
	}
	return oauthToken, nil
}

// GetAllByMeetingPlatformID retrieves all OAuthInfos based on the given MeetingPlatform ID.
func (o *OAuthInfoRepo) GetAllByMeetingPlatformID(ID uint) ([]*domain.OAuthInfo, error) {
	var oauthTokens []*domain.OAuthInfo
	if err := o.DB.Where("meeting_platform_id = ?", ID).Find(&oauthTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving OAuthInfos by MeetingPlatform ID %d: %w", ID, err)
	}
	return oauthTokens, nil
}

// GetAllByUserID retrieves all OAuthInfos based on the given User ID.
func (o *OAuthInfoRepo) GetAllByUserID(userID uint) ([]*domain.OAuthInfo, error) {
	var oauthTokens []*domain.OAuthInfo
	if err := o.DB.Where("user_id = ?", userID).Find(&oauthTokens).Error; err != nil {
		return nil, fmt.Errorf("error retrieving OAuthInfos by User ID %d: %w", userID, err)
	}
	return oauthTokens, nil
}

// Update updates the given OAuthInfo. The ID of the given OAuthInfo needs to be set
// in order to find the existing record in the database.
func (o *OAuthInfoRepo) Update(oauthToken *domain.OAuthInfo) error {
	if err := o.DB.Model(oauthToken).Updates(domain.OAuthInfo{
		UserID:            oauthToken.UserID,
		MeetingPlatformID: oauthToken.MeetingPlatformID,
		AccessToken:       oauthToken.AccessToken,
		RefreshToken:      oauthToken.RefreshToken,
		TokenType: 		   oauthToken.TokenType,
		Expiry: 		   oauthToken.Expiry,
	}).Error; err != nil {
		return fmt.Errorf("error updating OAuthInfo with ID %d: %w", oauthToken.ID, err)
	}
	return nil
}

// Delete deletes the OAuthInfo with the given ID.
func (o *OAuthInfoRepo) Delete(ID uint) error {
	if err := o.DB.Delete(&domain.OAuthInfo{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting OAuthInfo with id %d: %w", ID, err)
	}
	return nil
}

// GetByUserIDAndMeetingPlatformID retrieves the OAuthInfo by the given User ID and MeetingPlatform ID.
func (o *OAuthInfoRepo) GetByUserIDAndMeetingPlatformID(userID uint, meetingPlatformID uint) (*domain.OAuthInfo, error) {
	oauthToken := new(domain.OAuthInfo)

	if err := o.DB.First(oauthToken,"user_id = ? AND meeting_platform_id = ?", userID, meetingPlatformID).Error; err != nil {
		return nil, fmt.Errorf("error finding OAuthInfo for user id %d with meeting platform id %d: %w", userID, meetingPlatformID, err)
	}
	return oauthToken, nil
}


