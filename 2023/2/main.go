package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var part = flag.Int("part", 1, "Run part 1 or part 2?")

type colour string

func main() {
	flag.Parse()

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

func Part1(input *bufio.Scanner) int {
	var max = map[colour]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var sum, i int

	for input.Scan() {
		line := input.Text()
		if line == "" {
			continue
		}

		i++
		games := parseGames(line)

		var exceeded bool

	foo:
		for _, game := range games {
			exceeded = false
			for colour, val := range game {
				if val > max[colour] {
					exceeded = true
					break foo
				}
			}
		}

		if !exceeded {
			sum += i
		}
	}

	return sum
}

func parseGames(line string) []map[colour]int {
	lineParts := strings.Split(line, ": ")
	rawGames := strings.Split(lineParts[1], "; ")

	games := []map[colour]int{}

	for _, rawGame := range rawGames {
		rawDices := strings.Split(rawGame, ", ")

		game := map[colour]int{}

		for _, rawDice := range rawDices {
			diceParts := strings.Split(rawDice, " ")
			val, _ := strconv.Atoi(diceParts[0])
			game[colour(diceParts[1])] = val
		}

		games = append(games, game)
	}
	return games
}

func Part2(input *bufio.Scanner) int {
	var sum int

	for input.Scan() {
		line := input.Text()
		if line == "" {
			continue
		}

		games := parseGames(line)

		var max = map[colour]int{}
		for _, game := range games {
			for colour, val := range game {
				if val > max[colour] {
					max[colour] = val
				}
			}
		}
		sum += max["red"] * max["green"] * max["blue"]
	}

	return sum
}
