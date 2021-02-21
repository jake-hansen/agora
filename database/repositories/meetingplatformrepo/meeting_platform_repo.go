package meetingplatformrepo

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type MeetingPlatformRepo struct {
	DB *gorm.DB
}

func (m *MeetingPlatformRepo) Create(meetingProvider *domain.MeetingPlatform) (uint, error) {
	if err := m.DB.Create(&meetingProvider).Error; err != nil {
		return 0, fmt.Errorf("error creating meeting provider: %w", err)
	}
	return meetingProvider.ID, nil
}

func (m *MeetingPlatformRepo) GetAll() ([]*domain.MeetingPlatform, error) {
	var meetingProviders []*domain.MeetingPlatform

	if err := m.DB.Find(&meetingProviders).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all meeting providers: %w", err)
	}
	return meetingProviders, nil
}

func (m *MeetingPlatformRepo) GetByID(ID uint) (*domain.MeetingPlatform, error) {
	meetingProvider := new(domain.MeetingPlatform)
	if err := m.DB.First(meetingProvider, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving meeting provider with id %d: %w", ID, err)
	}
	return meetingProvider, nil
}

func (m *MeetingPlatformRepo) GetByPlatformName(providerName string) (*domain.MeetingPlatform, error) {
	meetingProvider := new(domain.MeetingPlatform)
	if err := m.DB.Where("name = ?", providerName).First(meetingProvider).Error; err != nil {
		return nil, fmt.Errorf("error retrieving meeting provider by name %s: %w", providerName, err)
	}
	return meetingProvider, nil
}

func (m *MeetingPlatformRepo) Update(meetingProvider *domain.MeetingPlatform) error {
	if err := m.DB.Model(meetingProvider).Updates(domain.MeetingPlatform{
		Name:  meetingProvider.Name,
	}).Error; err != nil {
		return fmt.Errorf("error updating meeting provider with id %d: %w", meetingProvider.ID, err)
	}
	return nil
}

func (m *MeetingPlatformRepo) Delete(ID uint) error {
	if err := m.DB.Delete(&domain.MeetingPlatform{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting meeting provider with id %d: %w", ID, err)
	}
	return nil
}

