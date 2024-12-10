package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 18, Part1(bufio.NewScanner(strings.NewReader(example1))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 9, Part2(bufio.NewScanner(strings.NewReader(example1))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	result := Part1(input)
	assert.Assert(t, result != 2500, "Adventofcode said no")
	assert.Equal(t, 2493, result)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 1890, Part2(input))
}
