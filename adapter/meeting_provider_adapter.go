package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

// MeetingProviderDomainToDTO converts the given MeetingPlatform from domain representation to DTO representation.
func MeetingProviderDomainToDTO(meetingProvider domain.MeetingPlatform) *dto.MeetingProvider {
	convertedMeetingProvider := &dto.MeetingProvider{
		Name:        meetingProvider.Name,
		RedirectURL: meetingProvider.OAuth.Config.AuthCodeURL("", oauth2.AccessTypeOffline),
	}
	return convertedMeetingProvider
}
