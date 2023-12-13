package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 114, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestNextInSequence(t *testing.T) {
	assert.Equal(t, 18, nextInSequence([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 28, nextInSequence([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 68, nextInSequence([]int{10, 13, 16, 21, 30, 45}))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 2, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPrevInSequence(t *testing.T) {
	assert.Equal(t, -3, prevInSequence([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 0, prevInSequence([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 5, prevInSequence([]int{10, 13, 16, 21, 30, 45}))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 2101499000, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 1089, Part2(input))
}
