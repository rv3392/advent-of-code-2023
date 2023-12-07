package main

import (
	"fmt"
	"math"
	"richal.au/aoc2023/utils"
	"slices"
	"strconv"
	"strings"
)

func trimSpaceOverList(ls []string) []string {
	transformed := make([]string, len(ls))
	for p, l := range ls {
		transformed[p] = strings.TrimSpace(l)
	}
	return transformed
}

func removeEmptyElements(ls []string) []string {
	transformed := make([]string, 0)
	for _, l := range ls {
		if len(strings.TrimSpace(l)) != 0 {
			transformed = append(transformed, l)
		}
	}
	return transformed
}

type card struct {
	n       int
	winning []string
	yours   []string
}

func parseCard(unparsed string) card {
	var idAndNumbers = strings.Split(unparsed, ":")
	if len(idAndNumbers) != 2 {
		panic("invalid card")
	}
	var strId = removeEmptyElements(trimSpaceOverList(strings.Split(idAndNumbers[0], " ")))[1]
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic("invalid card: id is not an int")
	}
	var numbers = strings.Split(idAndNumbers[1], "|")
	if len(numbers) != 2 {
		panic("invalid card: winning and your numbers are weird")
	}
	var winningNumbers = strings.Split(strings.TrimSpace(numbers[0]), " ")
	winningNumbers = trimSpaceOverList(winningNumbers)
	winningNumbers = removeEmptyElements(winningNumbers)
	var yourNumbers = strings.Split(strings.TrimSpace(numbers[1]), " ")
	yourNumbers = trimSpaceOverList(yourNumbers)
	yourNumbers = removeEmptyElements(yourNumbers)
	return card{id, winningNumbers, yourNumbers}
}

func getNumWins(c card) int {
	var numWins = 0
	for _, yours := range c.yours {
		if slices.Contains(c.winning, yours) {
			numWins += 1
		}
	}
	return numWins
}

func getTotalPoints(cards []card) int {
	var points = 0
	for _, c := range cards {
		var wins = getNumWins(c)
		points += getPoints(wins)
	}
	return points
}

func getPoints(numWins int) int {
	if numWins == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(numWins-1)))
}

func partOne(unparsed []string) int {
	cards := make([]card, 0)
	for _, u := range unparsed {
		cards = append(cards, parseCard(u))
	}
	return getTotalPoints(cards)
}

func partTwo(unparsed []string) int {
	cards := make([]card, 0)
	for _, u := range unparsed {
		cards = append(cards, parseCard(u))
	}
	
	// Use a queue to store each card (both original and copy) that needs to be processed
	// For each card in the queue:
	//  1. Get the number of wins
	//  2. Add the copies of the cards that are won to the queue
	//  3. Increment the number of copies of cards (the answer to part two)
	q := make([]card, len(cards))
	copy(q, cards)
	var numCopies = 0
	for len(q) > 0 {
		// Pop
		var c = q[0] // Get the first element
		q = q[1:]    // Remove the first element
		var wins = getNumWins(c)
		for _, w := range cards[c.n : c.n+wins] {
			q = append(q, w)
		}
		numCopies++
	}
	return numCopies
}

func main() {
	cards := utils.ReadIntoSliceOfStrings("day4.input")
	totalPoints := partOne(cards)
	fmt.Println(totalPoints)
	numCopies := partTwo(cards)
	fmt.Println(numCopies)
}
