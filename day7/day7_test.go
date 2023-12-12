package main

import (
	"richal.au/aoc2023/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{
			"Simple Hands",
			"day7_test.input",
			6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := partOne(utils.ReadIntoSliceOfStrings(tt.input))
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
			"Simple Hands",
			"day7_test.input",
			5905,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := partTwo(utils.ReadIntoSliceOfStrings(tt.input))
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
