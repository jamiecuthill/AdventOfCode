package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), 3749)
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part1(input), 2299996598890)
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example1))), 11387)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part2(input), 362646859298554)
}
