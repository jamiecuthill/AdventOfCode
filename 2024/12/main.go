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

type Coord struct {
	x, y int
}

func parse(input *bufio.Scanner) map[rune]map[Coord]struct{} {
	var world = make(map[rune]map[Coord]struct{})
	var y int
	for input.Scan() {
		line := input.Text()
		for x, c := range line {
			if world[c] == nil {
				world[c] = make(map[Coord]struct{})
			}
			world[c][Coord{x, y}] = struct{}{}
		}
		y++
	}
	return world
}

func Part1(input *bufio.Scanner) int {
	world := parse(input)

	var sum int
	for _, v := range world {
		sum += Price(v)
	}
	return sum
}

func Price(shape map[Coord]struct{}) int {
	var price int
	var visited = make(map[Coord]struct{})

	for p := range shape {
		var perimeter, area int
		var q = new(deque.Deque[Coord])

		if _, ok := visited[p]; ok {
			continue
		}

		q.PushBack(p)
		visited[p] = struct{}{}
		area++

		for q.Len() > 0 {
			u := q.PopFront()

			for _, d := range directions {
				n := Coord{u.x + d.x, u.y + d.y}
				if _, ok := shape[n]; !ok {
					perimeter++
					continue
				}

				if _, ok := visited[n]; !ok {
					q.PushBack(n)
					visited[n] = struct{}{}
					area++
				}
			}
		}

		price += area * perimeter
	}

	return price
}

func Part2(input *bufio.Scanner) int {
	world := parse(input)

	var sum int
	for _, v := range world {
		sum += PriceWithDiscount(v)
	}
	return sum
}

var directions = []Coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func PriceWithDiscount(cells map[Coord]struct{}) int {
	var price int
	var visited = make(map[Coord]struct{})

	for p := range cells {
		var area int
		var regionMap = make(map[Coord]struct{})
		var q = new(deque.Deque[Coord])

		// p is part of an existing region
		if _, ok := visited[p]; ok {
			continue
		}

		// Explore new region
		q.PushBack(p)
		visited[p] = struct{}{}
		regionMap[p] = struct{}{}
		area++

		for q.Len() > 0 {
			u := q.PopFront()

			for _, d := range directions {
				n := Coord{u.x + d.x, u.y + d.y}
				if _, ok := cells[n]; !ok {
					continue
				}

				if _, ok := visited[n]; !ok {
					q.PushBack(n)
					regionMap[n] = struct{}{}
					visited[n] = struct{}{}
					area++
				}
			}
		}

		// Count the corners in the region
		var corners int

		for p := range regionMap {
			var dx, dy = -1, 0

			for i := 0; i < 4; i++ {
				// ?O
				// OX
				if _, ok := regionMap[Coord{x: p.x + dx, y: p.y + dy}]; !ok {
					if _, ok := regionMap[Coord{x: p.x + dy, y: p.y - dx}]; !ok {
						corners++
					}
				}

				// OX
				// XX
				if _, ok := regionMap[Coord{x: p.x + dx, y: p.y + dy}]; ok {
					if _, ok := regionMap[Coord{x: p.x + dy, y: p.y - dx}]; ok {
						if _, ok := regionMap[Coord{x: p.x + dx + dy, y: p.y + dy - dx}]; !ok {
							corners++
						}
					}
				}

				dx, dy = dy, -dx
			}
		}

		price += area * corners
	}

	return price
}
