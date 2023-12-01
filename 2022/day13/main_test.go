package main

import (
	"bufio"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	in := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

	assert.Equal(t, 13, Part1(bufio.NewScanner(strings.NewReader(in))))
}

func TestCompare(t *testing.T) {
	ordered, _ := compare([]any{[]any{float64(2)}}, []any{float64(3)})
	assert.Equal(t, true, ordered)
}

func TestPart2(t *testing.T) {
	in := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

	assert.Equal(t, 140, Part2(bufio.NewScanner(strings.NewReader(in))))
}
