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

type Coord struct{ X, Y int }

func parse(input *bufio.Scanner) (map[byte][]Coord, int, int) {
	var antennas = make(map[byte][]Coord)

	var x, y int
	for input.Scan() {
		line := input.Text()
		for x = range line {
			if line[x] != '.' {
				if nil == antennas[line[x]] {
					antennas[line[x]] = make([]Coord, 0)
				}
				antennas[line[x]] = append(antennas[line[x]], Coord{x, y})
			}
		}
		y++
	}
	y--

	return antennas, x, y
}

func Part1(input *bufio.Scanner) int {
	antennas, x, y := parse(input)

	var antinodes = make(map[Coord]struct{})
	for k := range antennas {
		for i := 0; i < len(antennas[k]); i++ {
			for j := i; j < len(antennas[k]); j++ {
				if i == j {
					continue
				}

				var a = antennas[k][i]
				var b = antennas[k][j]

				var dx = a.X - b.X
				var dy = a.Y - b.Y

				for _, p := range []Coord{a, b} {
					xi := p.X + dx
					yi := p.Y + dy
					if xi >= 0 && xi <= x && yi >= 0 && yi <= y {
						antinodes[Coord{xi, yi}] = struct{}{}
					}
					dx = -dx
					dy = -dy
				}
			}
		}
	}

	return len(antinodes)
}

func Part2(input *bufio.Scanner) int {
	antennas, x, y := parse(input)

	var antinodes = make(map[Coord]struct{})
	for k := range antennas {
		for i := 0; i < len(antennas[k]); i++ {
			for j := i; j < len(antennas[k]); j++ {
				if i == j {
					continue
				}

				var a = antennas[k][i]
				var b = antennas[k][j]

				var dx = a.X - b.X
				var dy = a.Y - b.Y

				var i int
				for {
					xi := a.X + (i * dx)
					yi := a.Y + (i * dy)
					if xi < 0 || xi > x || yi < 0 || yi > y {
						break
					}
					antinodes[Coord{xi, yi}] = struct{}{}
					i++
				}
				dx = -dx
				dy = -dy
				i = 0
				for {
					i++
					xi := a.X + (i * dx)
					yi := a.Y + (i * dy)
					if xi < 0 || xi > x || yi < 0 || yi > y {
						break
					}
					antinodes[Coord{xi, yi}] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}
