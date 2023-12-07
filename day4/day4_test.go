package main

import (
	"testing"
)

func TestGetNumWins(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"4 Wins", "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 4},
		{"2 Wins", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"2 Wins", "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"1 Win", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"No wins", "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"No wins", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getNumWins(parseCard(tt.input))
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestGetPoints(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  int
	}{
		{"No Wins", 0, 0},
		{"1 Win", 1, 1},
		{"2 Wins", 2, 2},
		{"3 Wins", 3, 4},
		{"10 wins", 10, 512},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getPoints(tt.input)
			if ans != tt.want {
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
		{
			"Simple Cards",
			[]string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			13,
		},
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
		{
			"Simple Cards",
			[]string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			30,
		},
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
