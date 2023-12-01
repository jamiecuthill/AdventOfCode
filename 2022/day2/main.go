package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	_ int = iota
	rock
	paper
	scissors
)

const (
	loss int = 0
	draw     = 3
	win      = 6
)

var opponent = map[string]int{
	"A": rock,
	"B": paper,
	"C": scissors,
}

var guess = map[string]int{
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var target = map[string]int{
	"X": loss,
	"Y": draw,
	"Z": win,
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var total int
	var total2 int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		total += score(opponent[line[0]], guess[line[1]])
		total2 += scoreTarget(opponent[line[0]], target[line[1]])
	}

	fmt.Println("Total score: ", total)
	fmt.Println("Targetted total score: ", total2)
}

func Part1(input *bufio.Scanner) int {
	var total int
	var total2 int

	for input.Scan() {
		line := strings.Split(input.Text(), " ")
		total += score(opponent[line[0]], guess[line[1]])
		total2 += scoreTarget(opponent[line[0]], target[line[1]])
	}

	return total
}

func Part2(input *bufio.Scanner) int {
	var total int

	for input.Scan() {
		line := strings.Split(input.Text(), " ")
		total += scoreTarget(opponent[line[0]], target[line[1]])
	}

	return total
}

func score(opponent, me int) int {
	switch {
	case opponent == me:
		return me + draw
	case opponent == rock:
		switch me {
		case paper:
			return me + win
		}
	case opponent == paper:
		switch me {
		case scissors:
			return me + win
		}
	case opponent == scissors:
		switch me {
		case rock:
			return me + win
		}
	}
	return me + loss
}

func scoreTarget(opponent, outcome int) int {
	switch outcome {
	case win:
		return win + beat(opponent)
	case loss:
		return loss + lose(opponent)
	}
	return draw + opponent
}

func beat(move int) int {
	if move == scissors {
		return rock
	}
	return move + 1
}

func lose(move int) int {
	if move == rock {
		return scissors
	}
	return move - 1
}
