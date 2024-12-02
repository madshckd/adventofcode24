package main

import (
	"bufio"
	"fmt"

	"github.com/madshckd/adventofcode24/helpers"
)

/* global variables */
var totalSaferRecords int
var increasedTransition bool

/* part one : to find number of safer records,
each records, containing levels should either all
increase or decrease but gradually that is difference
should be atleast one and atmost three*/

func findSaferRecords(record []string) {
	var levelDiff int

	/* iterating through levels of a specific record */
	for count := range len(record) {

		/* if the record manages to iterate through all elements*/
		/* then it's a safer record */
		if (count + 1) == len(record) {
			totalSaferRecords++
			return
		} else {
			levelDiff = (helpers.ToInt(record[count]) - helpers.ToInt(record[count+1]))

			/* this cancels out enormous jump between levels */
			if levelDiff < -3 || levelDiff > 3 || levelDiff == 0 {
				return
			} else {
				/* following is to check, whether the level follows same increased  */
				/* or decreased transition, mixed won't be a safer record */
				if count == 0 {
					increasedTransition = isIncreasedTransition(levelDiff)
				} else {
					if increasedTransition != isIncreasedTransition(levelDiff) {
						return
					}
				}
			}
		}
	}
}

/* this function is to check whether level differences
are either increasing or decreasing
< 0 means the subsequent value is greater (increasing)
> 0 means the subsequent value is lower (decreasing) */

func isIncreasedTransition(value int) bool {
	if value < 0 {
		return true
	} else {
		return false
	}
}

/* main */
func main() {
	/* file operation */
	file := helpers.FileOpen("../inputs/day_02")
	defer file.Close()

	/* reading a file */
	scanner := bufio.NewScanner(file)

	/* reading line by line */
	for scanner.Scan() {
		record := (helpers.StringSplit(scanner.Text(), " "))
		findSaferRecords(record)
	}

	/* PART ONE */
	fmt.Println("TOTAL SAFER RECORDS : ", totalSaferRecords)
}
