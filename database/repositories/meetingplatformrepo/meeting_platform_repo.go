package meetingplatformrepo

import (
	"fmt"

	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

// MeetingPlatformRepo is a repository that holds information about MeetingPlatforms
// backed by a database.
type MeetingPlatformRepo struct {
	DB *gorm.DB
}

// Create creates the given MeetingPlatform in the database.
func (m *MeetingPlatformRepo) Create(meetingProvider *domain.MeetingPlatform) (uint, error) {
	if err := m.DB.Create(&meetingProvider).Error; err != nil {
		return 0, fmt.Errorf("error creating meeting provider: %w", err)
	}
	return meetingProvider.ID, nil
}

// GetAll retrieves all MeetingPlatforms in the database.
func (m *MeetingPlatformRepo) GetAll() ([]*domain.MeetingPlatform, error) {
	var meetingProviders []*domain.MeetingPlatform

	if err := m.DB.Find(&meetingProviders).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all meeting providers: %w", err)
	}
	return meetingProviders, nil
}

// GetByID retrieves the MeetingPlatform in the database with the provided ID.
func (m *MeetingPlatformRepo) GetByID(ID uint) (*domain.MeetingPlatform, error) {
	meetingProvider := new(domain.MeetingPlatform)
	if err := m.DB.First(meetingProvider, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving meeting provider with id %d: %w", ID, err)
	}
	return meetingProvider, nil
}

// GetByPlatformName retrieves the MeetingPlatform in the database with the provided name.
func (m *MeetingPlatformRepo) GetByPlatformName(providerName string) (*domain.MeetingPlatform, error) {
	meetingProvider := new(domain.MeetingPlatform)
	if err := m.DB.Where("name = ?", providerName).First(meetingProvider).Error; err != nil {
		return nil, fmt.Errorf("error retrieving meeting provider by name %s: %w", providerName, err)
	}
	return meetingProvider, nil
}

// Update updates the given MeetingPlatform. The ID of the given MeetingPlatform needs to be set
// in order to find the existing record in the database.
func (m *MeetingPlatformRepo) Update(meetingProvider *domain.MeetingPlatform) error {
	if err := m.DB.Model(meetingProvider).Updates(domain.MeetingPlatform{
		Name: meetingProvider.Name,
	}).Error; err != nil {
		return fmt.Errorf("error updating meeting provider with id %d: %w", meetingProvider.ID, err)
	}
	return nil
}

// Delete deletes the MeetingPlatform from the database with the given ID.
func (m *MeetingPlatformRepo) Delete(ID uint) error {
	if err := m.DB.Delete(&domain.MeetingPlatform{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting meeting provider with id %d: %w", ID, err)
	}
	return nil
}
