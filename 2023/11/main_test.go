package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 374, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Example(t *testing.T) {
	factor = 10
	assert.Equal(t, 1030, Part2(bufio.NewScanner(strings.NewReader(example))))

	factor = 100
	assert.Equal(t, 8410, Part2(bufio.NewScanner(strings.NewReader(example))))

	factor = 1_000_000
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 10313550, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 611998089572, Part2(input))
}

func TestBetween(t *testing.T) {
	assert.Equal(t, true, between(1, 2, 3))
	assert.Equal(t, true, between(3, 2, 1))
}
