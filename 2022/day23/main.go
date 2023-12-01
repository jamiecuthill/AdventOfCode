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

	scanner := bufio.NewScanner(f)
	switch *part {
	case 1:
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

func Part1(input *bufio.Scanner) int {
	var elves []point
	var y int
	for input.Scan() {
		line := input.Text()
		for x, c := range line {
			switch c {
			case '#':
				elves = append(elves, point{x, y})
			}
		}
		y++
	}

	for round := 1; round <= 10; round++ {
		var proposed []point
		for _, e := range elves {
			proposed = append(proposed, e.Evaluate(elves))
		}

		for i, new := range proposed {

			for {
				j := Exists(new, proposed, i)
				if j == -1 {
					break
				}
				proposed[j] = elves[j]
				proposed[i] = elves[i]
			}
		}

		elves = proposed
		// fmt.Println("Round", round)
		// print(elves)

		d := directions[0]
		directions = append(directions[1:], d)
	}

	min, max := minMax(elves)

	return (max.x-min.x+1)*(max.y-min.y+1) - len(elves)
}

func Part2(input *bufio.Scanner) int {
	var elves []point
	var y int
	for input.Scan() {
		line := input.Text()
		for x, c := range line {
			switch c {
			case '#':
				elves = append(elves, point{x, y})
			}
		}
		y++
	}

	var round int = 1
	for {
		var proposed []point

		var moved bool
		for _, e := range elves {
			p := e.Evaluate(elves)
			if p != e {
				moved = true
			}
			proposed = append(proposed, p)
		}

		if !moved {
			break
		}

		for i, new := range proposed {
			for {
				j := Exists(new, proposed, i)
				if j == -1 {
					break
				}
				proposed[j] = elves[j]
				proposed[i] = elves[i]
			}
		}

		elves = proposed
		fmt.Println("Round", round)
		print(elves)

		// shuffle direction
		d := directions[0]
		directions = append(directions[1:], d)
		round++
	}

	return round
}

func minMax(elves []point) (point, point) {
	var min, max point
	min.x, min.y = math.MaxInt, math.MaxInt
	for _, e := range elves {
		if e.x > max.x {
			max.x = e.x
		}
		if e.y > max.y {
			max.y = e.y
		}
		if e.x < min.x {
			min.x = e.x
		}
		if e.y < min.y {
			min.y = e.y
		}
	}
	return min, max
}

type point struct {
	x int
	y int
}

func (p point) Apply(d point) point {
	return point{p.x + d.x, p.y + d.y}
}

func (p point) Evaluate(elves []point) point {
	var neighbours byte

	for _, other := range elves {
		if other == p {
			continue
		}
		d := point{x: other.x - p.x, y: other.y - p.y}
		if abs(d.x) <= 1 && abs(d.y) <= 1 {
			var shift int = ((d.y + 1) * 3) + (d.x + 1)
			if shift == 8 {
				shift = 4
			}
			neighbours |= (1 << shift)
		}
	}

	if neighbours == 0 {
		return p
	}

	for _, i := range directions {
		if i.mask&^neighbours == i.mask {
			return p.Apply(i.d)
		}
	}

	// nothing?
	return p
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var directions = []struct {
	mask byte
	d    point
}{
	{0b00000111, point{0, -1}},
	{0b11010000, point{0, 1}},
	{0b01001001, point{-1, 0}},
	{0b00110100, point{1, 0}},
}

func Exists(new point, proposed []point, i int) int {
	for j := range proposed {
		if i == j {
			continue
		}
		if new == proposed[j] {
			return j
		}
	}
	return -1
}

func print(elves []point) {
	min, max := minMax(elves)

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if Exists(point{x, y}, elves, -1) > -1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
