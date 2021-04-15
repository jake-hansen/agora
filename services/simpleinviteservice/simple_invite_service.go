package simpleinviteservice

import (
	"github.com/jake-hansen/agora/domain"
)

type SimpleInviteService struct {
	inviteRepo domain.InviteRepository
	meetingService domain.MeetingPlatformService
	oauthService domain.OAuthInfoService
	userService domain.UserService
}

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

func (s *SimpleInviteService) GetAllReceivedInvites(userID uint) ([]*domain.Invite, error) {
	invites, err := s.inviteRepo.GetAllByInvitee(userID)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

func (s *SimpleInviteService) GetAllSentInvites(userID uint) ([]*domain.Invite, error) {
	invites, err := s.inviteRepo.GetAllByInviter(userID)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

func (s *SimpleInviteService) DeleteInvite(inviteID uint) error {
	return s.inviteRepo.Delete(inviteID)
}

func (s *SimpleInviteService) GetInvite(inviteID uint) (*domain.Invite, error) {
	return s.inviteRepo.GetByID(inviteID)
}


