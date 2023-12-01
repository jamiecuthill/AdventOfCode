package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
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
		fmt.Println(Part1(scanner, 2000000))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

type point struct {
	x int
	y int
}

func Part1(input *bufio.Scanner, y int) int {
	pairs := pairs(input)
	// world, from, to := createWorld(pairs)
	return count(pairs, y)
}

type span struct {
	lower int
	upper int
}

func Part2(input *bufio.Scanner) int {
	pairs := pairs(input)

	var loc point

	var ranges []span

	for y := 0; y <= 4000000; y++ {
		for _, p := range pairs {
			dy := abs(p.sensor.y - y)
			if dy > p.dist {
				continue
			}

			dx := p.dist - dy

			lx := p.sensor.x - dx
			ux := p.sensor.x + dx

			ranges = append(ranges, span{lower: lx, upper: ux})
		}

		// detect gap between ranges
		ranges = collapse(ranges)
		if len(ranges) > 1 {
			loc.x = ranges[0].upper + 1
			loc.y = y
			break
		}

		// clear
		ranges = ranges[:0]
	}

	return loc.x*4000000 + loc.y
}

func collapse(ranges []span) []span {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Sort(ByLower(ranges))

	for i := 1; i < len(ranges); i++ {
		if ranges[i-1].upper >= ranges[i].lower {
			ranges[i].lower = min(ranges[i-1].lower, ranges[i].lower)
			ranges[i].upper = max(ranges[i-1].upper, ranges[i].upper)
			ranges = append(ranges[:i-1], ranges[i:]...)
			i = 0
		}
	}

	return ranges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type ByLower []span

func (a ByLower) Len() int           { return len(a) }
func (a ByLower) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLower) Less(i, j int) bool { return a[i].lower < a[j].lower }

var matchSensor = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

type pair struct {
	sensor point
	beacon point
	dist   int
}

func pairs(input *bufio.Scanner) []pair {
	var pairs []pair
	for input.Scan() {
		matches := matchSensor.FindSubmatch(input.Bytes())
		sensorX, _ := strconv.Atoi(string(matches[1]))
		sensorY, _ := strconv.Atoi(string(matches[2]))
		beaconX, _ := strconv.Atoi(string(matches[3]))
		beaconY, _ := strconv.Atoi(string(matches[4]))

		s := point{x: sensorX, y: sensorY}
		b := point{x: beaconX, y: beaconY}

		pairs = append(pairs, pair{sensor: s, beacon: b, dist: abs(s.x-b.x) + abs(s.y-b.y)})

	}
	return pairs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func count(pairs []pair, y int) int {
	var ranges []span
	for _, p := range pairs {
		dy := abs(p.sensor.y - y)
		if dy > p.dist {
			continue
		}

		dx := p.dist - dy

		lx := p.sensor.x - dx
		ux := p.sensor.x + dx

		ranges = append(ranges, span{lower: lx, upper: ux})
	}

	ranges = collapse(ranges)

	return ranges[0].upper - ranges[0].lower
}
