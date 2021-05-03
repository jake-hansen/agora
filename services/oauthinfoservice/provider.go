package oauthinfoservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns an OAuthInfoService configured with the provided MeetingPlatformService and OAuthInfoRepository.
func Provide(platformService domain.MeetingPlatformService, repo domain.OAuthInfoRepository) *OAuthInfoService {
	return &OAuthInfoService{
		platformService: platformService,
		repo:            repo,
	}
}

var (
	// ProviderProductionSet provides an OAuthInfoService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.OAuthInfoService), new(*OAuthInfoService)))
)
