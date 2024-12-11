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

	scanner := bufio.NewScanner(f)
	switch *part {
	case 1:
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

func Part1(input *bufio.Scanner) int {
	var at, end point
	var blizzards []blizzard
	var w, h int
	var y int
	for input.Scan() {
		for x, c := range input.Text() {
			if y == 0 && c == '.' {
				at = point{x - 1, y - 1}
			}
			if c == '>' {
				blizzards = append(blizzards, blizzard{point{x - 1, y - 1}, point{1, 0}})
			}
			if c == 'v' {
				blizzards = append(blizzards, blizzard{point{x - 1, y - 1}, point{0, 1}})
			}
			if c == '<' {
				blizzards = append(blizzards, blizzard{point{x - 1, y - 1}, point{-1, 0}})
			}
			if c == '^' {
				blizzards = append(blizzards, blizzard{point{x - 1, y - 1}, point{0, -1}})
			}
		}
		y++
		w = len(input.Text()) - 2
	}
	h = y - 2

	end = point{w - 1, h}

	return Move(0, at, end, blizzards, w, h)
}

func Part2(input *bufio.Scanner) int {
	return 0
}

type point struct{ x, y int }

type blizzard struct {
	loc point
	d   point
}

func (b *blizzard) Move() {
	b.loc.x += b.d.x
	b.loc.y += b.d.y
}

var moves = []point{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
	{0, 0},
}

func gcd(x, y int) int {
	var small, gcd int
	if x > y {
		small = y
	} else {
		small = x
	}
	for i := 1; i <= small; i++ {
		if (x%i == 0) && (y%i == 0) {
			gcd = i
		}
	}

	return gcd
}

func Move(minute int, at, end point, blizzards []blizzard, w, h int) int {
	queue := new(deque.Deque[state])
	queue.PushFront(state{minute, at})
	lcm := w * h / gcd(w, h)
	var seen = map[state]struct{}{}

	for queue.Len() > 0 {
		curr := queue.PopFront()
		time := curr.t + 1

		for _, d := range moves {
			n := point{curr.p.x + d.y, curr.p.y + d.y}

			if n == end {
				return time
			}

			if (n.x < 0 || n.y < 0 || n.x >= w || n.y >= h) && !(n == point{0, -1}) {
				continue
			}

			if (n != point{0, -1}) && Occupied(n, time, blizzards, w, h) {
				continue
			}

			key := state{t: time % lcm, p: point{n.x, n.y}}
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			queue.PushBack(state{time, n})
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Occupied(p point, t int, blizzards []blizzard, w, h int) bool {
	for _, b := range blizzards {
		bl := point{
			x: (p.x - b.d.x*t) % w,
			y: (p.y - b.d.y*t) % h,
		}
		if b.loc == bl {
			return true
		}
	}
	return false
}

type state struct {
	t int
	p point
}
