package zoom

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jake-hansen/agora/platforms/common"

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
	err := common.CreateMeeting("Zoom", z.Client, BaseURLV2+url, oauth, zoomMeeting, &meetingResponse, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	return zoomadapter.ZoomMeetingToDomainMeeting(meetingResponse), err
}

// GetMeetings retrieves all meetings
func (z *ZoomActions) GetMeetings(oauth domain.OAuthInfo, pageReq domain.PageRequest) (*domain.Page, error) {
	path := "/users/me/meetings"

	paginationFunc := func(url url.URL) url.URL {
		q := url.Query()
		if pageReq.PageSize != 0 {
			q.Add("page_size", strconv.Itoa(pageReq.PageSize))
		}
		if pageReq.RequestedPage != "" {
			q.Add("next_page_token", pageReq.RequestedPage)
		}
		returnURL, _ := url.Parse(url.String())
		returnURL.RawQuery = q.Encode()
		return *returnURL
	}

	var meetings zoomdomain.MeetingList
	err := common.GetMeetings("Zoom", z.Client, BaseURLV2+path, oauth, paginationFunc, &meetings, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return zoomadapter.ZoomMeetingListToDomainMeetingPage(meetings), nil
}

// GetMeeting retrieves a single meeting by ID
func (z *ZoomActions) GetMeeting(oauth domain.OAuthInfo, meetingID string) (*domain.Meeting, error) {
	reqURL := "/meetings/" + url.QueryEscape(meetingID)

	var meeting zoomdomain.Meeting
	err := common.GetMeeting("Zoom", z.Client, BaseURLV2+reqURL, oauth, meetingID, &meeting, http.StatusOK)
	if err != nil {
		return nil, err
	}

	return zoomadapter.ZoomMeetingToDomainMeeting(meeting), nil
}

func (z *ZoomActions) DeleteMeeting(oauth domain.OAuthInfo, meetingID string) error {
	reqURL := "/meetings/" + url.QueryEscape(meetingID)

	err := common.DeleteMeeting("Webex", z.Client, BaseURLV2+reqURL, oauth, nil, http.StatusNoContent, meetingID)
	return err
}

