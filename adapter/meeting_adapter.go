package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
	"time"
)

func MeetingDTOToDomain(meeting *dto.Meeting) *domain.Meeting {
	domainMeeting := &domain.Meeting{
		Title:       meeting.Title,
		StartTime:   meeting.StartTime,
		Duration:    time.Duration(meeting.Duration),
		Description: meeting.Description,
	}
	return domainMeeting
}

func MeetingDomainToDTO(meeting *domain.Meeting) *dto.Meeting {
	dtoMeeting := &dto.Meeting{
		Title:       meeting.Title,
		StartTime:   meeting.StartTime,
		Duration:    dto.MeetingDuration(meeting.Duration),
		Description: meeting.Description,
	}
	return dtoMeeting
}
