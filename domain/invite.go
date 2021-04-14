package domain

import (
	"gorm.io/gorm"
	"time"
)

type Invite struct {
	gorm.Model
	MeetingID string
	MeetingEndTime time.Time
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
	SendInvite(invite *Invite) (uint, error)
	GetAllReceivedInvites(userID uint) ([]*Invite, error)
}
