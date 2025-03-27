package utils

import (
	"math"
	"strconv"
	"strings"
)

func Paginate(rows int, page int, total int) Pagination {

	d := float64(total) / float64(rows)
	pagination := Pagination{
		CurrentPage: page,
		NextPage:    page + 1,
		PrevPage:    page - 1,
		NoOfPage:    math.Ceil(d),
	}
	return pagination
}

func Ternary(val string, defVal int) int {
	if strings.Trim(val, " ") == "" {
		return defVal
	}
	newVal, _ := strconv.Atoi(val)
	return newVal
}
