package helpers

import (
	"fmt"
	"strconv"
)

func Pagination(limit, page string) string {
	limitInt, _ := strconv.Atoi(limit)
	pageInt, _ := strconv.Atoi(page)

	formula := limitInt * (pageInt - 1)

	result := strconv.Itoa(formula)

	fmt.Println(limitInt)
	fmt.Println(pageInt)
	fmt.Println(formula)

	return result
}
