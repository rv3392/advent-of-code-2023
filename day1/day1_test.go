package main

import "testing"

func TestGetCalibrationValue(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"1abc2 gives 12", "1abc2", 12},
		{"pqr3stu8vwx gives 38", "pqr3stu8vwx", 38},
		{"a1b2c3d4e5f gives 15", "a1b2c3d4e5f", 15},
		{"treb7uchet gives 77", "treb7uchet", 77},
		{"two1nine gives 11", "two1nine", 11},
		{"abcone2threexyz gives 22", "abcone2threexyz", 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getCalibrationValue(tt.input, false)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestGetCalibrationValueSpeltDigits(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"1abc2 gives 12", "1abc2", 12},
		{"pqr3stu8vwx gives 38", "pqr3stu8vwx", 38},
		{"a1b2c3d4e5f gives 15", "a1b2c3d4e5f", 15},
		{"treb7uchet gives 77", "treb7uchet", 77},
		{"two1nine gives 29", "two1nine", 29},
		{"eightwothree gives 83", "eightwothree", 83},
		{"abcone2threexyz gives 13", "abcone2threexyz", 13},
		{"xtwone3four gives 24", "xtwone3four", 24},
		{"4nineeightseven2 gives 42", "4nineeightseven2", 42},
		{"zoneight234 gives 14", "zoneight234", 14},
		{"7pqrstsixteen gives 76", "7pqrstsixteen", 76},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getCalibrationValue(tt.input, true)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
