package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
	"time"
)

func InviteDomainToDTO(invite *domain.Invite) *dto.Invite {
	dtoInvite := &dto.Invite{
		MeetingPlatformID: invite.MeetingPlatformID,
		InviterID:       invite.InviterID,
		InviteeID:       invite.InviteeID,
		Meeting:         dto.Meeting{
			ID:          invite.MeetingID,
			Title:       invite.MeetingTitle,
			StartTime:   invite.MeetingStartTime.Format(time.RFC3339),
			Duration: 	 dto.MeetingDuration(invite.MeetingDuration),
			Description: invite.MeetingDescription,
			JoinURL:     invite.MeetingJoinURL,
		},
	}

	return dtoInvite
}

func InviteRequestDTOToDomain(inviteRequest *dto.InviteRequest) *domain.InviteRequest {
	domainInviteRequest := &domain.InviteRequest{
		MeetingID:         inviteRequest.MeetingID,
		InviterID:         inviteRequest.InviterID,
		InviteeUsername:   inviteRequest.InviteeUsername,
		MeetingPlatformID: inviteRequest.MeetingPlatformID,
	}
	return domainInviteRequest
}
