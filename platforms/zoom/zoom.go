package zoom

import (
	"encoding/json"
	"github.com/jake-hansen/agora/platforms/common"
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

func (z *ZoomActions) closeBody(res *http.Response) error {
	err := res.Body.Close()
	return err
}

// CreateMeeting creates a meeting
func (z *ZoomActions) CreateMeeting(oauth domain.OAuthInfo, meeting *domain.Meeting) (*domain.Meeting, error) {
	url := "/users/me/meetings"

	zoomMeeting := zoomadapter.DomainMeetingToZoomMeeting(*meeting)

	var meetingResponse zoomdomain.Meeting
	err := common.CreateMeeting("Zoom", z.Client, BaseURLV2+url, oauth, zoomMeeting, meetingResponse)
	if err != nil {
		return nil, err
	}

	return zoomadapter.ZoomMeetingToDomainMeeting(meetingResponse), err
}

// GetMeetings retrieves all meetings
func (z *ZoomActions) GetMeetings(oauth domain.OAuthInfo, pageReq domain.PageRequest) (*domain.Page, error) {
	path := "/users/me/meetings"
	u, err := url.Parse("/users/me/meetings")
	if err != nil {
		return nil, common.NewRequestCreationError(BaseURLV2+path, err)
	}

	q := u.Query()
	if pageReq.PageSize != 0 {
		q.Add("page_size", strconv.Itoa(pageReq.PageSize))
	}
	if pageReq.RequestedPage != "" {
		q.Add("next_page_token", pageReq.RequestedPage)
	}
	u.RawQuery = q.Encode()

	req, err := common.CreateRequest(http.MethodGet, BaseURLV2+u.String(), nil, oauth)
	if err != nil {
		return nil, common.NewRequestCreationError(BaseURLV2+u.String(), err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, common.NewRequestExecutionError(BaseURLV2+u.String(), err)
	}
	defer z.closeBody(res)

	if res.StatusCode != http.StatusOK {
		return nil, common.NewAPIError("Zoom", "retrieve meetings", res.StatusCode)
	}

	var meetings zoomdomain.MeetingList
	err = json.NewDecoder(res.Body).Decode(&meetings)
	if err != nil {
		return nil, common.NewResponseDecodingError(BaseURLV2+u.String(), err)
	}

	return zoomadapter.ZoomMeetingListToDomainMeetingPage(meetings), nil
}

// GetMeeting retrieves a single meeting by ID
func (z *ZoomActions) GetMeeting(oauth domain.OAuthInfo, meetingID string) (*domain.Meeting, error) {
	reqURL := "/meetings/" + url.QueryEscape(meetingID)

	req, err := common.CreateRequest(http.MethodGet, BaseURLV2+reqURL, nil, oauth)
	if err != nil {
		return nil, common.NewRequestCreationError(BaseURLV2+reqURL, err)
	}

	res, err := z.Client.Do(req)
	if err != nil {
		return nil, common.NewRequestExecutionError(BaseURLV2+reqURL, err)
	}
	defer z.closeBody(res)

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, common.NewNotFoundError("meeting", meetingID, "user", strconv.Itoa(int(oauth.UserID)))
		}
		return nil, common.NewAPIError("Zoom", "retrieve meeting", res.StatusCode)
	}

	var meeting zoomdomain.Meeting
	err = json.NewDecoder(res.Body).Decode(&meeting)
	if err != nil {
		return nil, common.NewResponseDecodingError(reqURL, err)
	}

	return zoomadapter.ZoomMeetingToDomainMeeting(meeting), nil
}
