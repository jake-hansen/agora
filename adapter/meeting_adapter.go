package adapter

import (
	"time"

	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// ScheduledMeetingDTOToDomain converts the given scheduled Meeting from DTO representation to domain representation.
func ScheduledMeetingDTOToDomain(meeting *dto.Meeting) *domain.Meeting {
	domainMeeting := &domain.Meeting{
		Title:       meeting.Title,
		Duration:    time.Duration(meeting.Duration),
		Description: meeting.Description,
		Type:        domain.TypeScheduled,
	}

	if meeting.StartTime != "" {
		domainMeeting.StartTime, _ = time.Parse(time.RFC3339, meeting.StartTime)
	}

	return domainMeeting
}

// InstantMeetingDTOToDomain converts the given instant Meeting from DTO representation to domain representation.
func InstantMeetingDTOToDomain(meeting *dto.InstantMeeting) *domain.Meeting {
	domainMeeting := &domain.Meeting{
		Title:       meeting.Title,
		Description: meeting.Description,
		Type:        domain.TypeInstant,
	}

	return domainMeeting
}

// MeetingDomainToDTO converts the given Meeting from domain representation to DTO representation.
func MeetingDomainToDTO(meeting *domain.Meeting) *dto.Meeting {

	dtoMeeting := &dto.Meeting{
		ID:          meeting.ID,
		Title:       meeting.Title,
		Duration:    dto.MeetingDuration(meeting.Duration),
		Description: meeting.Description,
		JoinURL:     meeting.JoinURL,
		StartURL:    meeting.StartURL,
	}

	if !meeting.StartTime.IsZero() {
		dtoMeeting.StartTime = meeting.StartTime.Format(time.RFC3339)
	}

	return dtoMeeting
}

// MeetingPageDomainToDTO converts the given Page from domain representation to MeetingPage representation.
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
