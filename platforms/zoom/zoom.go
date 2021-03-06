package zoom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom/zoomadapter"
	"github.com/jake-hansen/agora/platforms/zoom/zoomdomain"
	"net/http"
	"time"
)

const (
	BaseURLV2 = "https://api.zoom.us/v2"
)

type ZoomActions struct {
	Client	*http.Client
}

func NewZoom() *ZoomActions {
	return &ZoomActions{Client: &http.Client{
		Timeout:       time.Minute,
	}}
}

func (z *ZoomActions) CreateMeeting(oauth domain.OAuthInfo, meeting *domain.Meeting) (*domain.Meeting, error) {
	zoomMeeting := zoomadapter.DomainMeetingToZoomMeeting(*meeting)

	requestBody, err := json.Marshal(zoomMeeting)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, BaseURLV2 + "/users/me/meetings", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oauth.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	res, err := z.Client.Do(req)

	defer func() error {
		closeErr := res.Body.Close()
		if err == nil {
			err = closeErr
		}
		return err
	}()

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, errors.New("could not create meeting with Zoom")
	}

	var meetingResponse zoomdomain.Meeting
	err = json.NewDecoder(res.Body).Decode(&meetingResponse)
	if err != nil {
		return nil, err
	}

	return zoomadapter.ZoomMeetingToDomainMeeting(meetingResponse), err
}

func (z *ZoomActions) GetMeetings(oauth domain.OAuthInfo) (*domain.Page, error) {

	req, err := http.NewRequest(http.MethodGet, BaseURLV2 + "/users/me/meetings", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oauth.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	res, err := z.Client.Do(req)

	defer func() error {
		closeErr := res.Body.Close()
		if err == nil {
			err = closeErr
		}
		return err
	}()

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve meetings from Zoom")
	}

	var meetings zoomdomain.MeetingList
	err = json.NewDecoder(res.Body).Decode(&meetings)
	if err != nil {
		return nil, err
	}

	return zoomadapter.ZoomMeetingListToDomainMeetingPage(meetings), nil
}
