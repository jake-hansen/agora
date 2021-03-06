package domain

import (
	"context"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

// MeetingPlatform represents an external MeetingPlatform.
type MeetingPlatform struct {
	gorm.Model
	Name		string
	OAuth		MeetingPlatformOAuthInfo	`gorm:"-"`
	Actions		MeetingPlatformActions		`gorm:"-"`
}

// ConfiguredPlatforms are a slice of MeetingPlatforms that are configured for
// the application.
type ConfiguredPlatforms []*MeetingPlatform

// MeetingPlatformOAuthInfo contains configuration information about the MeetingPlatform's
// OAuth implementation.
type MeetingPlatformOAuthInfo struct {
	Config	oauth2.Config
}

// MeetingPlatformActions are implementations of the MeetingPlatform's API.
type MeetingPlatformActions interface {
	CreateMeeting(oauth OAuthInfo, meeting *Meeting) (*Meeting, error)
	GetMeetings(oauth OAuthInfo) (*Page, error)
	GetMeeting(oauth OAuthInfo, meetingID string) (*Meeting, error)
}

// MeetingPlatformRepository stores information about MeetingPlatforms.
type MeetingPlatformRepository interface {
	Create(platform *MeetingPlatform) (uint, error)
	GetAll() ([]*MeetingPlatform, error)
	GetByID(ID uint) (*MeetingPlatform, error)
	GetByPlatformName(name string) (*MeetingPlatform, error)
	Update(platform *MeetingPlatform) error
	Delete(ID uint) error
}

// MeetingPlatformServices manages CRUD operations on a MeetingPlatform as well as
// OAuth.
type MeetingPlatformService interface {
	Save(platform *MeetingPlatform) (uint, error)
	GetAll() ([]*MeetingPlatform, error)
	Delete(ID uint) error
	GetByID(ID uint) (*MeetingPlatform, error)
	GetByPlatformName(name string) (*MeetingPlatform, error)
    GetOAuthToken(ctx context.Context, authorization string, platform *MeetingPlatform) (*oauth2.Token, error)
	RefreshOAuthToken(ctx context.Context, token *oauth2.Token, platform *MeetingPlatform) (*oauth2.Token, error)
}
