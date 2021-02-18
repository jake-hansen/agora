package domain

import "gorm.io/gorm"

type MeetingPlatform struct {
	gorm.Model
	Name		string
	RedirectURL string
}

type MeetingPlatformRepository interface {
	Create(meetingProvider *MeetingPlatform) (uint, error)
	GetAll() ([]*MeetingPlatform, error)
	GetByID(ID uint) (*MeetingPlatform, error)
	GetByProviderName(providerName string) (*MeetingPlatform, error)
	Update(meetingProvider *MeetingPlatform) error
	Delete(ID uint) error
}

type MeetingPlatformService interface {
	Create(meetingProvider *MeetingPlatform) (uint, error)
	GetAll() ([]*MeetingPlatform, error)
	GetByID(ID uint) (*MeetingPlatform, error)
	GetByProviderName(providerName string) (*MeetingPlatform, error)
	Update(meetingProvider *MeetingPlatform) error
	Delete(ID uint) error
}
