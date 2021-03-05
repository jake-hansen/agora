package zoomadapter

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom/zoomdomain"
	"strconv"
	"time"
)

func DomainMeetingToZoomMeeting(meeting domain.Meeting) *zoomdomain.Meeting {
	durationMinutes := meeting.Duration.Round(time.Minute)

	zoom := &zoomdomain.Meeting{
		Topic:       meeting.Title,
		Type:        zoomdomain.TypeScheduled,
		StartTime:   meeting.StartTime,
		Duration:    strconv.Itoa(int(durationMinutes.Minutes())),
		Agenda:      meeting.Description,
	}
	return zoom
}
