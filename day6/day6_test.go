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
			"Simple Almanac",
			"day6_test.input",
			288,
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
