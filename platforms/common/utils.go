package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jake-hansen/agora/domain"
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

// CreateMeeting creates a meeting.
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

	if res.StatusCode != successCode {
		return NewAPIError(platformName, "create meeting", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return NewResponseDecodingError(endpoint, err)
	}

	return nil
}

// GetMeeting gets a meeting.
func GetMeeting(platformName string, client *http.Client, endpoint string, oauth domain.OAuthInfo, meetingID string, result interface{}, successCode int) error {
	req, err := CreateRequest(http.MethodGet, endpoint, nil, oauth)
	if err != nil {
		return NewRequestCreationError(endpoint, err)
	}

	res, err := client.Do(req)
	if err != nil {
		return NewRequestExecutionError(endpoint, err)
	}
	defer CloseBody(res)

	if res.StatusCode != successCode {
		if res.StatusCode == http.StatusNotFound {
			return NewNotFoundError("meeting", meetingID, "user", strconv.Itoa(int(oauth.UserID)))
		}
		return NewAPIError(platformName, "retrieve meeting", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return NewResponseDecodingError(endpoint, err)
	}

	return nil
}

// AddPagination is a function that takes a URL, adds information to it to get a specific page,
// and returns the URL
type AddPagination func(url url.URL) url.URL

// GetMeetings gets all meetings.
func GetMeetings(platformName string, client *http.Client, endpoint string, oauth domain.OAuthInfo, paginationFunc AddPagination, result interface{}, successCode int) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return NewRequestCreationError(endpoint, err)
	}

	paginatedURL := u
	if paginationFunc != nil {
		newURL := paginationFunc(*u)
		paginatedURL = &newURL
	}

	req, err := CreateRequest(http.MethodGet, paginatedURL.String(), nil, oauth)
	if err != nil {
		return NewRequestCreationError(paginatedURL.String(), err)
	}

	res, err := client.Do(req)
	if err != nil {
		return NewRequestExecutionError(paginatedURL.String(), err)
	}
	defer CloseBody(res)

	if res.StatusCode != successCode {
		return NewAPIError(platformName, "retrieve meetings", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return NewResponseDecodingError(paginatedURL.String(), err)
	}

	return nil
}

// DeleteMeeting deletes a meeting.
func DeleteMeeting(platformName string, client *http.Client, endpoint string, oauth domain.OAuthInfo, result interface{}, successCode int, meetingID string) error {
	req, err := CreateRequest(http.MethodDelete, endpoint, nil, oauth)
	if err != nil {
		return NewRequestCreationError(endpoint, err)
	}

	res, err := client.Do(req)
	if err != nil {
		return NewRequestExecutionError(endpoint, err)
	}
	defer CloseBody(res)

	if res.StatusCode != successCode {
		if res.StatusCode == http.StatusNotFound {
			return NewNotFoundError("meeting", meetingID, "user", strconv.Itoa(int(oauth.UserID)))
		}
		return NewAPIError(platformName, "delete meeting", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil && !errors.Is(err, io.EOF) {
		return NewResponseDecodingError(endpoint, err)
	}

	return nil
}
