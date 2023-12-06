package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 288, Part1(bufio.NewScanner(strings.NewReader(`Time:      7  15   30
Distance:  9  40  200`))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 71503, Part2(bufio.NewScanner(strings.NewReader(`Time:      7  15   30
Distance:  9  40  200`))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 608902, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 46173809, Part2(input))
}
