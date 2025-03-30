package pagination

type PageInfo struct {
	CurrentPage int     `json:"currentPage"`
	NextPage    int     `json:"nextPage"`
	PrevPage    int     `json:"prevPage"`
	TotalPages  float64 `json:"totalPages"`
}

type PagingOptions struct {
	Page int `json:"page"`
	Rows int `json:"rows"`
}
type SortingOptions struct {
	OrderBy string `json:"orderBy"`
	SortBy  string `json:"sortBy"`
	GroupBy string `json:"groupBy"`
}

type QueryParams struct {
	PagingOptions
	SortingOptions
}
