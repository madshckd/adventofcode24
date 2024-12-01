package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

/*function to split strings based on specific split element*/
func StringSplit(str string, splitElem string) []string {
	return strings.Split(
		strings.Trim(str, splitElem),
		splitElem,
	)
}

/*function to convert string to integer*/
func ToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
	}

	return num
}
