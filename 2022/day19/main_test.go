package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 33, Part1(bufio.NewScanner(strings.NewReader(example))))
}

// func BenchmarkPart1Example(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Part1(bufio.NewScanner(strings.NewReader(example)))
// 	}
// }

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 301, Part2(bufio.NewScanner(strings.NewReader(example))))
}

// func BenchmarkPart2Example(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Part2(bufio.NewScanner(strings.NewReader(example)))
// 	}
// }

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 223971851179174, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 3379022190351, Part2(input))
}
