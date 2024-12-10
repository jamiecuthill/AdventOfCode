package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 21, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 525152, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestCalcPositions(t *testing.T) {
	assert.Equal(t, 1, CalcPositions("???.###", []int{1, 1, 3}))
	assert.Equal(t, 4, CalcPositions(".??..??...?##.", []int{1, 1, 3}))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 7857, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 28606137449920, Part2(input))
}
