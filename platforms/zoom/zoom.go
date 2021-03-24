package zoom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom/zoomadapter"
	"github.com/jake-hansen/agora/platforms/zoom/zoomdomain"
)

const (
	BaseURLV2 = "https://api.zoom.us/v2"
)

// ZoomActions contains the functions necessary to interact with the Zoom API.
type ZoomActions struct {
	Client *http.Client
}

// NewZoom returns a new ZoomActions configured with an http.Client configured with a timeout
// of one minute for requests.
func NewZoom() *ZoomActions {
	return &ZoomActions{Client: &http.Client{
		Timeout: time.Minute,
	}}
}

// createZoomRequest is a helper function that creates a request to be sent to Zoom. This function appends
// the provided OAuth token to the request in the necesseary headers.
func (z *ZoomActions) createZoomRequest(httpMethod string, url string, jsonBody interface{}, oauth domain.OAuthInfo) (*http.Request, error) {
	buffer := bytes.NewBuffer(nil)
	if jsonBody != nil {
		requestBody, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(requestBody)
	}
	req, err := http.NewRequest(httpMethod, BaseURLV2+url, buffer)
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

// CreateMeeting creates a meeting
func (z *ZoomActions) CreateMeeting(oauth domain.OAuthInfo, meeting *domain.Meeting) (*domain.Meeting, error) {
	url := "/users/me/meetings"

	zoomMeeting := zoomadapter.DomainMeetingToZoomMeeting(*meeting)

	req, err := z.createZoomRequest(http.MethodPost, url, zoomMeeting, oauth)
	if err != nil {
		return nil, NewRequestCreationError(BaseURLV2+url, err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, NewRequestExecutionError(BaseURLV2+url, err)
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

// GetMeetings retrieves all meetings
func (z *ZoomActions) GetMeetings(oauth domain.OAuthInfo, pageReq domain.PageRequest) (*domain.Page, error) {
	path := "/users/me/meetings"
	u, err := url.Parse("/users/me/meetings")
	if err != nil {
		return nil, NewRequestCreationError(BaseURLV2+path, err)
	}

	q := u.Query()
	if pageReq.PageSize != 0 {
		q.Add("page_size", strconv.Itoa(pageReq.PageSize))
	}
	if pageReq.RequestedPage != "" {
		q.Add("next_page_token", pageReq.RequestedPage)
	}
	u.RawQuery = q.Encode()

	req, err := z.createZoomRequest(http.MethodGet, u.String(), nil, oauth)
	if err != nil {
		return nil, NewRequestCreationError(BaseURLV2+u.String(), err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, NewRequestExecutionError(BaseURLV2+u.String(), err)
	}
	defer z.closeBody(res)

	if res.StatusCode != http.StatusOK {
		return nil, NewZoomAPIError("retrieve meetings", res.StatusCode)
	}

	var meetings zoomdomain.MeetingList
	err = json.NewDecoder(res.Body).Decode(&meetings)
	if err != nil {
		return nil, NewResponseDecodingError(BaseURLV2+u.String(), err)
	}

	return zoomadapter.ZoomMeetingListToDomainMeetingPage(meetings), nil
}

// GetMeeting retrieves a single meeting by ID
func (z *ZoomActions) GetMeeting(oauth domain.OAuthInfo, meetingID string) (*domain.Meeting, error) {
	reqURL := "/meetings/" + url.QueryEscape(meetingID)

	req, err := z.createZoomRequest(http.MethodGet, reqURL, nil, oauth)
	if err != nil {
		return nil, NewRequestCreationError(BaseURLV2+reqURL, err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, NewRequestExecutionError(BaseURLV2+reqURL, err)
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
