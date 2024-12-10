package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 18, Part1(bufio.NewScanner(strings.NewReader(example))))
}

// func BenchmarkPart1Example(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Part1(bufio.NewScanner(strings.NewReader(example)))
// 	}
// }

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 20, Part2(bufio.NewScanner(strings.NewReader(example))))
}

// func BenchmarkPart2Example(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Part2(bufio.NewScanner(strings.NewReader(example)))
// 	}
// }

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, -1, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 3379022190351, Part2(input))
}
