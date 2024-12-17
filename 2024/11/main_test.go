package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `125 17`

func TestBlink(t *testing.T) {
	r := Blink(0, 1)
	assert.Equal(t, r, 1)

	r = Blink(1, 1)
	assert.Equal(t, r, 1)

	r = Blink(10, 1)
	assert.Equal(t, r, 2)

	r = Blink(1000, 2)
	assert.Equal(t, r, 3)
}

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), 55312)
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part1(input), 194482)
}

var result int

func BenchmarkPart1(b *testing.B) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b.ResetTimer()

	var r int
	for i := 0; i < b.N; i++ {
		input := bufio.NewScanner(f)
		r = Part1(input)
	}

	result = r
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, Part2(input), 232454623677743)
}
