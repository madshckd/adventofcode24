package main

import (
	"bufio"
	"fmt"
	"slices"

	"github.com/madshckd/adventofcode24/helpers"
)

var (
	leftList      []int64
	rightList     []int64
	totalDistance int64
)

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

	fmt.Println(totalDistance)
}

func main() {
	file := helpers.FileOpen("../inputs/day_01")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := helpers.StringSplit(scanner.Text(), "   ")
		leftList = append(leftList, int64(helpers.ToInt(line[0])))
		rightList = append(rightList, int64(helpers.ToInt(line[1])))
	}

	findTotalDistance()
}
