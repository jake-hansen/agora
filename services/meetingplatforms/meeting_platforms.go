package meetingplatforms

import (
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

type Config struct {
	OAuthRedirectURL 	string
	OAuthClientID	 	string
	OAuthClientSecret	string
	OAuthScopes			[]string
	OAuthAuthURL		string
	OAuthTokenURL		string
}

func NewPlatform(name string, actions domain.MeetingPlatformActions, cfg *Config) *domain.MeetingPlatform {
	p := &domain.MeetingPlatform{
		Name:        name,
		RedirectURL: cfg.OAuthRedirectURL,
		OAuth:       domain.MeetingPlatformOAuthInfo{
			Config: oauth2.Config{
				ClientID:     cfg.OAuthClientID,
				ClientSecret: cfg.OAuthClientSecret,
				Endpoint:     oauth2.Endpoint{
					AuthURL:   cfg.OAuthAuthURL,
					TokenURL:  cfg.OAuthTokenURL,
				},
				Scopes:       cfg.OAuthScopes,
			},
		},
		Actions: actions,
	}

	return p
}