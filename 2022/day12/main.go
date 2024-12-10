package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var part = flag.Int("part", 1, "Run part 1 or part 2?")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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

func parseStartEnd(in rune) rune {
	if in == 'E' {
		return 'z'
	}
	if in == 'S' {
		return 'a'
	}
	return in
}

func Part1(input *bufio.Scanner) int {
	world, start, end := parse(input)

	solution := move(
		path{this: start},
		world[end.y][end.x].rune,
		make([]point, 0),
		world,
		func(next, prev rune) bool {
			return parseStartEnd(prev) >= parseStartEnd(next)-1
		})
	if solution == nil {
		return -1
	}

	var len int
	for solution.previous != nil {
		len++
		solution = solution.previous
	}
	return len
}

func Part2(input *bufio.Scanner) int {
	world, _, end := parse(input)

	solution := move(
		path{this: end},
		'a',
		make([]point, 0),
		world,
		func(next, prev rune) bool {
			return parseStartEnd(prev) <= parseStartEnd(next)+1
		})
	if solution == nil {
		return -1
	}

	var len int
	for solution.previous != nil {
		len++
		solution = solution.previous
	}
	return len
}

var directions = []point{
	{y: 1},
	{x: 1},
	{y: -1},
	{x: -1},
}

type path struct {
	this     point
	previous *path
}

func move(curr path, end rune, track []point, world [][]loc, valid func(rune, rune) bool) *path {
	var queue = []path{curr}
	world[curr.this.y][curr.this.x].visited = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if world[v.this.y][v.this.x].rune == end {
			return &v
		}

		for _, d := range directions {
			prev := v
			next := path{
				this: point{
					x: v.this.x + d.x,
					y: v.this.y + d.y,
				},
				previous: &prev,
			}

			// is valid grid movement
			if next.this.x < 0 ||
				next.this.x == len(world[prev.this.y]) ||
				next.this.y < 0 ||
				next.this.y == len(world) ||
				world[next.this.y][next.this.x].visited {
				continue
			}

			if !valid(world[next.this.y][next.this.x].rune, world[prev.this.y][prev.this.x].rune) {
				continue
			}

			queue = append(queue, next)
			world[next.this.y][next.this.x].visited = true
		}
	}

	return nil
}

type point struct {
	x int
	y int
}

type loc struct {
	rune
	visited bool
}

func parse(input *bufio.Scanner) (world [][]loc, start point, end point) {
	world = make([][]loc, 0)
	var y int
	for input.Scan() {
		line := input.Text()
		world = append(world, make([]loc, 0))
		for x, c := range line {
			char := rune(c)
			if char == 'S' {
				start = point{x: x, y: y}
				// char = 'a'
			}
			if char == 'E' {
				end = point{x: x, y: y}
				// char = 'z'
			}
			world[y] = append(world[y], loc{rune: char})

		}
		y++
	}
	return
}
