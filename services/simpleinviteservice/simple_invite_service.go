package simpleinviteservice

import (
	"github.com/jake-hansen/agora/domain"
)

// SimpoleInviteService manages operations on Invites.
type SimpleInviteService struct {
	inviteRepo     domain.InviteRepository
	meetingService domain.MeetingPlatformService
	oauthService   domain.OAuthInfoService
	userService    domain.UserService
}

// SendInvite sends an Invite to another user based on the information contained in the InviteRequest.
func (s *SimpleInviteService) SendInvite(invite *domain.InviteRequest) (uint, error) {
	invitee, err := s.userService.GetByUsername(invite.InviteeUsername)
	if err != nil {
		return 0, err
	}

	if invitee.ID == invite.InviterID {
		return 0, NewInviterSameAsInviteeErr()
	}

	platform, err := s.meetingService.GetByID(invite.MeetingPlatformID)
	if err != nil {
		return 0, err
	}

	oauth, err := s.oauthService.GetOAuthInfo(invite.InviterID, platform)
	if err != nil {
		return 0, err
	}

	meeting, err := platform.Actions.GetMeeting(*oauth, invite.MeetingID)
	if err != nil {
		return 0, err
	}

	domainInvite := &domain.Invite{
		MeetingID:          meeting.ID,
		MeetingStartTime:   meeting.StartTime,
		MeetingDuration:    domain.MeetingDuration(meeting.Duration),
		MeetingTitle:       meeting.Title,
		MeetingDescription: meeting.Description,
		MeetingPlatformID:  platform.ID,
		MeetingJoinURL:     meeting.JoinURL,
		InviterID:          invite.InviterID,
		InviteeID:          invitee.ID,
	}

	return s.inviteRepo.Create(domainInvite)
}

// GetAllReceivedInvites returns all Invites that have been received by the provided
// userID.
func (s *SimpleInviteService) GetAllReceivedInvites(userID uint) ([]*domain.Invite, error) {
	invites, err := s.inviteRepo.GetAllByInvitee(userID)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

// GetAllSentInvites returns all Invites that have been sent by the provided userID.
func (s *SimpleInviteService) GetAllSentInvites(userID uint) ([]*domain.Invite, error) {
	invites, err := s.inviteRepo.GetAllByInviter(userID)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

// DeleteInvite deletes the Invite with the provided ID.
func (s *SimpleInviteService) DeleteInvite(inviteID uint) error {
	return s.inviteRepo.Delete(inviteID)
}

// GetInvite returns the Invite with the provided ID.
func (s *SimpleInviteService) GetInvite(inviteID uint) (*domain.Invite, error) {
	return s.inviteRepo.GetByID(inviteID)
}

// DeleteAllInvitesByMeetingID deletes all Invites which contain the provided meetingID.
func (s *SimpleInviteService) DeleteAllInvitesByMeetingID(meetingID string) error {
	return s.inviteRepo.DeleteAllByMeetingID(meetingID)
}
