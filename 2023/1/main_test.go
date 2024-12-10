package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 142, Part1(bufio.NewScanner(strings.NewReader(`1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 281, Part2(bufio.NewScanner(strings.NewReader(`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`))))
}

func TestLastWithSpelling(t *testing.T) {
	assert.Equal(t, "9", lastWithSpelling("two1nine"))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 54304, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 54418, Part2(input))
}
