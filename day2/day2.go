package main

import (
	"fmt"
	"richal.au/aoc2023/utils"
	"strconv"
	"strings"
)

type round = map[string]int
type game struct {
	id     int
	rounds []round
}

func parseId(unparsedId string) int {
	var parts = strings.Split(unparsedId, " ")
	if len(parts) != 2 {
		fmt.Errorf("found invalid game")
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Errorf("found invalid game: id is not int")
	}
	return id
}

func parseRounds(unparsedRounds string) []round {
	rounds := make([]round, 0)
	for _, r := range strings.Split(unparsedRounds, ";") {
		picks := map[string]int{}
		stripped := strings.TrimSpace(r)
		for _, b := range strings.Split(stripped, ",") {
			var parts = strings.Split(strings.TrimSpace(b), " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Errorf("found invalid game: count of balls is invalid")
			}
			picks[parts[1]] = count
		}
		rounds = append(rounds, picks)
	}
	return rounds
}

func parseGame(gameStr string) game {
	if !strings.HasPrefix(gameStr, "Game ") {
		fmt.Errorf("found invalid game: no Game prefix")
	}

	var colonSplit = strings.Split(gameStr, ":")
	if len(colonSplit) != 2 {
		fmt.Errorf("found invalid game: more than 1 colon")
	}
	if len(colonSplit) != 1 {
		fmt.Errorf("found invalid game")
	}
	id := parseId(colonSplit[0])
	rounds := parseRounds(colonSplit[1])
	return game{id, rounds}

}

func parseGames(unparsedGames []string) []game {
	games := make([]game, 0)
	for _, unparsedGame := range unparsedGames {
		var g = parseGame(unparsedGame)
		games = append(games, g)
	}
	return games
}

func partOne(games []string) int {
	const maxRed = 12
	const maxGreen = 13
	const maxBlue = 14

	var idSum = 0
	for _, game := range parseGames(games) {
		var validGame = true
		for _, round := range game.rounds {
			if round["red"] > maxRed {
				validGame = false
			}
			if round["green"] > maxGreen {
				validGame = false
			}
			if round["blue"] > maxBlue {
				validGame = false
			}
		}
		if validGame {
			idSum += game.id
		}
	}
	return idSum
}

func partTwo(games []string) int {
	var sumPower = 0
	for _, game := range games {
		var minRed = 0
		var minGreen = 0
		var minBlue = 0
		parsed := parseGame(game)
		for _, round := range parsed.rounds {
			minRed = max(minRed, round["red"])
			minGreen = max(minGreen, round["green"])
			minBlue = max(minBlue, round["blue"])
		}
		sumPower += minRed * minGreen * minBlue
	}
	return sumPower
}

func main() {
	var games = utils.ReadIntoSliceOfStrings("day2.input")

	p1 := partOne(games)
	fmt.Println(p1)
	p2 := partTwo(games)
	fmt.Println(p2)
}
