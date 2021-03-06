package zoom

import "fmt"

var (
	ErrNotFound = NotFoundError{}
)

type NotFoundError struct {
	Resource	string
	ResourceID	string
	Requester	string
	RequesterID string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("the %s with id %s was not found for the %s with id %s", n.Resource, n.ResourceID, n.Requester, n.RequesterID)
}

func NewNotFoundError(resource string, resourceID string, requester string, requesterID string) NotFoundError {
	return NotFoundError{
		Resource:   resource,
		ResourceID: resourceID,
		Requester:  requester,
		RequesterID: requesterID,
	}
}

func (n NotFoundError) Is(tgt error) bool {
	_, ok := tgt.(NotFoundError)
	return ok
}


