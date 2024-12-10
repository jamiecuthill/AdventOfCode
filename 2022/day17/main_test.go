package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 3068, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func BenchmarkPart1Example(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(bufio.NewScanner(strings.NewReader(example)))
	}
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 1514285714288, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 4886370, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 11374534948438, Part2(input))
}
