package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = ``
const example2 = ``

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), 0)
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part1(input), 0)
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example2))), 0)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part2(input), 0)
}
