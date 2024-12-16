package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `2333133121414131402`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), 1928)
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	result := Part1(input)
	assert.Equal(t, result, 6367087064415)
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, Part2(bufio.NewScanner(strings.NewReader(example1))), 2858)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	result := Part2(input)
	assert.Assert(t, result < 6413841924593)
	assert.Assert(t, result > 1183608194968)
	assert.Equal(t, result, 6390781891880)
}
