package main

import (
	"fmt"
	"math"
	"richal.au/aoc2023/utils"
	"strconv"
	"strings"
	"unicode"
)

type rangemap struct {
	dstStart int
	srcStart int
	num      int
}
type catmap = []rangemap
type almanac = []catmap

func (r rangemap) isInRange(value int) bool {
	if value < r.srcStart {
		return false
	}
	if value > r.srcStart+r.num {
		return false
	}
	return true
}

func (r rangemap) mapToDst(value int) int {
	var offset = value - r.srcStart
	return r.dstStart + offset
}

func addMapping(mappings *catmap, toAdd string) {
	parts := strings.Split(toAdd, " ")
	if len(parts) != 3 {
		panic("unexpected number of parts")
	}
	dstStart, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("unexpected destination range start")
	}
	srcStart, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("unexpected source range start")
	}
	num, err := strconv.Atoi(parts[2])
	if err != nil {
		panic("unexpected range size")
	}
	mapping := rangemap{dstStart, srcStart, num}
	*mappings = append(*mappings, mapping)
}

func parseSeedsToPlant(unparsed string) []int {
	seeds := make([]int, 0)
	values, hasPrefix := strings.CutPrefix(unparsed, "seeds: ")
	if !hasPrefix {
		panic("seeds string is not valid")
	}
	for _, u := range strings.Split(values, " ") {
		seed, err := strconv.Atoi(u)
		if err != nil {
			panic("seed is not a valid int")
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseAlmanac(unparsed []string) almanac {
	as := make(almanac, 0)
	for _, u := range unparsed {
		if len(u) == 0 {
			continue
		}
		if strings.HasSuffix(u, "map:") {
			a := make(catmap, 0)
			as = append(as, a)
		} else if unicode.IsDigit(rune(u[0])) {
			addMapping(&(as[len(as)-1]), u)
		}
	}
	return as
}

func traverseAlmanac(as almanac, seed int) int {
	var curValue = seed
	for _, c := range as {
		for _, r := range c {
			if !r.isInRange(curValue) {
				continue
			}
			curValue = r.mapToDst(curValue)
			break
		}
	}
	return curValue
}

func partOne(unparsed []string) int {
	ss := parseSeedsToPlant(unparsed[0])
	as := parseAlmanac(unparsed[1:])
	var minLoc = math.MaxInt32
	for _, s := range ss {
		var loc = traverseAlmanac(as, s)
		minLoc = min(loc, minLoc)
	}
	return minLoc
}

func main() {
	almanac := utils.ReadIntoSliceOfStrings("day5.input")
	var minLoc = partOne(almanac)
	fmt.Println(minLoc)
}
