package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

func MeetingProviderDTOToDomain(meetingProvider dto.MeetingProvider) *domain.MeetingPlatform {
	convertedMeetingProvider := &domain.MeetingPlatform{
		Name: meetingProvider.Name,
		RedirectURL: meetingProvider.RedirectURL,
	}
	return convertedMeetingProvider
}

func MeetingProviderDomainToDTO(meetingProvider domain.MeetingPlatform) *dto.MeetingProvider {
	convertedMeetingProvider := &dto.MeetingProvider{
		Name: meetingProvider.Name,
		RedirectURL: meetingProvider.RedirectURL,
	}
	return convertedMeetingProvider
}
