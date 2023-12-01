package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

	scanner := bufio.NewScanner(f)
	switch *part {
	case 1:
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

type unit struct {
	occupied bool
	wall     bool
}

type point struct {
	x int
	y int
}

func Part1(input *bufio.Scanner) int {
	world, maxY := createWorld(input)

	return fill(world, point{x: 500}, maxY)
}

func Part2(input *bufio.Scanner) int {
	world, maxY := createWorld(input)

	return fillFloored(world, point{x: 500}, maxY, func() bool {
		return world[point{x: 500}].occupied
	})
}

func createWorld(input *bufio.Scanner) (map[point]unit, int) {
	var world = map[point]unit{}

	var maxY = 0
	for input.Scan() {
		points := mapSlice(strings.Split(input.Text(), " -> "), toPoint)

		for i := 1; i < len(points); i++ {
			from, to := points[i-1], points[i]
			step := point{
				x: direction(from.x, to.x),
				y: direction(from.y, to.y),
			}
			curr := from
			for {
				world[curr] = unit{occupied: true, wall: true}

				if curr == to {
					break
				}

				curr.x += step.x
				curr.y += step.y

				if curr.y > maxY {
					maxY = curr.y
				}
			}
		}
	}
	return world, maxY
}

func fill(world map[point]unit, start point, maxY int) int {
	var particles int
	var gravity = point{y: 1}
	var landed bool

	sand := start

	for {
		particles++

		for {
			sand, landed = move(sand, gravity, world)
			if landed {
				break
			}

			if sand.y > maxY {
				return particles - 1
			}
		}

		sand = start
	}
}

func move(sand point, gravity point, world map[point]unit) (point, bool) {
	next := point{y: sand.y + gravity.y, x: sand.x}
	if !world[next].occupied {
		return next, false
	}

	// if  it's occupied by sand look to the left
	next.x--
	if !world[next].occupied {
		return next, false
	}

	// if  it's occupied by sand look to the right
	next.x += 2
	if !world[next].occupied {
		return next, false
	}

	// sand landed
	world[sand] = unit{occupied: true}
	return sand, true

}

func fillFloored(world map[point]unit, start point, maxY int, goal func() bool) int {
	var particles int
	var gravity = point{y: 1}
	var landed bool

	sand := start

	for !goal() {
		particles++

		for {
			sand, landed = moveFloored(sand, gravity, world, maxY)
			if landed {
				break
			}
		}

		sand = start
	}

	return particles
}

func moveFloored(sand point, gravity point, world map[point]unit, maxY int) (point, bool) {
	next := point{y: sand.y + gravity.y, x: sand.x}

	// Reached floor
	if next.y == maxY+2 {
		world[sand] = unit{occupied: true}
		return sand, true
	}

	if !world[next].occupied {
		return next, false
	}

	// if  it's occupied by sand look to the left
	next.x--
	if !world[next].occupied {
		return next, false
	}

	// if  it's occupied by sand look to the right
	next.x += 2
	if !world[next].occupied {
		return next, false
	}

	// sand landed
	world[sand] = unit{occupied: true}
	return sand, true

}

func direction(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return -1
}

func mapSlice[A any, B any](in []A, mapfn func(A) B) []B {
	var out = make([]B, 0, len(in))
	for _, a := range in {
		out = append(out, mapfn(a))
	}
	return out
}

func toPoint(in string) point {
	parts := strings.Split(in, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return point{x: x, y: y}
}

// func render(world map[point]unit, start point) {
// 	var (
// 		minx, miny, maxx, maxy int
// 	)

// 	minx = math.MaxInt

// 	for k := range world {
// 		if k.x < minx {
// 			minx = k.x
// 		}
// 		if k.x > maxx {
// 			maxx = k.x
// 		}
// 		if k.y > maxy {
// 			maxy = k.y
// 		}
// 	}

// 	for y := miny; y <= maxy; y++ {
// 		for x := minx; x <= maxx; x++ {
// 			p := point{x: x, y: y}
// 			if p == start {
// 				fmt.Print("+")
// 			} else {
// 				if world[p].occupied && world[p].wall {
// 					fmt.Print("#")
// 				} else if world[p].occupied {
// 					fmt.Print("o")
// 				} else {
// 					fmt.Print(".")
// 				}
// 			}
// 		}
// 		fmt.Println("")
// 	}
// }
