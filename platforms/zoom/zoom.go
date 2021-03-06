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
	"net/url"
	"strconv"
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

func (z *ZoomActions) createZoomRequest(httpMethod string, url string, jsonBody interface{}, oauth domain.OAuthInfo) (*http.Request, error) {
	buffer := bytes.NewBuffer(nil)
	if jsonBody != nil {
		requestBody, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(requestBody)
	}
	req, err := http.NewRequest(httpMethod, BaseURLV2 + url, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oauth.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (z *ZoomActions) closeBody(res *http.Response) error {
	err := res.Body.Close()
	return err
}

func (z *ZoomActions) CreateMeeting(oauth domain.OAuthInfo, meeting *domain.Meeting) (*domain.Meeting, error) {
	url := "/users/me/meetings"

	zoomMeeting := zoomadapter.DomainMeetingToZoomMeeting(*meeting)

	req, err := z.createZoomRequest(http.MethodPost, url, zoomMeeting, oauth)
	if err != nil {
		return nil, NewRequestCreationError(BaseURLV2 + url, err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, NewRequestExecutionError(BaseURLV2 + url, err)
	}
	defer z.closeBody(res)

	if res.StatusCode != http.StatusCreated {
		return nil, NewZoomAPIError("create meeting", res.StatusCode)
	}

	var meetingResponse zoomdomain.Meeting
	err = json.NewDecoder(res.Body).Decode(&meetingResponse)
	if err != nil {
		return nil, NewResponseDecodingError(url, err)
	}

	return zoomadapter.ZoomMeetingToDomainMeeting(meetingResponse), err
}

func (z *ZoomActions) GetMeetings(oauth domain.OAuthInfo) (*domain.Page, error) {
	url := "/users/me/meetings"

	req, err := z.createZoomRequest(http.MethodGet, url, nil, oauth)
	if err != nil {
		return nil, NewRequestCreationError(BaseURLV2 + url, err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, NewRequestExecutionError(BaseURLV2 + url, err)
	}
	defer z.closeBody(res)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve meetings from Zoom")
	}

	var meetings zoomdomain.MeetingList
	err = json.NewDecoder(res.Body).Decode(&meetings)
	if err != nil {
		return nil, NewResponseDecodingError(url, err)
	}

	return zoomadapter.ZoomMeetingListToDomainMeetingPage(meetings), nil
}

func (z *ZoomActions) GetMeeting(oauth domain.OAuthInfo, meetingID string) (*domain.Meeting, error) {
	reqURL := "/meetings/" + url.QueryEscape(meetingID)

	req, err := z.createZoomRequest(http.MethodGet, reqURL, nil, oauth)
	if err != nil {
		return nil, NewRequestCreationError(BaseURLV2 + reqURL, err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, NewRequestExecutionError(BaseURLV2 + reqURL, err)
	}
	defer z.closeBody(res)

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, NewNotFoundError("meeting", meetingID, "user", strconv.Itoa(int(oauth.UserID)))
		}
		return nil, NewZoomAPIError("retrieve meeting", res.StatusCode)
	}

	var meeting zoomdomain.Meeting
	err = json.NewDecoder(res.Body).Decode(&meeting)
	if err != nil {
		return nil, NewResponseDecodingError(reqURL, err)
	}

	return zoomadapter.ZoomMeetingToDomainMeeting(meeting), nil
}
