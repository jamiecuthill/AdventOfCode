package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
		visited := Solve(make([]point, 2), scanner)
		fmt.Println(visited)
	case 2:
		visited := Solve(make([]point, 10), scanner)
		fmt.Println(visited)
	}
}

func Solve(rope []point, scanner *bufio.Scanner) uint {
	var world = make(worldgrid, 0, maxcap)
	var ok bool
	var visited uint
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		direction := rune(line[0])
		distanceStr := string(line[2:])
		distance, _ := strconv.Atoi(distanceStr)

		for ; distance > 0; distance-- {
			rope[0] = move(direction, rope[0])
			for i := 1; i < len(rope); i++ {
				rope[i] = follow(rope[i-1], rope[i])
			}
			if world, ok = exists(world, rope[len(rope)-1].x, rope[len(rope)-1].y); !ok {
				visited++
			}
		}
	}
	return visited
}

func diff(a int32) int32 {
	if a == 0 {
		return 0
	}
	if a > 0 {
		return 1
	}
	return -1
}

func abs(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

// follow moves the tail given the head and tail positions
func follow(h, t point) point {
	if abs(t.x-h.x) > 1 || abs(t.y-h.y) > 1 {
		return point{x: t.x + diff(h.x-t.x), y: t.y + diff(h.y-t.y)}
	}
	return t
}

func move(direction rune, p point) point {
	switch direction {
	case 'U':
		p.y++
	case 'D':
		p.y--
	case 'L':
		p.x--
	case 'R':
		p.x++
	}
	return p
}

type point struct {
	x int32
	y int32
}

const gridsize = 8
const maxcap = 20

type worldgrid [][][4]uint64

func exists(world worldgrid, x, y int32) (worldgrid, bool) {
	d := 0
	if x < 0 {
		d++
		x = -x
	}
	if y < 0 {
		d += 2
		y = -y
	}

	xi := x >> 3
	for i := int32(len(world)); i < xi+1; i++ {
		world = append(world, make([][4]uint64, 0, maxcap))
	}

	yi := y >> 3
	for i := int32(len(world[xi])); i < yi+1; i++ {
		world[xi] = append(world[xi], [4]uint64{})
	}

	shift := uint((x&(gridsize-1))<<3 + (y & (gridsize - 1)))

	bitmask := uint64(1 << shift)

	// Position has already been visited
	if world[xi][yi][d]&bitmask > 0 {
		return world, true
	}

	// Register that we've visited this location
	world[xi][yi][d] |= bitmask
	return world, false
}
