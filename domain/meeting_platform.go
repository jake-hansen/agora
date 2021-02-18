package domain

import (
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type MeetingPlatform struct {
	gorm.Model
	Name		string
	RedirectURL string
	OAuth		MeetingPlatformOAuthInfo
	Actions		MeetingPlatformActions
}

type MeetingPlatformOAuthInfo struct {
	Config	oauth2.Config
}

type MeetingPlatformActions interface {
	CreateMeeting()
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
