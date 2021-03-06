package platforms

import (
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

// Config contains OAuth configuration information for a MeetingPlatform.
type Config struct {
	OAuthRedirectURL  string
	OAuthClientID     string
	OAuthClientSecret string
	OAuthScopes       []string
	OAuthAuthURL      string
	OAuthTokenURL     string
}

// NewPlatform returns a MeetingPlatform with the given name, actions, and config.
func NewPlatform(name string, actions domain.MeetingPlatformActions, cfg *Config) *domain.MeetingPlatform {
	p := &domain.MeetingPlatform{
		Name: name,
		OAuth: domain.MeetingPlatformOAuthInfo{
			Config: oauth2.Config{
				ClientID:     cfg.OAuthClientID,
				ClientSecret: cfg.OAuthClientSecret,
				Endpoint: oauth2.Endpoint{
					AuthURL:  cfg.OAuthAuthURL,
					TokenURL: cfg.OAuthTokenURL,
				},
				Scopes:      cfg.OAuthScopes,
				RedirectURL: cfg.OAuthRedirectURL,
			},
		},
		Actions: actions,
	}

	return p
}
