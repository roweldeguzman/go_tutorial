package pagination

type PageParam struct {
	Page int
	Rows int
}
type SortParam struct {
	OrderBy string
	SortBy  string
	GroupBy string
}

type PaginationParam struct {
	PageParam
	SortParam
}
