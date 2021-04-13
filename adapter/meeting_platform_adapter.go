package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

// MeetingPlatformDomainToDTO converts the given MeetingPlatform from domain representation to DTO representation.
func MeetingPlatformDomainToDTO(meetingProvider domain.MeetingPlatform) *dto.MeetingPlatform {
	convertedMeetingProvider := &dto.MeetingPlatform{
		Name:        meetingProvider.Name,
		RedirectURL: meetingProvider.OAuth.Config.AuthCodeURL("", oauth2.AccessTypeOffline),
	}
	return convertedMeetingProvider
}
