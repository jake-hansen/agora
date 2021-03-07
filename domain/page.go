package domain

type Page struct {
	PageCount		int
	PageNumber		int
	PageSize		int
	TotalRecords	int
	NextPageToken	string
	PreviousPageToken	string
	Records			[]interface{}
}

type PageRequest struct {
	PageSize		int
	RequestedPage	string
}
