package dto

type PageRequest struct {
	PageSize		int
	RequestedPage	string
}

func NewPageReq(size int, requestedPage string) *PageRequest {
	return &PageRequest{
		PageSize:      size,
		RequestedPage: requestedPage,
	}
}