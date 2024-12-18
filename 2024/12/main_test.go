package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `AAAA
BBCD
BBCC
EEEC`
const example2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
const example3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), 140)
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example2))), 772)
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example3))), 1930)
}

func TestPrice(t *testing.T) {
	assert.Equal(t, Price(nil), 0)
	assert.Equal(t, Price(make(map[Coord]struct{})), 0)
	assert.Equal(t, Price(toMap([]Coord{{0, 0}})), 4)
	assert.Equal(t, Price(toMap([]Coord{{0, 0}, {4, 4}})), 8)
	assert.Equal(t, Price(toMap([]Coord{{0, 0}, {0, 1}})), 12)
	assert.Equal(t, Price(toMap([]Coord{{0, 0}, {0, 1}, {0, 2}})), 24)
	assert.Equal(t, Price(toMap([]Coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}})), 32)
	assert.Equal(t, Price(toMap([]Coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}})), 40)
}

func toMap(s []Coord) map[Coord]struct{} {
	m := make(map[Coord]struct{})
	for _, c := range s {
		m[c] = struct{}{}
	}
	return m
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part1(input), 1400386)
}

const example4 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
const example5 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

func TestPart2Example(t *testing.T) {
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example1))), 80)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example4))), 236)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example5))), 368)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example2))), 436)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example3))), 1206)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part2(input), 851994)
}
