package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
const example2 = `T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), 14)
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	result := Part1(input)
	assert.Equal(t, result, 394)
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example2))), 9)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example1))), 34)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part2(input), 1277)
}
