package common

import "fmt"

var (
	ErrNotFound     = NotFoundError{}
	ErrReqCreation  = RequestCreationError{}
	ErrReqExecution = RequestExecutionError{}
	ErrResDecoding  = ResponseDecodingError{}
)

type NotFoundError struct {
	Resource    string
	ResourceID  string
	Requester   string
	RequesterID string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("the %s with id %s was not found for the %s with id %s", n.Resource, n.ResourceID, n.Requester, n.RequesterID)
}

func NewNotFoundError(resource string, resourceID string, requester string, requesterID string) NotFoundError {
	return NotFoundError{
		Resource:    resource,
		ResourceID:  resourceID,
		Requester:   requester,
		RequesterID: requesterID,
	}
}

func (n NotFoundError) Is(tgt error) bool {
	_, ok := tgt.(NotFoundError)
	return ok
}

type RequestCreationError struct {
	Err error
	URL string
}

func (r RequestCreationError) Error() string {
	return fmt.Sprintf("an error occurred while creating the request for url %s: %s", r.URL, r.Err.Error())
}

func (r RequestCreationError) Is(tgt error) bool {
	_, ok := tgt.(RequestCreationError)
	return ok
}

func NewRequestCreationError(url string, err error) RequestCreationError {
	return RequestCreationError{
		Err: err,
		URL: url,
	}
}

type RequestExecutionError struct {
	Err error
	URL string
}

func (r RequestExecutionError) Error() string {
	return fmt.Sprintf("an error occurred while executing the request for url %s: %s", r.URL, r.Err.Error())
}

func (r RequestExecutionError) Is(tgt error) bool {
	_, ok := tgt.(RequestExecutionError)
	return ok
}

func NewRequestExecutionError(url string, err error) RequestExecutionError {
	return RequestExecutionError{
		Err: err,
		URL: url,
	}
}

type APIError struct {
	Platform string
	Action string
	Code   int
}

func (z APIError) Error() string {
	return fmt.Sprintf("an error occurred while performing action '%s' with the %s API. http code %d", z.Action, z.Platform, z.Code)
}

func (z APIError) Is(tgt error) bool {
	_, ok := tgt.(APIError)
	return ok
}

func NewAPIError(platform string, action string, code int) APIError {
	return APIError{
		Platform: platform,
		Action: action,
		Code:   code,
	}
}

type ResponseDecodingError struct {
	Err error
	URL string
}

func (r ResponseDecodingError) Error() string {
	return fmt.Sprintf("an error occurred while decoding the response for url %s: %s", r.URL, r.Err.Error())
}

func (r ResponseDecodingError) Is(tgt error) bool {
	_, ok := tgt.(ResponseDecodingError)
	return ok
}

func NewResponseDecodingError(url string, err error) ResponseDecodingError {
	return ResponseDecodingError{
		Err: err,
		URL: url,
	}
}
