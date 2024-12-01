package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func StringSplit(str string, splitElem string) []string {
	return strings.Split(
		strings.Trim(str, splitElem),
		splitElem,
	)
}

func ToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
	}

	return num
}
