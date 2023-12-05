package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 4361, Part1(bufio.NewScanner(strings.NewReader(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 467835, Part2(bufio.NewScanner(strings.NewReader(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 540131, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 86879020, Part2(input))
}

func TestInlineTouches(t *testing.T) {
	assert.Equal(t, 2816, Part2(bufio.NewScanner(strings.NewReader(`.....50..481.........=.....643...........@......%............*.......815......681..263....*........*...5.....256*11.377$....872.903.*.......`))))
}
