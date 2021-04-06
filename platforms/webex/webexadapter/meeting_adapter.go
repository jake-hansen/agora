package webexadapter

import (
	"time"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/webex/webexdomain"
)

func DomainMeetingToWebexMeeting(meeting domain.Meeting) *webexdomain.Meeting {
	endTime := meeting.StartTime.Add(meeting.Duration)

	webex := &webexdomain.Meeting{
		Title:     meeting.Title,
		Agenda:    meeting.Description,
		Start:     meeting.StartTime.Format(time.RFC3339),
		End:       endTime.Format(time.RFC3339),
		SendEmail: false,
	}
	return webex
}

func WebexMeetingToDomainMeeting(meeting webexdomain.Meeting) *domain.Meeting {
	startTime, _ := time.Parse(time.RFC3339, meeting.Start)
	endTime, _ := time.Parse(time.RFC3339, meeting.End)
	duration := endTime.Sub(startTime)

	domainMeeting := &domain.Meeting{
		ID:          meeting.ID,
		Title:       meeting.Title,
		StartTime:   startTime,
		Duration:    duration,
		Description: meeting.Agenda,
		JoinURL:     meeting.WebLink,
		StartURL:    meeting.WebLink,
		Type:        domain.TypeScheduled,
	}
	return domainMeeting
}
