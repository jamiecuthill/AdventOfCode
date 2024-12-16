package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gammazero/deque"
)

var part = flag.Int("part", 1, "Run part 1 or part 2?")

func main() {
	flag.Parse()

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

const End = 9

type Coord struct {
	x, y  int
	Value int
}

func (c Coord) Neighbours(world map[Coord]int) []Coord {
	n := []Coord{}

	for _, d := range []Coord{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: -1}, {x: 0, y: 1}} {
		nc := Coord{x: c.x + d.x, y: c.y + d.y}
		if _, ok := world[nc]; ok && world[nc] == world[c]+1 {
			n = append(n, nc)
		}
	}

	return n
}

func parse(input *bufio.Scanner) (width, height int, world map[Coord]int, trailheads []Coord) {
	world = make(map[Coord]int)
	for input.Scan() {
		line := input.Text()
		for i := range line {
			if i > width {
				width = i
			}
			p := Coord{x: i, y: height}
			if line[i] == '0' {
				trailheads = append(trailheads, p)
			}
			world[p] = int(line[i]) - '0'
		}
		height++
	}
	height--
	return
}

func Part1(input *bufio.Scanner) int {
	_, _, world, trailheads := parse(input)

	var sum int
	for _, t := range trailheads {
		var visited = make(map[Coord]struct{})
		var q = new(deque.Deque[Coord])

		q.PushBack(t)
		visited[t] = struct{}{}

		for q.Len() > 0 {
			u := q.PopFront()

			if world[u] == End {
				t.Value++
			}

			for _, n := range u.Neighbours(world) {
				if _, ok := visited[n]; ok {
					continue
				}

				q.PushBack(n)
				visited[n] = struct{}{}
			}
		}

		sum += t.Value
	}

	return sum
}

type Trail struct {
	Head  Coord
	Tail  Coord
	Trail map[Coord]struct{}
}

func (t Trail) Next(world map[Coord]int) []Trail {
	var next []Trail

	for _, n := range t.Tail.Neighbours(world) {
		var newTrail = make(map[Coord]struct{})
		for k, v := range t.Trail {
			newTrail[k] = v
		}
		newTrail[n] = struct{}{}

		if world[n] == world[t.Tail]+1 {
			next = append(next, Trail{Head: t.Head, Tail: n, Trail: newTrail})
		}
	}

	return next
}

func Part2(input *bufio.Scanner) int {
	_, _, world, trailheads := parse(input)

	var sum int
	for _, t := range trailheads {
		var q = new(deque.Deque[Trail])

		q.PushBack(Trail{Head: t, Tail: t, Trail: map[Coord]struct{}{t: {}}})

		for q.Len() > 0 {
			u := q.PopFront()

			if world[u.Tail] == End {
				t.Value++
			}

			for _, n := range u.Next(world) {
				q.PushBack(n)
			}
		}

		sum += t.Value
	}

	return sum
}
