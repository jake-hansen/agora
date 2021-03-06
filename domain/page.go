package domain

type Page struct {
	PageCount		int
	PageNumber		int
	PageSize		int
	TotalRecords	int
	Records			[]interface{}
}
