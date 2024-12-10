package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 11, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 31, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 2192892, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 22962826, Part2(input))
}
