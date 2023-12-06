package utils

import (
	"bufio"
	"os"
)

func ReadIntoSliceOfStrings(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func ReadIntoMap(filename string) [][]rune {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	input := make([][]rune, 0)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}
