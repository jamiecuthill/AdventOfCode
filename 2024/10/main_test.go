package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `0123
1234
8765
9876`
const example2 = `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
const example3 = `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
const example4 = `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`
const example5 = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), 1)
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example2))), 2)
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example3))), 4)
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example4))), 3)
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example5))), 36)
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part1(input), 461)
}

const example6 = `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`
const example7 = `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
const example8 = `012345
123456
234567
345678
4.6789
56789.`

func TestPart2Example(t *testing.T) {
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example6))), 3)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example7))), 13)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example8))), 227)
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example5))), 81)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part2(input), 875)
}
