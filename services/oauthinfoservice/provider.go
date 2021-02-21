package oauthinfoservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(platformService domain.MeetingPlatformService, repo domain.OAuthInfoRepository) *OAuthInfoService {
	return &OAuthInfoService{
		platformService: platformService,
		repo:            repo,
	}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.OAuthInfoService), new(*OAuthInfoService)))
)
