package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example1 = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`
const example2 = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=07870, Y=06450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestMachineCost(t *testing.T) {
	// Not achieveable
	assert.Equal(t, MachineCost(P{1, 1}, P{1, 2}, P{3, 4}), uint64(0))

	// One move
	assert.Equal(t, MachineCost(P{1, 2}, P{1, 2}, P{3, 4}), uint64(3))
	assert.Equal(t, MachineCost(P{3, 4}, P{1, 2}, P{3, 4}), uint64(1))

	// Two moves
	assert.Equal(t, MachineCost(P{2, 4}, P{1, 2}, P{3, 4}), uint64(6))
	assert.Equal(t, MachineCost(P{4, 6}, P{6, 6}, P{2, 3}), uint64(2))

	// No intersection
	assert.Equal(t, MachineCost(P{10, 10}, P{5, 1}, P{6, 2}), uint64(0))
	assert.Equal(t, MachineCost(P{10, 10}, P{6, 2}, P{5, 1}), uint64(0))

	// Some random examples
	assert.Equal(t, MachineCost(P{12748, 12748}, P{26, 66}, P{67, 21}), uint64(0))
	assert.Equal(t, MachineCost(P{18641, 10279}, P{69, 23}, P{27, 71}), uint64(0))
}

func TestMachineCostFromInput(t *testing.T) {
	// Button A: X+46, Y+68
	// Button B: X+34, Y+14
	// Prize: X=11306, Y=10856
	assert.Equal(t, MachineCost(P{11306, 10856}, P{46, 68}, P{34, 14}), uint64(0))

	// Button A: X+34, Y+13
	// Button B: X+36, Y+49
	// Prize: X=268, Y=12318
	assert.Equal(t, MachineCost(P{268, 12318}, P{34, 13}, P{36, 49}), uint64(0))
}

func TestPart1Example(t *testing.T) {
	assert.Equal(t, Part1(bufio.NewScanner(strings.NewReader(example1))), uint64(480))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	result := Part1(input)
	assert.Equal(t, result, uint64(30973))
}

func TestPart2Example(t *testing.T) {
	result := Part2(bufio.NewScanner(strings.NewReader(example2)))
	assert.Equal(t, result, uint64(875_318_608_908))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	// Greater than 46_313_727_003_084
	assert.Equal(t, Part2(input), uint64(95_688_837_203_288))
}
