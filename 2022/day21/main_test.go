package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 152, Part1(bufio.NewScanner(strings.NewReader(example))))
}

// func BenchmarkPart1Example(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Part1(bufio.NewScanner(strings.NewReader(example)))
// 	}
// }

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 301, Part2(bufio.NewScanner(strings.NewReader(example))))
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

	assert.Equal(t, 223971851179174, Part1(input))
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
