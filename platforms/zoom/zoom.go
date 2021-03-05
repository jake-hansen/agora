package zoom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom/zoomadapter"
	"io/ioutil"
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
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer func() error {
		closeErr := res.Body.Close()
		if err == nil {
			err = closeErr
		}
		return err
	}()

	fmt.Println(string(body))
	return meeting, nil
}




