package meetingplatforms

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/meetingplatforms/zoom"
	"github.com/spf13/viper"
)

func Cfg(v *viper.Viper, name string) *Config {
	c := Config{
		OAuthRedirectURL:  v.GetString(fmt.Sprintf("platforms.%s.oauth.url.redirect", name)),
		OAuthClientID:     v.GetString(fmt.Sprintf("platforms.%s.oauth.client.id", name)),
		OAuthClientSecret: v.GetString(fmt.Sprintf("platforms.%s.oauth.client.secret", name)),
		OAuthScopes:       v.GetStringSlice(fmt.Sprintf("platforms.%s.oauth.client.id", name)),
		OAuthAuthURL:      v.GetString(fmt.Sprintf("platforms.%s.oauth.url.auth", name)),
		OAuthTokenURL:     v.GetString(fmt.Sprintf("platforms.%s.oauth.url.token", name)),
	}
	return &c
}

func Provide(zoomActions *zoom.Zoom, v *viper.Viper) []*domain.MeetingPlatform {
	var platforms []*domain.MeetingPlatform

	platforms = append(platforms, NewPlatform("Zoom", zoomActions, Cfg(v, "zoom")))

	return platforms
}

var (
	ProviderSet = wire.NewSet(Provide, zoom.ProviderProductionSet)
)