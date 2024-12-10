package main

import (
	"bufio"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 24, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 93, Part2(bufio.NewScanner(strings.NewReader(example))))
}
