package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 24000, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 45000, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 67450, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 199357, Part2(input))
}
