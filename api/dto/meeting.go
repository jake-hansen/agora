package dto

import (
	"strconv"
	"strings"
	"time"
)

type MeetingDuration time.Duration

type Meeting struct {
	ID			string	`json:"id,omitempty"`
	Title       string `json:"title" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	Duration    MeetingDuration `json:"duration" binding:"required"`
	Description string `json:"description" binding:"required"`
	JoinURL		string	`json:"join_url,omitempty"`
	StartURL	string	`json:"start_url,omitempty"`
}

func (m *MeetingDuration) MarshalJSON() ([]byte, error) {
	minutes := time.Duration(*m) / time.Minute
	return []byte(strconv.Itoa(int(minutes))), nil
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

type MeetingPage struct {
	PageCount		int	`json:"page_count"`
	PageNumber		int	`json:"page_number"`
	PageSize		int	`json:"page_size"`
	TotalRecords	int	`json:"total_records"`
	Records			[]*Meeting	`json:"meetings"`
}
