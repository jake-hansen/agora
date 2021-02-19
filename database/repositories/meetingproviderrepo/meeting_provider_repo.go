package meetingproviderrepo

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type MeetingProviderRepo struct {
	DB *gorm.DB
}

func (m *MeetingProviderRepo) Create(meetingProvider *domain.MeetingProvider) (uint, error) {
	if err := m.DB.Create(&meetingProvider).Error; err != nil {
		return 0, fmt.Errorf("error creating meeting provider: %w", err)
	}
	return meetingProvider.ID, nil
}

func (m *MeetingProviderRepo) GetAll() ([]*domain.MeetingProvider, error) {
	var meetingProviders []*domain.MeetingProvider

	if err := m.DB.Find(&meetingProviders).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all meeting providers: %w", err)
	}
	return meetingProviders, nil
}

func (m *MeetingProviderRepo) GetByID(ID uint) (*domain.MeetingProvider, error) {
	meetingProvider := new(domain.MeetingProvider)
	if err := m.DB.First(meetingProvider, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving meeting provider with id %d: %w", ID, err)
	}
	return meetingProvider, nil
}

func (m *MeetingProviderRepo) GetByProviderName(providerName string) (*domain.MeetingProvider, error) {
	meetingProvider := new(domain.MeetingProvider)
	if err := m.DB.Where("name = ?", providerName).First(meetingProvider).Error; err != nil {
		return nil, fmt.Errorf("error retrieving meeting provider by name %s: %w", providerName, err)
	}
	return meetingProvider, nil
}

func (m *MeetingProviderRepo) Update(meetingProvider *domain.MeetingProvider) error {
	if err := m.DB.Model(meetingProvider).Updates(domain.MeetingProvider{
		Name:  meetingProvider.Name,
	}).Error; err != nil {
		return fmt.Errorf("error updating meeting provider with id %d: %w", meetingProvider.ID, err)
	}
	return nil
}

func (m *MeetingProviderRepo) Delete(ID uint) error {
	if err := m.DB.Delete(&domain.MeetingProvider{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting meeting provider with id %d: %w", ID, err)
	}
	return nil
}

