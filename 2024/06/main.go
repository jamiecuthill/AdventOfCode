package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

func parse(input *bufio.Scanner) (maxX int, maxY int, obsticles map[[2]int]struct{}, g Guard) {
	obsticles = make(map[[2]int]struct{})
	var y int
	for input.Scan() {
		line := input.Text()
		for x, c := range line {
			if x > maxX {
				maxX = x
			}
			if c == '#' {
				obsticles[[2]int{x, y}] = struct{}{}
			} else if c == '^' {
				g = Guard{StartingPosition: Coord{x, y}}
			}
		}
		y++
	}

	maxY = y - 1

	return
}

func Part1(input *bufio.Scanner) int {
	var x, y, obsticles, g = parse(input)
	return g.Walk(x, y, obsticles, false)
}

func Part2(input *bufio.Scanner) int {
	var x, y, obsticles, g = parse(input)

	g.Walk(x, y, obsticles, false)

	travelledPath := g.Visited
	var count int

	for xi := 0; xi <= x; xi++ {
		for yi := 0; yi <= y; yi++ {
			if _, ok := travelledPath[xi+yi*y]; !ok {
				continue
			}

			obsticles[[2]int{xi, yi}] = struct{}{}
			if g.Walk(x, y, obsticles, true) == -1 {
				count++
			}
			delete(obsticles, [2]int{xi, yi})
		}
	}

	return count
}

type Guard struct {
	StartingPosition Coord
	Visited          map[int]struct{}
	Turns            map[[4]int]struct{}
}

func (g *Guard) Walk(maxX, maxY int, obsticles map[[2]int]struct{}, detectLoops bool) int {
	g.Visited = make(map[int]struct{})
	g.Turns = make(map[[4]int]struct{})

	var heading = Coord{0, -1}
	var position = g.StartingPosition
	g.Visited[position.X+position.Y*maxX] = struct{}{}

	for {
		position.Move(heading)

		if position.Outside(maxX, maxY) {
			break
		}

		// Have we moved into an obsticle?
		if _, blocked := obsticles[position.Key()]; blocked {
			if detectLoops {
				// Are we revisiting a previous turn?
				if _, ok := g.Turns[[4]int{position.X, position.Y, heading.X, heading.Y}]; ok {
					return -1
				}
				g.Turns[[4]int{position.X, position.Y, heading.X, heading.Y}] = struct{}{}
			}

			position.Backup(heading)
			heading.Turn()
			continue
		}

		if !detectLoops {
			g.Visited[position.X+position.Y*maxX] = struct{}{}
		}
	}
	return len(g.Visited)
}

type Coord struct {
	X, Y int
}

func (c *Coord) Move(heading Coord) {
	c.X += heading.X
	c.Y += heading.Y
}

func (c *Coord) Backup(heading Coord) {
	c.X -= heading.X
	c.Y -= heading.Y
}

func (c *Coord) Turn() {
	c.X, c.Y = -c.Y, c.X
}

func (c Coord) Key() [2]int {
	return [2]int{c.X, c.Y}
}

func (c Coord) Outside(maxX, maxY int) bool {
	return c.X < 0 || c.X > maxX || c.Y < 0 || c.Y > maxY
}
