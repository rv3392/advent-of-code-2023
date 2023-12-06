package main

import (
	"fmt"
	"richal.au/aoc2023/utils"
	"strconv"
	"unicode"
)

func rangeContainsSymbol(x1 int, x2 int, y1 int, y2 int, schematic [][]rune) bool {
	lines := schematic[y1 : y2+1]
	for _, line := range lines {
		for _, v := range line[x1 : x2+1] {
			if !unicode.IsDigit(v) && v != '.' {
				return true
			}
		}
	}
	return false
}

func rangeContainsGear(x1 int, x2 int, y1 int, y2 int, schematic [][]rune) string {
	lines := schematic[y1 : y2+1]
	for py, line := range lines {
		for px, v := range line[x1 : x2+1] {
			if v == '*' {
				return strconv.Itoa(x1+px) + "," + strconv.Itoa(y1+py)
			}
		}
	}
	return ""
}

func partOne(schematic [][]rune) int {
	var totalPartNumbers = 0
	for py, line := range schematic {
		var curNum = ""
		var startx = 0
		for px, value := range line {
			if unicode.IsDigit(value) {
				if len(curNum) == 0 {
					startx = max(px-1, 0)
				}
				curNum += string(value)
			}

			if (!unicode.IsDigit(value) && len(curNum) != 0) || (px == len(line)-1 && len(curNum) != 0) {
				var endx = px
				var starty = max(py-1, 0)
				var endy = min(py+1, len(schematic)-1)
				if rangeContainsSymbol(startx, endx, starty, endy, schematic) {
					partNum, err := strconv.Atoi(curNum)
					if err != nil {
						fmt.Errorf("invalid part number found")
					}
					totalPartNumbers += partNum
				}
				curNum = ""
			}
		}
	}
	return totalPartNumbers
}

func partTwo(schematic [][]rune) int {
	foundGears := make(map[string]int)
	var totalGearRatio = 0
	for py, line := range schematic {
		var curNum = ""
		var startx = 0
		for px, value := range line {
			if unicode.IsDigit(value) {
				if len(curNum) == 0 {
					startx = max(px-1, 0)
				}
				curNum += string(value)
			}

			if (!unicode.IsDigit(value) && len(curNum) != 0) || (px == len(line)-1 && len(curNum) != 0) {
				var endx = px
				var starty = max(py-1, 0)
				var endy = min(py+1, len(schematic)-1)
				if gearLoc := rangeContainsGear(startx, endx, starty, endy, schematic); gearLoc != "" {
					partNum1, err := strconv.Atoi(curNum)
					if err != nil {
						fmt.Errorf("invalid part number found")
					}

					if partNum2, previouslySeen := foundGears[gearLoc]; previouslySeen {
						totalGearRatio += partNum1 * partNum2
						delete(foundGears, gearLoc)
					} else {
						foundGears[gearLoc] = partNum1
					}
				}
				curNum = ""
			}
		}
	}
	return totalGearRatio
}

func main() {
	schematic := utils.ReadIntoMap("day3.input")
	var totalPartNumbers = partOne(schematic)
	fmt.Println(totalPartNumbers)
	var totalGearRatio = partTwo(schematic)
	fmt.Println(totalGearRatio)
}
