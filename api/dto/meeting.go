package dto

import (
	"strings"
	"time"
)

type MeetingDuration time.Duration

type Meeting struct {
	Title       string `json:"title" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	Duration    MeetingDuration `json:"duration" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (m *MeetingDuration) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*m = MeetingDuration(dur)
	return nil
}

