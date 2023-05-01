package domain

type SortDirectionEnum string

const (
	SortDirectionEnumDesc SortDirectionEnum = "desc"
	SortDirectionEnumAsc  SortDirectionEnum = "asc"
)

type PaginationResultType struct {
	StartIndex     int  `json:"startIndex"`
	EndIndex       int  `json:"endIndex"`
	HasNextPage    bool `json:"hasNextPage"`
	HasPreviusPage bool `json:"hasPreviusPage"`
}

type PaginationInputType struct {
	StartIndex    int               `json:"startIndex"`
	First         int               `json:"first"`
	OrderBy       string            `json:"orderBy"`
	SortDirection SortDirectionEnum `json:"sortDirection"`
}

type PaginationResponseType struct {
	Items    interface{}
	PageInfo *PaginationResultType
}
