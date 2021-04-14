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

	invitee, err := s.userService.GetByUsername(invite.InviteeUsername)
	if err != nil {
		return 0, err
	}

	domainInvite := &domain.Invite{
		MeetingID:          meeting.ID,
		MeetingStartTime:   meeting.StartTime,
		MeetingDuration: 	int(meeting.Duration.Minutes()),
		MeetingTitle:       meeting.Title,
		MeetingDescription: meeting.Description,
		MeetingPlatformID:  platform.ID,
		MeetingJoinURL: 	meeting.JoinURL,
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
