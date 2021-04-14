package domain

import (
	"gorm.io/gorm"
	"time"
)

type InviteRequest struct {
	MeetingID string
	InviterID uint
	InviteeUsername string
	MeetingPlatformID uint
}

type Invite struct {
	gorm.Model
	MeetingID string
	MeetingStartTime time.Time
	MeetingDuration int
	MeetingTitle string
	MeetingDescription string
	MeetingPlatformID uint
	MeetingJoinURL string
	InviterID uint
	InviteeID uint
}

type InviteRepository interface {
	Create(invite *Invite) (uint, error)
	GetByID(ID uint) (*Invite, error)
	GetAllByInvitee(inviteeID uint) ([]*Invite, error)
	Delete(ID uint) error
}

type InviteService interface {
	SendInvite(invite *InviteRequest) (uint, error)
	GetAllReceivedInvites(userID uint) ([]*Invite, error)
}
