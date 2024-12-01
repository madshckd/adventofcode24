package helpers

import (
	"fmt"
	"os"
)

/*function to open up an input file*/
func FileOpen(fileName string) *os.File {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}
