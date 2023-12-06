package main

import (
	"reflect"
	"richal.au/aoc2023/utils"
	"testing"
)

func TestRangeContainsSymbol(t *testing.T) {
	var schematic = [][]rune{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '$', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	var ans = rangeContainsSymbol(0, 3, 0, 1, schematic)
	if ans != true {
		t.Errorf("got %v, want %v", ans, true)
	}
	ans = rangeContainsSymbol(1, 4, 1, 2, schematic)
	if ans != true {
		t.Errorf("got %v, want %v", ans, true)
	}
	ans = rangeContainsSymbol(5, 9, 1, 3, schematic)
	if ans != true {
		t.Errorf("got %v, want %v", ans, true)
	}
	ans = rangeContainsSymbol(4, 8, 0, 1, schematic)
	if ans != false {
		t.Errorf("got %v, want %v", ans, true)
	}

}

func TestReadIntoMap(t *testing.T) {
	ans := utils.ReadIntoMap("day3test.input")
	want := [][]rune{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '.', '6', '3', '3'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
	}
	if !reflect.DeepEqual(want, ans) {
		t.Errorf("got %v, want %v", ans, want)
	}
}

func TestPartOne(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{
			"Simple Schematic",
			"day3test.input",
			4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := partOne(utils.ReadIntoMap(tt.input))
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{
			"Simple Schematic",
			"day3test.input",
			467835,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := partTwo(utils.ReadIntoMap(tt.input))
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
