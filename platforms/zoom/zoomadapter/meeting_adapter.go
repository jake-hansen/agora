package zoomadapter

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom/zoomdomain"
	"strconv"
	"time"
)

func DomainMeetingToZoomMeeting(meeting domain.Meeting) *zoomdomain.Meeting {
	durationMinutes := meeting.Duration.Round(time.Minute)
	gmt := meeting.StartTime.Format(time.RFC3339)

	zoom := &zoomdomain.Meeting{
		Topic:       meeting.Title,
		Type:        zoomdomain.TypeScheduled,
		StartTime:   gmt,
		Duration:    int(durationMinutes.Minutes()),
		Agenda:      meeting.Description,
	}
	return zoom
}

func ZoomMeetingToDomainMeeting(meeting zoomdomain.Meeting) *domain.Meeting {
	parsedTime, _ := time.Parse(time.RFC3339, meeting.StartTime)

	domainMeeting := &domain.Meeting {
		ID:			 strconv.Itoa(meeting.ID),
		Title:       meeting.Topic,
		StartTime:   parsedTime,
		Duration:    time.Duration(meeting.Duration) * time.Minute,
		Description: meeting.Agenda,
		JoinURL: 	 meeting.JoinURL,
		StartURL: 	 meeting.StartURL,
	}
	return domainMeeting
}

func ZoomMeetingListToDomainMeetingPage(meetingList zoomdomain.MeetingList) *domain.Page {
	var meetings []interface{}
	for _, meeting := range meetingList.Meetings {
		meetings = append(meetings, ZoomMeetingToDomainMeeting(*meeting))
	}

	page := &domain.Page{
		PageCount:    meetingList.PageCount,
		PageNumber:   meetingList.PageNumber,
		PageSize:     meetingList.PageSize,
		TotalRecords: meetingList.TotalRecords,
		Records:      meetings,
	}
	return page
}
