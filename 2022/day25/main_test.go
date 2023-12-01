package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, "2=-1=0", Part1(bufio.NewScanner(strings.NewReader(example))))
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

	assert.Equal(t, "2-10==12-122-=1-1-22", Part1(input))
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

func TestToDec(t *testing.T) {
	tests := []struct {
		dec   int
		snafu string
	}{
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{7, "12"},
		{8, "2="},
		{9, "2-"},
		{10, "20"},
		{15, "1=0"},
		{20, "1-0"},
		{2022, "1=11-2"},
		{12345, "1-0---0"},
		{314159265, "1121-1110-1=0"},
		{1747, "1=-0-2"},
		{906, "12111"},
		{198, "2=0="},
		{11, "21"},
		{201, "2=01"},
		{31, "111"},
		{1257, "20012"},
		{32, "112"},
		{353, "1=-1="},
		{107, "1-12"},
		{7, "12"},
		{3, "1="},
		{37, "122"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.dec), func(t *testing.T) {
			assert.Equal(t, tt.dec, ToDec(tt.snafu))
		})
	}
}

func TestToSnafu(t *testing.T) {
	tests := []struct {
		dec   int
		snafu string
	}{
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{7, "12"},
		{8, "2="},
		{9, "2-"},
		{10, "20"},
		{15, "1=0"},
		{20, "1-0"},
		{2022, "1=11-2"},
		{12345, "1-0---0"},
		{314159265, "1121-1110-1=0"},
		{1747, "1=-0-2"},
		{906, "12111"},
		{198, "2=0="},
		{11, "21"},
		{201, "2=01"},
		{31, "111"},
		{1257, "20012"},
		{32, "112"},
		{353, "1=-1="},
		{107, "1-12"},
		{7, "12"},
		{3, "1="},
		{37, "122"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.dec), func(t *testing.T) {
			assert.Equal(t, tt.snafu, ToSnafu(tt.dec))
		})
	}
}
