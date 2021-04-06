package webex

import (
	"github.com/jake-hansen/agora/domain"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://webexapis.com/v1"
)

type WebexActions struct {
	Client *http.Client
}

func NewWebex() *WebexActions {
	return &WebexActions{Client: &http.Client{
		Timeout: time.Minute,
	}}
}

func (w WebexActions) CreateMeeting(oauth domain.OAuthInfo, meeting *domain.Meeting) (*domain.Meeting, error) {
	panic("implement me")
}

func (w WebexActions) GetMeetings(oauth domain.OAuthInfo, pageReq domain.PageRequest) (*domain.Page, error) {
	panic("implement me")
}

func (w WebexActions) GetMeeting(oauth domain.OAuthInfo, meetingID string) (*domain.Meeting, error) {
	panic("implement me")
}
