package main

import "testing"

func TestIsVisible(t *testing.T) {
	forest := [][]int{
		{3, 2, 6, 3, 3},
		{0, 5, 5, 3, 5},
		{3, 5, 3, 5, 3},
		{7, 1, 3, 4, 9},
		{3, 2, 2, 9, 0},
	}

	if !isVisible(tree{1, 1, 5}, forest) {
		t.Fatal("Tree should be visible")
	}

	if got := score(coordinate{2, 1}, forest); got != 4 {
		t.Fatalf("Tree score should be %d got %d", 4, got)
	}

	if got := score(coordinate{2, 3}, forest); got != 8 {
		t.Fatalf("Tree score should be %d got %d", 8, got)
	}

	if got := score(coordinate{0, 0}, forest); got != 0 {
		t.Fatalf("Tree score should be %d got %d", 0, got)
	}
}
