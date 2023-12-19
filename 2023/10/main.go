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

func Part1(input *bufio.Scanner) int {
	field := []string{}
	start := [2]int{}
	for input.Scan() {
		line := input.Text()
		field = append(field, line)
		for x := 0; x < len(line); x++ {
			if line[x] == 'S' {
				start = [2]int{x, len(field) - 1}
			}
		}
	}

	pos := start
	var last int
	var len int

	for {
		for dir, next := range directions(pos[x], pos[y]) {
			if isBackwards(last, dir) {
				// That's where we came from, so we can't go there.
				continue
			}

			if !isValid(dir, rune(field[pos[y]][pos[x]])) {
				// The pipe doesn't go in that direction
				continue
			}

			if isConnected(dir, rune(field[next[y]][next[x]])) {
				// We found a connection
				pos = next
				last = dir
				len++
				break
			}
		}

		if pos == start {
			break
		}
	}

	return len / 2
}

const (
	x = iota
	y
)

const (
	up    = iota // 0
	right        // 1
	down         // 2
	left         // 3
)

func isBackwards(last, dir int) bool {
	switch dir {
	case up:
		return last == down
	case right:
		return last == left
	case down:
		return last == up
	case left:
		return last == right
	default:
		return false
	}
}

func directions(x, y int) [4][2]int {
	return [4][2]int{
		up:    {x, y - 1},
		right: {x + 1, y},
		down:  {x, y + 1},
		left:  {x - 1, y},
	}
}

func isValid(dir int, c rune) bool {
	if c == 'S' {
		return true
	}
	switch dir {
	case up:
		return c == '|' || c == 'J' || c == 'L'
	case right:
		return c == 'L' || c == '-' || c == 'F'
	case down:
		return c == '|' || c == 'F' || c == '7'
	case left:
		return c == '7' || c == '-' || c == 'J'
	default:
		return false
	}
}

func isConnected(dir int, to rune) bool {
	if to == 'S' {
		return true
	}
	switch dir {
	case down:
		return to == '|' || to == 'J' || to == 'L'
	case left:
		return to == 'F' || to == '-' || to == 'L'
	case up:
		return to == '|' || to == 'F' || to == '7'
	case right:
		return to == '7' || to == '-' || to == 'J'
	default:
		return false
	}
}

func Part2(input *bufio.Scanner) int {
	field := []string{}
	start := [2]int{}
	for input.Scan() {
		line := input.Text()
		field = append(field, line)
		for x := 0; x < len(line); x++ {
			if line[x] == 'S' {
				start = [2]int{x, len(field) - 1}
			}
		}
	}

	pos := start
	var last int

	poly := [][2]int{start}

	for {
		for dir, next := range directions(pos[x], pos[y]) {
			if isBackwards(last, dir) {
				// That's where we came from, so we can't go there.
				continue
			}

			if !isValid(dir, rune(field[pos[y]][pos[x]])) {
				// The pipe doesn't go in that direction
				continue
			}

			if next[y] < 0 || next[x] < 0 || next[y] >= len(field) || next[x] >= len(field[next[y]]) {
				// We're out of bounds
				continue
			}

			if isConnected(dir, rune(field[next[y]][next[x]])) {
				// We found a connection
				pos = next
				last = dir
				// Might need to simplify this polygon by collapsing straight lines
				poly = append(poly, next)
				break
			}
		}

		if pos == start {
			break
		}
	}

	var inside int
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if isInsidePolygon([2]int{x, y}, poly) {
				inside++
			}
		}
	}
	return inside
}

// isInsidePolygon returns true if the point is inside the polygon.
// Based upon Even-odd rule https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule
func isInsidePolygon(p [2]int, poly [][2]int) bool {
	j := len(poly) - 1
	var c bool

	for i := 0; i < len(poly); i++ {
		if p[x] == poly[i][x] && p[y] == poly[i][y] {
			// Exclude points on the corners
			return false
		}

		if (poly[i][y] > p[y]) != (poly[j][y] > p[y]) {
			slope := (p[x]-poly[i][x])*(poly[j][y]-poly[i][y]) - (poly[j][x]-poly[i][x])*(p[y]-poly[i][y])
			if slope == 0 {
				// Exclude points on the line
				return false
			}
			if (slope < 0) != (poly[j][1] < poly[i][1]) {
				c = !c
			}
		}

		j = i
	}

	return c
}
