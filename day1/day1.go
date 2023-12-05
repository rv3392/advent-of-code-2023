package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isSpeltDigit(startPos int, line string) (bool, int, int) {
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for offset := 0; offset <= 5; offset++ {
		if len(line) < startPos+offset {
			break
		}
		var check = line[startPos : startPos+offset]
		for digitVal, digit := range digits {
			if check == digit {
				return true, digitVal + 1, offset - 1
			}
		}
	}
	return false, -1, -1
}

func getCalibrationValue(line string, parseSpeltDigits bool) int {
	digits := make([]int, 0)
	for pos := 0; pos < len(line); pos++ {
		c := rune(line[pos])
		if unicode.IsDigit(c) {
			digits = append(digits, int(c-'0'))
			continue
		}

		if !parseSpeltDigits {
			continue
		}

		isDigit, digit, offset := isSpeltDigit(pos, line)
		if !isDigit {
			continue
		}

		digits = append(digits, digit)
		pos += offset
	}

	if len(digits) == 1 {
		return digits[0]*10 + digits[0]
	}

	return digits[0]*10 + digits[len(digits)-1]
}

func partOne(calibrationDoc []string) {
	var total = 0
	for _, line := range calibrationDoc {
		total += getCalibrationValue(line, false)
	}
	fmt.Println(total)
}

func partTwo(calibrationDoc []string) {
	var total = 0
	for _, line := range calibrationDoc {
		total += getCalibrationValue(line, true)
	}
	fmt.Println(total)
}

func main() {
	f, err := os.Open("day1.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	partOne(input)
	partTwo(input)
}
