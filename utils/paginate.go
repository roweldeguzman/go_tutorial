package utils

import (
	"api/struct/pagination"
	"math"
	"strconv"
	"strings"
)

func Paginate(rows int, page int, total int) pagination.PageInfo {

	d := float64(total) / float64(rows)
	pagination := pagination.PageInfo{
		CurrentPage: page,
		NextPage:    page + 1,
		PrevPage:    page - 1,
		TotalPages:  math.Ceil(d),
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
