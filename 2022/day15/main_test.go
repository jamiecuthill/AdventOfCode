package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 26, Part1(bufio.NewScanner(strings.NewReader(example)), 10))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 56000011, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 4886370, Part1(input, 2000000))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 11374534948438, Part2(input))
}

func TestDete(t *testing.T) {
	ranges := collapse([]span{{-8, 12}, {6, 10}, {12, 14}, {14, 26}})
	assert.Equal(t, 1, len(ranges))
	assert.Equal(t, span{lower: -8, upper: 26}, ranges[0])

	ranges = collapse([]span{{-3, 3}, {2, 2}, {3, 13}, {11, 13}, {15, 25}, {15, 17}})
	assert.Equal(t, 2, len(ranges))
	assert.Equal(t, 14, ranges[0].upper+1)
}
