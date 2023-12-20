package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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
	uni := read(input)

	uni.expand(2)

	gs := uni.galaxies()

	var sum int
	var calculated = map[int]struct{}{}
	for _, a := range gs {
		for _, b := range gs {
			if a == b {
				continue
			}
			if _, ok := calculated[key(a, b)]; ok {
				continue
			}
			d := uni.distance(a, b)
			sum += d
			calculated[key(a, b)] = struct{}{}
		}
	}

	return sum
}

var factor = 1_000_000

func Part2(input *bufio.Scanner) int {
	uni := read(input)

	uni.expand(factor)

	gs := uni.galaxies()

	var sum int
	var calculated = map[int]struct{}{}
	for _, a := range gs {
		for _, b := range gs {
			if a == b {
				continue
			}
			if _, ok := calculated[key(a, b)]; ok {
				continue
			}
			d := uni.distance(a, b)
			sum += d
			calculated[key(a, b)] = struct{}{}
		}
	}

	return sum
}

func key(x, y int) int {
	if x < y {
		return x*1000 + y
	}
	return y*1000 + x
}

func read(input *bufio.Scanner) *universe {
	u := universe{}

	var count int = 1

	for input.Scan() {
		line := input.Text()
		row := make([]int, len(line))
		for i, c := range line {
			if c == '#' {
				row[i] = count
				count++
			}
		}
		u.grid = append(u.grid, row)
	}

	return &u
}

type universe struct {
	grid            [][]int
	expansionFactor int
	emptyY          []int
	emptyX          []int
	index           map[int][2]int
}

// expand mutates the universe to account for expansion that occurred during the observation of the input.
func (u *universe) expand(factor int) {
	u.expansionFactor = factor

	// check for empty rows
	for y := 0; y < len(u.grid); y++ {
		if isEmpty(u.grid[y]) {
			u.emptyY = append(u.emptyY, y)
		}
	}

	// check for empty columns
	for x := 0; x < len(u.grid[0]); x++ {
		isEmpty := true
		for y := 0; y < len(u.grid); y++ {
			if u.grid[y][x] != 0 {
				isEmpty = false
			}
		}
		if isEmpty {
			u.emptyX = append(u.emptyX, x)
		}
	}
}

func isEmpty(row []int) bool {
	for _, c := range row {
		if c != 0 {
			return false
		}
	}
	return true
}

func (u *universe) galaxies() []int {
	if u.index == nil {
		u.index = make(map[int][2]int)
	}
	var list []int
	for y, row := range u.grid {
		for x, c := range row {
			if c != 0 {
				u.index[c] = [2]int{x, y}
				list = append(list, c)
			}
		}
	}
	return list
}

func (u *universe) distance(f, t int) int {
	a := u.index[f]
	b := u.index[t]

	var dx, dy int
	for _, x := range u.emptyX {
		if between(a[0], x, b[0]) {
			dx += 1
		}
	}

	for _, y := range u.emptyY {
		if between(a[1], y, b[1]) {
			dy += 1
		}
	}

	x1, x2 := a[0], b[0]
	y1, y2 := a[1], b[1]

	return abs(x1-x2) + (dx * u.expansionFactor) - dx + abs(y1-y2) + (dy * u.expansionFactor) - dy
}

func between(a, b, c int) bool {
	if c < a {
		c, a = a, c
	}
	return a < b && b < c
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
