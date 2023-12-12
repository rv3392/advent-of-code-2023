package main

import (
	"fmt"
	"richal.au/aoc2023/utils"
	"strconv"
	"strings"
)

func parseListOfInts(line string, prefix string) []int {
	values := make([]int, 0)
	for _, value := range strings.Split(strings.TrimSpace(strings.TrimPrefix(line, prefix)), " ") {
		if strings.TrimSpace(value) == "" {
			continue
		}
		parsedValue, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			panic("value is not a valid int")
		}
		values = append(values, parsedValue)
	}
	return values
}

func parseSpacedInts(line string, prefix string) int {
	values := ""
	for _, value := range strings.Split(strings.TrimSpace(strings.TrimPrefix(line, prefix)), " ") {
		if strings.TrimSpace(value) == "" {
			continue
		}
		values += value
	}
	parsedValues, err := strconv.Atoi(values)
	if err != nil {
		panic("not a valid int")
	}
	return parsedValues
}

func numWaysToBreakRecord(time int, record int) int {
	var broken = 0
	for i := 0; i < time; i++ {
		var distance = (time - i) * i
		if distance > record {
			broken++
		}
	}
	return broken
}

func partOne(lines []string) int {
	var multipliedMarginOfError = 1
	var times = parseListOfInts(lines[0], "Time:")
	var distances = parseListOfInts(lines[1], "Distance:")

	for i, time := range times {
		multipliedMarginOfError *= numWaysToBreakRecord(time, distances[i])
	}

	fmt.Println(times)
	fmt.Println(distances)

	return multipliedMarginOfError
}

func partTwo(lines []string) int {
	var multipliedMarginOfError = 1
	var time = parseSpacedInts(lines[0], "Time:")
	var distance = parseSpacedInts(lines[1], "Distance:")
	multipliedMarginOfError *= numWaysToBreakRecord(time, distance)

	fmt.Println(time)
	fmt.Println(distance)

	return multipliedMarginOfError
}

func main() {
	lines := utils.ReadIntoSliceOfStrings("day6.input")
	multipliedMargin := partOne(lines)
	fmt.Println(multipliedMargin)
	multipliedMarginKerning := partTwo(lines)
	fmt.Println(multipliedMarginKerning)
}
