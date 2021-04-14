package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

/*func InviteDomainToDTO(invite *domain.Invite) *dto.Invite {
	dtoInvite := &dto.Invite{
		MeetingID:       invite.MeetingID,
		MeetingPlatform: ,
		InviterID:       0,
		Invitee:         "",
	}

	return dtoInvite
}*/

func InviteRequestDTOToDomain(inviteRequest *dto.InviteRequest) *domain.InviteRequest {
	domainInviteRequest := &domain.InviteRequest{
		MeetingID:         inviteRequest.MeetingID,
		InviterID:         inviteRequest.InviterID,
		InviteeUsername:   inviteRequest.InviteeUsername,
		MeetingPlatformID: inviteRequest.MeetingPlatformID,
	}
	return domainInviteRequest
}
