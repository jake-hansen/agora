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
		JoinURL: 	 meeting.JoinURL,
		StartURL: 	 meeting.StartURL,
	}
	return dtoMeeting
}

func DomainMeetingPageToDTOMeetingPage(page *domain.Page) *dto.MeetingPage {
	var meetings []*dto.Meeting
	for _, meeting := range page.Records {
		meetings = append(meetings, MeetingDomainToDTO(meeting.(*domain.Meeting)))
	}

	dtoMeetingPage := &dto.MeetingPage{
		PageCount:    page.PageCount,
		PageNumber:   page.PageNumber,
		PageSize:     page.PageSize,
		TotalRecords: page.TotalRecords,
		Records:      meetings,
	}
	return dtoMeetingPage
}
