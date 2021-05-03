package webexadapter

import (
	"time"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/webex/webexdomain"
)

// DomainMeetingToWebexMeeting converts a Meeting from domain representation to Webex representation.
func DomainMeetingToWebexMeeting(meeting domain.Meeting) *webexdomain.Meeting {
	if meeting.Type == domain.TypeInstant {
		meeting.StartTime = time.Now().Add(time.Second * 30)
		meeting.Duration = time.Minute * 30

		if meeting.Description == "" {
			meeting.Description = "Instant meeting"
		}

		if meeting.Title == "" {
			meeting.Title = "Instant meeting"
		}
	}

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

// WebexMeetingToDomainMeeting converts a Meeting from Webex representation to domain representation.
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

// WebexMeetingListToDomainMeetingPage converts a MeetingList to a Page.
func WebexMeetingListToDomainMeetingPage(meetingList webexdomain.MeetingList) *domain.Page {
	size := len(meetingList.Items)
	var meetings []interface{}
	for _, meeting := range meetingList.Items {
		meetings = append(meetings, WebexMeetingToDomainMeeting(*meeting))
	}

	page := &domain.Page{
		PageCount:         0,
		PageNumber:        1,
		PageSize:          size,
		TotalRecords:      size,
		NextPageToken:     "",
		PreviousPageToken: "",
		Records:           meetings,
	}
	return page
}
