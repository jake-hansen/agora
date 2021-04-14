package simpleinviteservice

import "github.com/jake-hansen/agora/domain"

type SimpleInviteService struct {
	inviteRepo domain.InviteRepository
}

func (s *SimpleInviteService) SendInvite(invite *domain.Invite) (uint, error) {
	return s.inviteRepo.Create(invite)
}

func (s *SimpleInviteService) GetAllReceivedInvites(userID uint) ([]*domain.Invite, error) {
	invites, err := s.inviteRepo.GetAllByInvitee(userID)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

