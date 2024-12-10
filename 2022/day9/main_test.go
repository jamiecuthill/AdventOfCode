package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestFollow(t *testing.T) {
	tail := follow(point{1, 2}, point{0, 0})
	if tail.x != 1 || tail.y != 1 {
		t.Fatalf("Expected 1,1 got %+v", tail)
	}

	tail = follow(point{0, 0}, point{1, 2})
	if tail.x != 0 || tail.y != 1 {
		t.Fatalf("Expected 0,1 got %+v", tail)
	}

	tail = follow(point{0, 1}, point{2, 0})
	if tail.x != 1 || tail.y != 1 {
		t.Fatalf("Expected 1,1 got %+v", tail)
	}

	tail = follow(point{2, 0}, point{0, 1})
	if tail.x != 1 || tail.y != 0 {
		t.Fatalf("Expected 1,0 got %+v", tail)
	}
}

func TestPart1(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	in := bufio.NewScanner(strings.NewReader(input))
	if got := Solve(make([]point, 2), in); got != 13 {
		t.Fatalf("Expected 13, got %d", got)
	}
}

func TestPart2(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	in := bufio.NewScanner(strings.NewReader(input))
	if got := Solve(make([]point, 10), in); got != 1 {
		t.Fatalf("Expected 1, got %d", got)
	}

	input = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	in = bufio.NewScanner(strings.NewReader(input))
	if got := Solve(make([]point, 10), in); got != 36 {
		t.Fatalf("Expected 36, got %d", got)
	}
}
