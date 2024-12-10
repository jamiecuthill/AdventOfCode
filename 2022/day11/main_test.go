package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	got := monkeyBusinessLevel([]*monkey{
		{inspections: 1},
		{inspections: 2},
		{inspections: 3},
	})
	if got != 6 {
		t.Fatalf("Expected 6, got %d", got)
	}

	cap := 10
	monkeys := []*monkey{
		{
			items:  (&ring{d: make([]int, cap)}).append(79).append(98), //[]int{79, 98},
			op:     func(i int) int { return i * 19 },
			modulo: 23,
		},
		{
			items:  (&ring{d: make([]int, cap)}).append(54).append(65).append(75).append(74), //[]int{54, 65, 75, 74},
			op:     func(i int) int { return i + 6 },
			modulo: 19,
		},
		{
			items:  (&ring{d: make([]int, cap)}).append(79).append(60).append(97), //[]int{79, 60, 97},
			op:     func(i int) int { return i * i },
			modulo: 13,
		},
		{
			items:  (&ring{d: make([]int, cap)}).append(74), //[]int{74},
			op:     func(i int) int { return i + 3 },
			modulo: 17,
		},
	}

	monkeys[0].a = monkeys[2]
	monkeys[0].b = monkeys[3]

	monkeys[1].a = monkeys[2]
	monkeys[1].b = monkeys[0]

	monkeys[2].a = monkeys[1]
	monkeys[2].b = monkeys[3]

	monkeys[3].a = monkeys[0]
	monkeys[3].b = monkeys[1]

	got = Solve(monkeys, func(i int) int { return i / 3 })
	if got != 10605 {
		t.Fatalf("Expected 10605, got %d", got)
	}
}

func BenchmarkSolve2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cap := 10
		monkeys := []*monkey{
			{
				items:  (&ring{d: make([]int, cap)}).append(79).append(98), //[]int{79, 98},
				op:     func(i int) int { return i * 19 },
				modulo: 23,
			},
			{
				items:  (&ring{d: make([]int, cap)}).append(54).append(65).append(75).append(74), //[]int{54, 65, 75, 74},
				op:     func(i int) int { return i + 6 },
				modulo: 19,
			},
			{
				items:  (&ring{d: make([]int, cap)}).append(79).append(60).append(97), //[]int{79, 60, 97},
				op:     func(i int) int { return i * i },
				modulo: 13,
			},
			{
				items:  (&ring{d: make([]int, cap)}).append(74), //[]int{74},
				op:     func(i int) int { return i + 3 },
				modulo: 17,
			},
		}

		monkeys[0].a = monkeys[2]
		monkeys[0].b = monkeys[3]

		monkeys[1].a = monkeys[2]
		monkeys[1].b = monkeys[0]

		monkeys[2].a = monkeys[1]
		monkeys[2].b = monkeys[3]

		monkeys[3].a = monkeys[0]
		monkeys[3].b = monkeys[1]

		got := Solve2(monkeys, func(i int) int {
			if i < 96577 {
				return i
			}
			return i % 96577
		})
		_ = got
	}
}

func TestRing(t *testing.T) {
	// Can loop around the ring
	r := ring{d: make([]int, 10)}
	for i := 1; i <= 100; i++ {
		r.append(i)
		assert.Equal(t, i, r.pop())
	}

	// Does not panic when filled
	r = ring{d: make([]int, 10)}
	for i := 1; i <= 10; i++ {
		r.append(i)
		assert.Equal(t, i, r.len())
	}

	for i := 1; i <= 10; i++ {
		j := r.pop()
		assert.Equal(t, i, j)
	}
}
