package main

import (
	"reflect"
	"testing"
)

func TestParseRound(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  []map[string]int
	}{
		{"3 blue, 4 red", "3 blue, 4 red", []map[string]int{{"blue": 3, "red": 4}}},
		{"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			[]map[string]int{{"blue": 3, "red": 4}, {"red": 1, "green": 2, "blue": 6}, {"green": 2}}},
		{"20 green, 5 red; 12870 green, 15 blue, 3 red", "20 green, 5 red; 12870 green, 15 blue, 3 red",
			[]map[string]int{{"green": 20, "red": 5}, {"green": 12870, "blue": 15, "red": 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := parseRounds(tt.input)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestParseGame(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  game
	}{
		{"Game 1: 3 blue, 4 red", "Game 1: 3 blue, 4 red", game{1, []map[string]int{{"blue": 3, "red": 4}}}},
		{"Game 256: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 256: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			game{256, []map[string]int{{"blue": 3, "red": 4}, {"red": 1, "green": 2, "blue": 6}, {"green": 2}}}},
		// This is not parsed correctly.
		{"Game  10: 20 green, 5 red; 12870 green, 15 blue, 3 red", "Game  10: 20 green, 5 red; 12870 green, 15 blue, 3 red",
			game{0, []map[string]int{{"green": 20, "red": 5}, {"green": 12870, "blue": 15, "red": 3}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := parseGame(tt.input)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  int
	}{
		{"Simple Game", []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := partOne(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  int
	}{
		{"Simple Game", []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}, 2286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := partTwo(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
