package webex

import (
	"net/http"
	"net/url"
	"time"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/common"
	"github.com/jake-hansen/agora/platforms/webex/webexadapter"
	"github.com/jake-hansen/agora/platforms/webex/webexdomain"
)

const (
	BaseURLV1 = "https://webexapis.com/v1"
)

// WebexActions contains actions that can be performed on the Webex API.
type WebexActions struct {
	Client *http.Client
}

// NewWebex returns a WebexActions configured with a default http.Client.
func NewWebex() *WebexActions {
	return &WebexActions{Client: &http.Client{
		Timeout: time.Minute,
	}}
}

// CreateMeeting creates a meeting.
func (w *WebexActions) CreateMeeting(oauth domain.OAuthInfo, meeting *domain.Meeting) (*domain.Meeting, error) {
	url := "/meetings"

	webexMeeting := webexadapter.DomainMeetingToWebexMeeting(*meeting)

	var meetingResponse webexdomain.Meeting
	err := common.CreateMeeting("Webex", w.Client, BaseURLV1+url, oauth, webexMeeting, &meetingResponse, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return webexadapter.WebexMeetingToDomainMeeting(meetingResponse), nil
}

// GetMeetings gets all meetings for a user.
func (w *WebexActions) GetMeetings(oauth domain.OAuthInfo, pageReq domain.PageRequest) (*domain.Page, error) {
	reqURL := "/meetings"

	var meetings webexdomain.MeetingList
	err := common.GetMeetings("Webex", w.Client, BaseURLV1+reqURL, oauth, nil, &meetings, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return webexadapter.WebexMeetingListToDomainMeetingPage(meetings), nil
}

// GetMeeting gets a meeting.
func (w *WebexActions) GetMeeting(oauth domain.OAuthInfo, meetingID string) (*domain.Meeting, error) {
	reqURL := "/meetings/" + url.QueryEscape(meetingID)

	var meeting webexdomain.Meeting
	err := common.GetMeeting("Webex", w.Client, BaseURLV1+reqURL, oauth, meetingID, &meeting, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return webexadapter.WebexMeetingToDomainMeeting(meeting), nil
}

// DeleteMeeting deletes a meeting.
func (w *WebexActions) DeleteMeeting(oauth domain.OAuthInfo, meetingID string) error {
	reqURL := "/meetings/" + url.QueryEscape(meetingID)

	err := common.DeleteMeeting("Webex", w.Client, BaseURLV1+reqURL, oauth, nil, http.StatusNoContent, meetingID)
	return err
}
