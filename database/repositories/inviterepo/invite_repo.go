package inviterepo

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jake-hansen/agora/database/repositories"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
	"strconv"
)

type InviteRepo struct {
	DB *gorm.DB
}

func (i *InviteRepo) Create(invite *domain.Invite) (uint, error) {
	if err := i.DB.Create(&invite).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return 0, repositories.NewDuplicateEntryError(repositories.DATABASE_ACTION_CREATE, "invite", "", "")
		}
		return 0, fmt.Errorf("error creating invite: %w", err)
	}
	return invite.ID, nil
}

func (i *InviteRepo) GetByID(ID uint) (*domain.Invite, error) {
	invite := new(domain.Invite)
	if err := i.DB.First(invite, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.NewNotFoundError("delete", "invite", strconv.Itoa(int(ID)), "by id")
		}
		return nil, fmt.Errorf("error retrieving invite with id %d: %w", ID, err)
	}
	return invite, nil
}

func (i *InviteRepo) GetAllByInvitee(inviteeID uint) ([]*domain.Invite, error) {
	var invites []*domain.Invite
	if err := i.DB.Where("invitee_id = ? AND deleted_at IS NULL", inviteeID).Find(&invites).Error; err != nil {
		return nil, fmt.Errorf("error retrieving invites by invitee id %d: %w", inviteeID, err)
	}
	return invites, nil
}

func (i *InviteRepo) GetAllByInviter(inviterID uint) ([]*domain.Invite, error) {
	var invites []*domain.Invite
	if err := i.DB.Where("inviter_id = ? AND deleted_at IS NULL", inviterID).Find(&invites).Error; err != nil {
		return nil, fmt.Errorf("error retrieving invites by inviter id %d: %w", inviterID, err)
	}
	return invites, nil
}

func (i *InviteRepo) Delete(ID uint) error {
	if err := i.DB.Unscoped().Delete(&domain.Invite{}, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repositories.NewNotFoundError("delete", "invite", strconv.Itoa(int(ID)), "by id")
		}
		return fmt.Errorf("error deleting invite with id %d: %w", ID, err)
	}
	return nil
}

func (i *InviteRepo) DeleteAllByMeetingID(meetingID string) error {
	if err := i.DB.Unscoped().Delete(&domain.Invite{}, "meeting_id = ?", meetingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repositories.NewNotFoundError("delete", "invite", meetingID, "by meeting id")
		}
		return fmt.Errorf("error deleting invites with meeting id %s: %w", meetingID, err)
	}
	return nil
}
