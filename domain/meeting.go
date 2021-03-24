package domain

import "time"

const (
	TypeInstant   = 1
	TypeScheduled = 2
)

// Meeting represents a meeting on a MeetingPlatform.
type Meeting struct {
	ID          string
	Title       string
	StartTime   time.Time
	Duration    time.Duration
	Description string
	JoinURL     string
	StartURL    string
	Type        int
}
