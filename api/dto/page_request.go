package dto

// PageRequest contains information about a page to retrieve.
type PageRequest struct {
	PageSize      int
	RequestedPage string
}

// NewPageReq returns a new PageRequest with the specified parameters.
func NewPageReq(size int, requestedPage string) *PageRequest {
	return &PageRequest{
		PageSize:      size,
		RequestedPage: requestedPage,
	}
}
