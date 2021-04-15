package domain

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"time"
)

type MeetingDuration time.Duration

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
	MeetingDuration MeetingDuration
	MeetingTitle string
	MeetingDescription string
	MeetingPlatformID uint
	MeetingJoinURL string
	InviterID uint
	InviteeID uint
}

func (m *MeetingDuration) Scan(src interface{}) error {
	var duration time.Duration = time.Duration(src.(int64))
	*m = MeetingDuration(time.Minute * duration)
	return nil
}

func (m MeetingDuration) Value() (driver.Value, error) {
	return time.Duration(m).Minutes(), nil
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
