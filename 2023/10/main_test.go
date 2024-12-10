package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

const example2 = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 4, Part1(bufio.NewScanner(strings.NewReader(example1))))
	assert.Equal(t, 8, Part1(bufio.NewScanner(strings.NewReader(example2))))
}

const example3 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const example4 = `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`

const example5 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

const example6 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 4, Part2(bufio.NewScanner(strings.NewReader(example3))))
	assert.Equal(t, 4, Part2(bufio.NewScanner(strings.NewReader(example4))))
	assert.Equal(t, 8, Part2(bufio.NewScanner(strings.NewReader(example5))))
	assert.Equal(t, 10, Part2(bufio.NewScanner(strings.NewReader(example6))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 7012, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 395, Part2(input))
}
