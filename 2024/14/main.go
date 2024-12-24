package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/gdamore/tcell/v2"
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
		fmt.Println(Part1(input, 101, 103))
	case 2:
		fmt.Println(Part2(input, 101, 103))
	}
}

type Robot struct {
	px, py int
	vx, vy int
}

func (r *Robot) Move(width, height int) {
	r.px += r.vx
	r.py += r.vy

	if r.px < 0 {
		r.px = width + r.px
	}

	if r.py < 0 {
		r.py = height + r.py
	}

	if r.px >= width {
		r.px = r.px - width
	}

	if r.py >= height {
		r.py = r.py - height
	}
}

func (r *Robot) Reverse(width, height int) {
	r.px -= r.vx
	r.py -= r.vy

	if r.px < 0 {
		r.px = width + r.px
	}

	if r.py < 0 {
		r.py = height + r.py
	}

	if r.px >= width {
		r.px = r.px - width
	}

	if r.py >= height {
		r.py = r.py - height
	}
}

func parse(input *bufio.Scanner) []Robot {
	var robots []Robot
	for input.Scan() {
		line := input.Text()

		var r Robot
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.px, &r.py, &r.vx, &r.vy)

		robots = append(robots, r)
	}

	return robots
}

func countQuadrants(robots []Robot, width, height int) [2][2]int {
	var quadrants = [2][2]int{{0, 0}, {0, 0}}
	for _, r := range robots {
		if r.px == width/2 || r.py == height/2 {
			continue
		}
		var x, y int
		if r.px < width/2 {
			x = 0
		} else {
			x = 1
		}
		if r.py < height/2 {
			y = 0
		} else {
			y = 1
		}
		quadrants[x][y]++
	}

	return quadrants
}

func Part1(input *bufio.Scanner, width, height int) int {
	var robots []Robot = parse(input)

	for i := 0; i < 100; i++ {
		for j := range robots {
			robots[j].Move(width, height)
		}
	}

	// count number of robots in each quadrant
	var quadrants = countQuadrants(robots, width, height)
	return quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1]
}

func Part2(input *bufio.Scanner, width, height int) int {
	var robots []Robot = parse(input)

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	screen.Clear()

	var updated = make(chan struct{})

	var i = new(atomic.Int64)
	go func() {
		for {
			time.Sleep(80 * time.Millisecond)
			for j := range robots {
				robots[j].Move(width, height)
			}
			i.Add(1)
			updated <- struct{}{}
		}
	}()

	quit := make(chan struct{})
	events := make(chan tcell.Event)
	go screen.ChannelEvents(events, quit)

evloop:
	for {
		select {
		case <-updated:
			Render(screen, robots, width, height)
		case ev := <-events:
			switch ev := ev.(type) {
			case *tcell.EventResize:
				screen.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyLeft {
					i.Add(-2)
					for j := range robots {
						robots[j].Reverse(width, height)
						robots[j].Reverse(width, height)
					}
					continue
				}
				if ev.Key() == tcell.KeyRight {
					continue
				}
				if ev.Key() == tcell.KeyEscape {
					close(quit)
					screen.Fini()
					break evloop
				}
			}
		}

	}

	return int(i.Load())
}

// Render will draw the robots on the screen
func Render(screen tcell.Screen, robots []Robot, width, height int) {
	screen.Clear()

	for i, r := range robots {
		botStyle := tcell.StyleDefault.Foreground(tcell.Color(i))
		screen.SetContent(r.px, r.py, '#', nil, botStyle)
	}

	screen.Show()
}
