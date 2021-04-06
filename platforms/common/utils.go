package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"net/http"
)

// CreateRequest is a helper function that creates a request to be sent to a platform API. This function appends
// the provided OAuth token to the request in the Authorization header.
func CreateRequest(httpMethod string, url string, jsonBody interface{}, oauth domain.OAuthInfo) (*http.Request, error) {
	buffer := bytes.NewBuffer(nil)
	if jsonBody != nil {
		requestBody, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(requestBody)
	}
	req, err := http.NewRequest(httpMethod, url, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oauth.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func CloseBody(res *http.Response) error {
	err := res.Body.Close()
	return err
}

// CreateMeeting creates a meeting
func CreateMeeting(platformName string, client *http.Client, endpoint string, oauth domain.OAuthInfo, meeting interface{}, result interface{}, successCode int) error {
	req, err := CreateRequest(http.MethodPost, endpoint, meeting, oauth)
	if err != nil {
		return NewRequestCreationError(endpoint, err)
	}

	res, err := client.Do(req)
	if err != nil {
		return NewRequestExecutionError(endpoint, err)
	}
	defer CloseBody(res)

	if res.StatusCode != successCode{
		return NewAPIError(platformName, "create meeting", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return NewResponseDecodingError(endpoint, err)
	}

	return nil
}
