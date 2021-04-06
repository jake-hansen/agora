package webex

import (
	"net/http"
	"time"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/common"
	"github.com/jake-hansen/agora/platforms/webex/webexadapter"
	"github.com/jake-hansen/agora/platforms/webex/webexdomain"
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
	url := "/meetings"

	webexMeeting := webexadapter.DomainMeetingToWebexMeeting(*meeting)

	var meetingResponse webexdomain.Meeting
	err := common.CreateMeeting("Webex", w.Client, BaseURLV1+url, oauth, webexMeeting, &meetingResponse, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return webexadapter.WebexMeetingToDomainMeeting(meetingResponse), nil
}

func (w WebexActions) GetMeetings(oauth domain.OAuthInfo, pageReq domain.PageRequest) (*domain.Page, error) {
	panic("implement me")
}

func (w WebexActions) GetMeeting(oauth domain.OAuthInfo, meetingID string) (*domain.Meeting, error) {
	panic("implement me")
}
