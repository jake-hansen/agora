package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

func MeetingProviderDTOToDomain(meetingProvider dto.MeetingProvider) *domain.MeetingProvider {
	convertedMeetingProvider := &domain.MeetingProvider{
		Name: meetingProvider.Name,
	}
	return convertedMeetingProvider
}

func MeetingProviderDomainToDTO(meetingProvider domain.MeetingProvider) *dto.MeetingProvider {
	convertedMeetingProvider := &dto.MeetingProvider{
		Name: meetingProvider.Name,
	}
	return convertedMeetingProvider
}
