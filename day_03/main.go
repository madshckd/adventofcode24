package main

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/madshckd/adventofcode24/helpers"
)

/* global variables */
var mulTotal int64
var mulInstruction = regexp.MustCompile(`mul\([0-9]+?,[0-9]+?\)`)

/* function to extract values of mul instruction */
/* perform multiplication and adding results to total */
func doMultiplication(mul []string) {
	for index := range len(mul) {
		/* removing mul () characters and only extracting values */
		valuesStr := mul[index][4:(len(mul[index]) - 1)]
		absValues := helpers.StringSplit(valuesStr, ",")
		/* performing multiplicatiion and adding it to the total */
		mulTotal += int64((helpers.ToInt(absValues[0])) * (helpers.ToInt(absValues[1])))
	}
}

/* main */
func main() {

	/* file operations */
	file := helpers.FileOpen("../inputs/day_03")
	defer file.Close()

	/* scanner */
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mul := mulInstruction.FindAllString(scanner.Text(), -1)
		doMultiplication(mul)
	}

	/* part one */
	fmt.Println(mulTotal)

}
