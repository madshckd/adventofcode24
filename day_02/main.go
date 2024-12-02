package main

import (
	"bufio"
	"fmt"

	"github.com/madshckd/adventofcode24/helpers"
)

/* global variables */
var totalSaferRecords, levelDownSaferRecords int
var increasedTransition bool

/* part one : to find number of safer records,
each records, containing levels should either all
increase or decrease but gradually that is difference
should be atleast one and atmost three*/

func findSaferRecords(record []string) bool {
	var levelDiff int

	/* iterating through levels of a specific record */
	for count := range len(record) {
		/* if the record manages to iterate through all elements*/
		/* then it's a safer record */
		if (count + 1) == len(record) {
			totalSaferRecords++
		} else {
			levelDiff = (helpers.ToInt(record[count]) - helpers.ToInt(record[count+1]))

			/* this cancels out enormous jump between levels */
			if levelDiff < -3 || levelDiff > 3 || levelDiff == 0 {
				return false
			} else {
				/* following is to check, whether the level follows same increased  */
				/* or decreased transition, mixed won't be a safer record */
				if count == 0 {
					increasedTransition = isIncreasedTransition(levelDiff)
				} else {
					if increasedTransition != isIncreasedTransition(levelDiff) {
						return false
					}
				}
			}
		}
	}

	/* safer record */
	return true
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

/*
	this function is used to generate possible records

with one level popped out to find the safer records
with one bad level
*/
func makeOneLevelDown(record []string) [][]string {
	var levelDownRecords [][]string

	for index := range len(record) {
		if index == 0 {
			levelDownRecords = append(levelDownRecords, record[1:])
		} else if index == (len(record) - 1) {
			levelDownRecords = append(levelDownRecords, record[:index])
		} else {
			levelDownRecords = append(levelDownRecords,
				append(append([]string{}, record[:index]...),
					record[index+1:]...))
		}
	}

	return levelDownRecords
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

		/* if the record fails to attain the safe parameter */
		/* then generate records with one level popped out */
		/* recheck if any possbile record set is safe */
		if findSaferRecords(record) == false {
			levelDownRecords := makeOneLevelDown(record)

			/* iterating through possible one level down records  */
			/* for already failed record */
			for index := range len(levelDownRecords) {
				if findSaferRecords(levelDownRecords[index]) == true {
					/* count of safer records with one bad level */
					levelDownSaferRecords++
					break
				}
			}
		}
	}

	/* PART ONE */
	fmt.Println("SAFER RECORDS : ", (totalSaferRecords - levelDownSaferRecords))

	/* PARt TWO */
	fmt.Println("SAFER RECORDS WITH ONLY ONE BAD LEVEL : ", totalSaferRecords)
}
