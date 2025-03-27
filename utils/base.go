package utils

type ResCode struct {
	OK       int
	EXIST    int
	NOTFOUND int
	ISE      int
	INVALID  int
}

var Code = ResCode{
	OK:       200,
	EXIST:    201,
	NOTFOUND: 404,
	ISE:      500,
	INVALID:  406,
}

type Pagination struct {
	CurrentPage int     `json:"currentPage"`
	NextPage    int     `json:"nextPage"`
	PrevPage    int     `json:"prevPage"`
	NoOfPage    float64 `json:"noOfPage"`
}

type ResponseData[T any] struct {
	StatusCode int        `json:"statusCode"`
	DevMessage T          `json:"devMessage"`
	Paginate   Pagination `json:"paginate"`
}
