package adapter

import (
	"time"

	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// InviteDomainToDTO converts the given Invite from domain representation to DTO representation.
func InviteDomainToDTO(invite *domain.Invite) *dto.Invite {
	dtoInvite := &dto.Invite{
		ID:                invite.ID,
		MeetingPlatformID: invite.MeetingPlatformID,
		InviterID:         invite.InviterID,
		Meeting: dto.Meeting{
			ID:          invite.MeetingID,
			Title:       invite.MeetingTitle,
			StartTime:   invite.MeetingStartTime.Format(time.RFC3339),
			Duration:    dto.MeetingDuration(invite.MeetingDuration),
			Description: invite.MeetingDescription,
			JoinURL:     invite.MeetingJoinURL,
		},
	}

	return dtoInvite
}

// InviteRequestDTOToDomain converts the given InviteRequest from DTO representation to domain representation.
func InviteRequestDTOToDomain(inviteRequest *dto.InviteRequest) *domain.InviteRequest {
	domainInviteRequest := &domain.InviteRequest{
		MeetingID:         inviteRequest.MeetingID,
		InviterID:         inviteRequest.InviterID,
		InviteeUsername:   inviteRequest.InviteeUsername,
		MeetingPlatformID: inviteRequest.MeetingPlatformID,
	}
	return domainInviteRequest
}
