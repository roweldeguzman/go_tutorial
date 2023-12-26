package utils

import (
	"math"
	"strconv"
	"strings"
)

func Paginate(rows int, page int, total int) map[string]interface{} {

	d := float64(total) / float64(rows)
	return map[string]interface{}{
		"currentPage": page,
		"nextPage":    page + 1,
		"prevPage":    page - 1,
		"noOfPage":    math.Ceil(d),
	}
}

func PagerTernary(val string, defVal int) int {
	if strings.Trim(val, " ") == "" {
		return defVal
	}
	newVal, isError := strconv.Atoi(val)
	if isError != nil {
		return defVal
	}

	return newVal
}
