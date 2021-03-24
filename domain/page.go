package domain

// Page represents a paginated response.
type Page struct {
	PageCount         int
	PageNumber        int
	PageSize          int
	TotalRecords      int
	NextPageToken     string
	PreviousPageToken string
	Records           []interface{}
}

// PageRequest represents a paginated request.
type PageRequest struct {
	PageSize      int
	RequestedPage string
}
