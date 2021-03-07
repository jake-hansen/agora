package adapter

import (
	"time"

	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
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
		ID:          meeting.ID,
		Title:       meeting.Title,
		StartTime:   meeting.StartTime,
		Duration:    dto.MeetingDuration(meeting.Duration),
		Description: meeting.Description,
		JoinURL:     meeting.JoinURL,
		StartURL:    meeting.StartURL,
	}
	return dtoMeeting
}

func MeetingPageDomainToDTO(page *domain.Page) *dto.MeetingPage {
	var meetings []*dto.Meeting
	for _, meeting := range page.Records {
		meetings = append(meetings, MeetingDomainToDTO(meeting.(*domain.Meeting)))
	}

	dtoMeetingPage := &dto.MeetingPage{
		PageCount:     page.PageCount,
		PageNumber:    page.PageNumber,
		PageSize:      page.PageSize,
		TotalRecords:  page.TotalRecords,
		NextPageToken: page.NextPageToken,
		Records:       meetings,
	}
	return dtoMeetingPage
}
