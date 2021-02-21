package domain

import (
	"context"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type MeetingPlatform struct {
	gorm.Model
	Name		string
	OAuth		MeetingPlatformOAuthInfo	`gorm:"-"`
	Actions		MeetingPlatformActions		`gorm:"-"`
}

type MeetingPlatformOAuthInfo struct {
	Config	oauth2.Config
}

type MeetingPlatformActions interface {
	CreateMeeting()
}

type MeetingPlatformRepository interface {
	Create(platform *MeetingPlatform) (uint, error)
	GetAll() ([]*MeetingPlatform, error)
	GetByID(ID uint) (*MeetingPlatform, error)
	GetByPlatformName(name string) (*MeetingPlatform, error)
	Update(platform *MeetingPlatform) error
	Delete(ID uint) error
}

type MeetingPlatformService interface {
	Save(platform *MeetingPlatform) (uint, error)
	GetAll() ([]*MeetingPlatform, error)
	Delete(ID uint) error
	GetByID(ID uint) (*MeetingPlatform, error)
	GetByPlatformName(name string) (*MeetingPlatform, error)
    GetOAuthToken(ctx context.Context, authorization string, platform *MeetingPlatform) (*oauth2.Token, error)
	RefreshOAuthToken(ctx context.Context, token *oauth2.Token, platform *MeetingPlatform) (*oauth2.Token, error)
}
