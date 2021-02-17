package domain

import "gorm.io/gorm"

type MeetingProvider struct {
	gorm.Model
	Name		string
	RedirectURL string
}

type MeetingProviderRepository interface {
	Create(meetingProvider *MeetingProvider) (uint, error)
	GetAll() ([]*MeetingProvider, error)
	GetByID(ID uint) (*MeetingProvider, error)
	GetByProviderName(providerName string) (*MeetingProvider, error)
	Update(meetingProvider *MeetingProvider) error
	Delete(ID uint) error
}

type MeetingProviderService interface {
	Create(meetingProvider *MeetingProvider) (uint, error)
	GetAll() ([]*MeetingProvider, error)
	GetByID(ID uint) (*MeetingProvider, error)
	GetByProviderName(providerName string) (*MeetingProvider, error)
	Update(meetingProvider *MeetingProvider) error
	Delete(ID uint) error
}
