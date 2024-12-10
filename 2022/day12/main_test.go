package main

import (
	"bufio"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	in := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	s := bufio.NewScanner(strings.NewReader(in))
	assert.Equal(t, 31, Part1(s))
}

func TestPart2(t *testing.T) {
	in := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	s := bufio.NewScanner(strings.NewReader(in))
	assert.Equal(t, 29, Part2(s))
}
