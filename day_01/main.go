package main

import (
	"bufio"
	"fmt"
	"slices"

	"github.com/madshckd/adventofcode24/helpers"
)

/*global variables */
var (
	leftList        []int64
	rightList       []int64
	totalDistance   int64
	totalSimilarity int64
)

/*part one : to find total distance between consecutive smallest of
numbers in both lists */

func findTotalDistance() {
	slices.Sort(leftList)
	slices.Sort(rightList)

	for index := 0; index < len(leftList); index++ {
		if leftList[index] < rightList[index] {
			totalDistance += rightList[index] - leftList[index]
		} else {
			totalDistance += leftList[index] - rightList[index]
		}
	}

	/*part one solution */
	fmt.Println("TOTAL DISTANCE : ", totalDistance)
}

/*part two : to find total similarity by finding out number of similar
counts for a number in left list and adding the multiplication of that
number with it's repeat count*/

func findTotalSimilarity() {
	for index := 0; index < len(leftList); index++ {
		similarCount := 0
		similarPoint := slices.Index(rightList, leftList[index])

		if similarPoint == -1 {
			continue
		}

		for leftList[index] == rightList[similarPoint] {
			similarCount++
			similarPoint++
		}

		totalSimilarity += leftList[index] * int64(similarCount)

	}

	/*part two solution */
	fmt.Println("TOTAL SIMILARITY : ", totalSimilarity)
}

/*main*/
func main() {
	/* Opens and closes input file */
	file := helpers.FileOpen("../inputs/day_01")
	defer file.Close()

	/*scanner to read the file */
	scanner := bufio.NewScanner(file)

	/* reads each line and makes two separate slices for left and right list */
	for scanner.Scan() {
		line := helpers.StringSplit(scanner.Text(), "   ")
		leftList = append(leftList, int64(helpers.ToInt(line[0])))
		rightList = append(rightList, int64(helpers.ToInt(line[1])))
	}

	/*PART ONE */
	findTotalDistance()

	/*PART TWO */
	findTotalSimilarity()
}
