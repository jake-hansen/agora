package dto

import (
	"strconv"
	"strings"
	"time"
)

// MeetingDuration is the duration for a meeting.
type MeetingDuration time.Duration

// Meeting represents a information about a scheduled meeting on a meeting platform.
type Meeting struct {
	ID          string          `json:"id,omitempty"`
	Title       string          `json:"title,omitempty" binding:"required"`
	StartTime   string          `json:"start_time,omitempty" binding:"required,valid meeting time"`
	Duration    MeetingDuration `json:"duration,omitempty" binding:"required"`
	Description string          `json:"description,omitempty" binding:"required"`
	JoinURL     string          `json:"join_url,omitempty"`
	StartURL    string          `json:"start_url,omitempty"`
}

// InstantMeeting is a meeting that is going to take place at the current moment.
type InstantMeeting struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// MarshalJSON converts a MeetingDuration to JSON format.
func (m *MeetingDuration) MarshalJSON() ([]byte, error) {
	minutes := time.Duration(*m) / time.Minute
	return []byte(strconv.Itoa(int(minutes))), nil
}

// UnmarshalJSON parses JSON and attempts to convert the data to a MeetingDuration.
func (m *MeetingDuration) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*m = MeetingDuration(dur)
	return nil
}

// MeetingPage represents a paged collection of Meetings.
type MeetingPage struct {
	PageCount     int        `json:"page_count,omitempty"`
	PageNumber    int        `json:"page_number,omitempty"`
	PageSize      int        `json:"page_size,omitempty"`
	TotalRecords  int        `json:"total_records,omitempty"`
	NextPageToken string     `json:"next_page_token,omitempty"`
	Records       []*Meeting `json:"meetings"`
}
