package main

import (
	"fmt"
	"richal.au/aoc2023/utils"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	cards    string
	bid      int
	category HandType
}

func compareFaceCard(face1, face2 rune, wildcard bool) bool {
	var encoding = map[rune]int{'T': 0, 'J': 1, 'Q': 2, 'K': 3, 'A': 4}
	if wildcard {
		encoding['J'] = -1
	}
	return encoding[face1] > encoding[face2]
}

func compareFaceCardWithWildcards(face1, face2 rune) bool {
	if face1 == 'J' && unicode.IsDigit(face2) {
		return false
	} else if unicode.IsDigit(face1) && face2 == 'J' {
		return true
	}
	var encoding = map[rune]int{'T': 0, 'Q': 1, 'K': 2, 'A': 3}
	return encoding[face1] > encoding[face2]
}

func lessThan(a, b Hand, wildcard bool) bool {
	if a.category < b.category {
		return true
	} else if a.category > b.category {
		return false
	}

	for c := 0; c < len(a.cards); c++ {
		cardA := rune(a.cards[c])
		cardB := rune(b.cards[c])
		if unicode.IsDigit(cardA) && unicode.IsDigit(cardB) {
			if cardA != cardB {
				return cardA > cardB
			}
		} else if unicode.IsDigit(cardA) {
			if cardB == 'J' {
				return true
			}
			return false
		} else if unicode.IsDigit(cardB) {
			if cardA == 'J' {
				return false
			}
			return true
		} else {
			if cardA != cardB {
				return compareFaceCard(cardA, cardB, wildcard)
			}
		}
	}
	return false
}

func determineType(cards string, wildcard bool) HandType {
	counts := make(map[rune]int)
	for _, card := range cards {
		counts[card]++
	}
	var numFours = 0
	var numThrees = 0
	var numTwos = 0
	var numSingles = 0
	var numWildcards = 0
	for card, count := range counts {
		if card == 'J' && wildcard {
			numWildcards = count
		} else if count == 5 {
			return FiveOfAKind
		} else if count == 4 {
			numFours++
		} else if count == 3 {
			numThrees++
		} else if count == 2 {
			numTwos++
		} else if count == 1 {
			numSingles++
		}
	}

	for i := 0; i < numWildcards; i++ {
		if numFours != 0 {
			return FiveOfAKind
		} else if numThrees != 0 {
			numFours++
			numThrees--
		} else if numTwos != 0 {
			numThrees++
			numTwos--
		} else if numSingles != 0 {
			numTwos++
			numSingles--
		} else {
			numSingles++
		}
	}

	if numFours == 1 {
		return FourOfAKind
	} else if numThrees == 1 && numTwos == 1 {
		return FullHouse
	} else if numThrees == 1 {
		return ThreeOfAKind
	} else if numTwos == 2 {
		return TwoPair
	} else if numTwos == 1 {
		return OnePair
	}
	return HighCard
}

func parseHandsAndBids(handsAndBids []string, wildcard bool) (hands []Hand) {
	hands = make([]Hand, 0)

	for _, handAndBid := range handsAndBids {
		parts := strings.Split(handAndBid, " ")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("bid is not a valid int: " + parts[1])
		}
		handType := determineType(parts[0], wildcard)
		hand := Hand{parts[0], bid, handType}
		hands = append(hands, hand)
	}

	return hands
}

func partOne(handsAndBids []string) int {
	var totalWinnings = 0
	hands := parseHandsAndBids(handsAndBids, false)
	sort.Slice(hands, func(i, j int) bool {
		return lessThan(hands[j], hands[i], false)
	})
	for rank, hand := range hands {
		totalWinnings += hand.bid * (rank + 1)
	}
	return totalWinnings
}

func partTwo(handsAndBids []string) int {
	var totalWinnings = 0
	hands := parseHandsAndBids(handsAndBids, true)
	sort.Slice(hands, func(i, j int) bool {
		return lessThan(hands[j], hands[i], true)
	})
	fmt.Println(hands)
	for rank, hand := range hands {
		totalWinnings += hand.bid * (rank + 1)
	}
	return totalWinnings
}

func main() {
	handsAndBids := utils.ReadIntoSliceOfStrings("day7.input")
	totalWinnings := partOne(handsAndBids)
	fmt.Println(totalWinnings)
	totalWinningsWithWildcards := partTwo(handsAndBids)
	fmt.Println(totalWinningsWithWildcards)
}
