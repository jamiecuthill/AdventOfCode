package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 64, Part1(bufio.NewScanner(strings.NewReader(example))))
}

// func BenchmarkPart1Example(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Part1(bufio.NewScanner(strings.NewReader(example)))
// 	}
// }

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 58, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Simple(t *testing.T) {
	assert.Equal(t, 10, Part2(bufio.NewScanner(strings.NewReader(`1,1,1
1,1,2`))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 3500, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	count := Part2(input)
	assert.Equal(t, 2048, count)
}
