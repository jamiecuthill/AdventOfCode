package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `1
2
-3
3
-2
0
4`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 3, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func BenchmarkPart1Example(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x = Part1(bufio.NewScanner(strings.NewReader(example)))
	}
	_ = x
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 1623178306, Part2(bufio.NewScanner(strings.NewReader(example))))
}

// func BenchmarkPart2Example(b *testing.B) {
// 	var x int
// 	for i := 0; i < b.N; i++ {
// 		x = Part2(bufio.NewScanner(strings.NewReader(example)))
// 	}
// 	_ = x
// }

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 8764, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	count := Part2(input)
	assert.Assert(t, count > 2034, "count=%v", count)
}
