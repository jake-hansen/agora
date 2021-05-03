package platforms

import (
	"fmt"

	"github.com/jake-hansen/agora/platforms/webex"

	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom"
	"github.com/spf13/viper"
)

// Cfg returns a Config for a MeetingPlatform with the provided name using a Viper to get
// the configuration information.
func Cfg(v *viper.Viper, name string) *Config {
	c := Config{
		OAuthRedirectURL:  v.GetString(fmt.Sprintf("platforms.%s.oauth.url.redirect", name)),
		OAuthClientID:     v.GetString(fmt.Sprintf("platforms.%s.oauth.client.id", name)),
		OAuthClientSecret: v.GetString(fmt.Sprintf("platforms.%s.oauth.client.secret", name)),
		OAuthScopes:       v.GetStringSlice(fmt.Sprintf("platforms.%s.oauth.client.scopes", name)),
		OAuthAuthURL:      v.GetString(fmt.Sprintf("platforms.%s.oauth.url.auth", name)),
		OAuthTokenURL:     v.GetString(fmt.Sprintf("platforms.%s.oauth.url.token", name)),
	}
	return &c
}

// Provide returns ConfiguredPlatforms for the application using the provided ZoomActions, WebexActions, and Viper.
func Provide(zoomActions *zoom.ZoomActions, webexActions *webex.WebexActions, v *viper.Viper) domain.ConfiguredPlatforms {
	var platforms []*domain.MeetingPlatform

	platforms = append(platforms, NewPlatform("zoom", zoomActions, Cfg(v, "zoom")))
	platforms = append(platforms, NewPlatform("teams", nil, Cfg(v, "teams")))
	platforms = append(platforms, NewPlatform("webex", webexActions, Cfg(v, "webex")))

	return platforms
}

var (
	// ProviderSet provides ConfiguredPlatforms for use in production.
	ProviderSet = wire.NewSet(Provide, zoom.ProviderProductionSet, webex.ProviderProductionSet)
)
