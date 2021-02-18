package meetingplatformservice

import (
	"context"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

type MeetingPlatformService struct {
	repo domain.MeetingPlatformRepository
}

func (m *MeetingPlatformService) Create(meetingProvider *domain.MeetingPlatform) (uint, error) {
	return m.repo.Create(meetingProvider)
}

func (m *MeetingPlatformService) GetAll() ([]*domain.MeetingPlatform, error) {
	return m.repo.GetAll()
}

func (m *MeetingPlatformService) GetByID(ID uint) (*domain.MeetingPlatform, error) {
	return m.repo.GetByID(ID)
}

func (m *MeetingPlatformService) GetByProviderName(providerName string) (*domain.MeetingPlatform, error) {
	return m.repo.GetByProviderName(providerName)
}

func (m *MeetingPlatformService) Update(meetingProvider *domain.MeetingPlatform) error {
	return m.repo.Update(meetingProvider)
}

func (m *MeetingPlatformService) Delete(ID uint) error {
	return m.repo.Delete(ID)
}

func (m *MeetingPlatformService) GetOAuthTokens(ctx context.Context, authorization string, platform *domain.MeetingPlatform) (*oauth2.Token, error) {
	token, err := platform.OAuth.Config.Exchange(ctx, authorization)
	return token, err
}

