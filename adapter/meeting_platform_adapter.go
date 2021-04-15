package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
)

// MeetingPlatformDomainToDTO converts the given MeetingPlatform from domain representation to DTO representation.
func MeetingPlatformDomainToDTO(meetingPlatform domain.MeetingPlatform) *dto.MeetingPlatform {
	convertedMeetingProvider := &dto.MeetingPlatform{
		ID:          meetingPlatform.ID,
		Name:        meetingPlatform.Name,
		RedirectURL: meetingPlatform.OAuth.Config.AuthCodeURL("", oauth2.AccessTypeOffline),
	}
	return convertedMeetingProvider
}
