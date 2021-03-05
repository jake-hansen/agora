package zoomadapter

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom/zoomdomain"
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
		Title:       meeting.Topic,
		StartTime:   parsedTime,
		Duration:    time.Duration(meeting.Duration) * time.Minute,
		Description: meeting.Agenda,
		JoinURL: 	 meeting.JoinURL,
		StartURL: 	 meeting.StartURL,
	}
	return domainMeeting
}
