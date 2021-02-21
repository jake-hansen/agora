package meetingplatformservice

import (
	"context"
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
	"strings"
)

type MeetingPlatformService struct {
	dbRepo 				domain.MeetingPlatformRepository
	configuredPlatforms map[string]*domain.MeetingPlatform
}

func New(repository domain.MeetingPlatformRepository, configuredPlatforms []*domain.MeetingPlatform) *MeetingPlatformService {
	p := MeetingPlatformService{
		dbRepo:              repository,
		configuredPlatforms: make(map[string]*domain.MeetingPlatform),
	}

	for _, platform := range configuredPlatforms {
		p.configuredPlatforms[platform.Name] = platform
	}

	return &p
}

func (m *MeetingPlatformService) Save(meetingProvider *domain.MeetingPlatform) (uint, error) {
	return m.dbRepo.Create(meetingProvider)
}

func (m *MeetingPlatformService) Delete(ID uint) error {
	return m.dbRepo.Delete(ID)
}

func (m *MeetingPlatformService) GetAll() ([]*domain.MeetingPlatform, error) {
	var platformList []*domain.MeetingPlatform
	dbPlatforms, err := m.dbRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, dbPlatform := range dbPlatforms {
		p, err := m.combine(dbPlatform)
		if err != nil {
			return nil, err
		}
		platformList = append(platformList, p)
	}

	return platformList, nil
}

func (m *MeetingPlatformService) GetByID(ID uint) (*domain.MeetingPlatform, error) {
	dbPlatform, err := m.dbRepo.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return m.combine(dbPlatform)
}

func (m *MeetingPlatformService) GetByPlatformName(name string) (*domain.MeetingPlatform, error) {
	dbPlatform, err := m.dbRepo.GetByPlatformName(name)
	if err != nil {
		return nil, err
	}

	return m.combine(dbPlatform)
}

func (m *MeetingPlatformService) GetOAuthToken(ctx context.Context, authorization string, platform *domain.MeetingPlatform) (*oauth2.Token, error) {
	token, err := platform.OAuth.Config.Exchange(ctx, authorization)
	return token, err
}

func (m *MeetingPlatformService) RefreshOAuthToken(ctx context.Context, token *oauth2.Token, platform *domain.MeetingPlatform) (*oauth2.Token, error) {
	tokenSource := platform.OAuth.Config.TokenSource(ctx, token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("an error occurred while refreshing an access token: %w", err)
	}

	return newToken, nil
}

func (m *MeetingPlatformService) combine(platform *domain.MeetingPlatform) (*domain.MeetingPlatform, error) {
	p := m.configuredPlatforms[strings.ToLower(platform.Name)]
	if p == nil {
		return nil,fmt.Errorf("meeting platform with name %s is not configured", platform.Name)
	} else {
		p.Model = platform.Model
		return p, nil
	}
}
