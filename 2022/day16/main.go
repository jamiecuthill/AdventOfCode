package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

// type valve struct {
// 	name    string
// 	rate    int
// 	tunnels []string
// }

var valveParser = regexp.MustCompile(`Valve ([A-Z]+) has flow rate=(\d+)`)

type dist struct {
	d int
	v string
}

func Part1(input *bufio.Scanner) int {
	valves, tunnels := parse(input)

	var nonempty []string
	var dists = make(map[string]map[string]int)

	for valve := range valves {
		if valve != "AA" && valves[valve] == 0 {
			continue
		}

		if valve != "AA" {
			nonempty = append(nonempty, valve)
		}

		dists[valve] = map[string]int{valve: 0, "AA": 0}
		visited := map[string]struct{}{valve: {}}

		queue := new(deque.Deque[dist])
		queue.Grow(len(valves))
		queue.PushFront(dist{d: 0, v: valve})

		for queue.Len() > 0 {
			d := queue.PopFront()
			for _, neighbour := range tunnels[d.v] {
				if _, in := visited[neighbour]; in {
					continue
				}
				visited[neighbour] = struct{}{}
				if valves[neighbour] > 0 {
					dists[valve][neighbour] = d.d + 1
				}
				queue.PushBack(dist{d: d.d + 1, v: neighbour})
			}
		}

		delete(dists[valve], valve)
		if valve != "AA" {
			// why?
			delete(dists[valve], "AA")
		}
	}

	// fmt.Println(dists)
	// fmt.Println(nonempty)

	indicies := map[string]int{}
	for i, v := range nonempty {
		indicies[v] = i
	}

	s := &search{
		cache:    map[key]int{},
		dists:    dists,
		indicies: indicies,
		valves:   valves,
	}

	return s.dfs(30, "AA", 0)

	// var minute int
	// var location string = "AA"
	// var open []valve
	// var released int

	// for {
	// 	minute++
	// 	if minute > part1Len {
	// 		break
	// 	}

	// 	fmt.Println("Minute", minute)
	// 	fmt.Println("Currently at", location, valves[location].tunnels)
	// 	fmt.Println("Open valves", open)
	// 	for i := range open {
	// 		fmt.Printf("Releasing %d from valve %s\n", open[i].rate, open[i].name)
	// 		released += open[i].rate
	// 	}

	// 	open, location = openOrMove(valves, location, open, minute)
	// 	fmt.Println("")
	// }

	// return released
}

type key struct {
	t int
	v string
	m int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var cache = map[key]int{}

type search struct {
	cache    map[key]int
	dists    map[string]map[string]int
	indicies map[string]int
	valves   map[string]int
}

func (s *search) dfs(time int, valve string, bitmask int) int {
	k := key{t: time, v: valve, m: bitmask}
	if v, ok := s.cache[k]; ok {
		return v
	}

	var maxval int
	for neighbour := range s.dists[valve] {
		bit := 1 << s.indicies[neighbour]
		if bitmask&bit == bit {
			// fmt.Printf("bitmask ignoring neighbour %s of valve %s\n", neighbour, valve)
			continue
		}
		remtime := time - s.dists[valve][neighbour] - 1
		if remtime <= 0 {
			continue
		}
		maxval = max(maxval, s.dfs(remtime, neighbour, bitmask|bit)+s.valves[neighbour]*remtime)
	}
	cache[key{t: time, v: valve, m: bitmask}] = maxval
	return maxval
}

func Part2(input *bufio.Scanner) int {
	valves, tunnels := parse(input)

	var nonempty []string
	var dists = make(map[string]map[string]int)

	for valve := range valves {
		if valve != "AA" && valves[valve] == 0 {
			continue
		}

		if valve != "AA" {
			nonempty = append(nonempty, valve)
		}

		dists[valve] = map[string]int{valve: 0, "AA": 0}
		visited := map[string]struct{}{valve: {}}

		queue := new(deque.Deque[dist])
		queue.Grow(len(valves))
		queue.PushFront(dist{d: 0, v: valve})

		for queue.Len() > 0 {
			d := queue.PopFront()
			for _, neighbour := range tunnels[d.v] {
				if _, in := visited[neighbour]; in {
					continue
				}
				visited[neighbour] = struct{}{}
				if valves[neighbour] > 0 {
					dists[valve][neighbour] = d.d + 1
				}
				queue.PushBack(dist{d: d.d + 1, v: neighbour})
			}
		}

		delete(dists[valve], valve)
		if valve != "AA" {
			// why?
			delete(dists[valve], "AA")
		}
	}

	// fmt.Println(dists)
	// fmt.Println(nonempty)

	indicies := map[string]int{}
	for i, v := range nonempty {
		indicies[v] = i
	}

	s := &search{
		cache:    map[key]int{},
		dists:    dists,
		indicies: indicies,
		valves:   valves,
	}

	b := (1 << len(nonempty)) - 1
	// fmt.Printf("%b", b)
	var m int

	for i := 0; i < (b+1)/2; i++ {
		// fmt.Printf("%b - %b: %d\n", i, b^i, m)
		m = max(m, s.dfs(26, "AA", i)+s.dfs(26, "AA", b^i))
	}

	return m
}

// openOrMove returns list of open valves and new location
// func openOrMove(valves map[string]valve, location string, open []valve, minute int) ([]valve, string) {
// 	var suggestion struct {
// 		v    valve
// 		t    string
// 		dest string
// 		cost int
// 	} = struct {
// 		v    valve
// 		t    string
// 		dest string
// 		cost int
// 	}{v: valves[location], cost: (part1Len - minute) * valves[location].rate}

// 	if isOpen(open, location) {
// 		suggestion.cost = 0
// 	} else {
// 		fmt.Printf("Value of opening %s is %d\n", location, suggestion.cost)
// 	}

// 	for _, v := range valves {
// 		if v.name == location || isOpen(open, v.name) {
// 			continue
// 		}
// 		pth := path(valves, location, v.name)
// 		// val := pth.totalValue()
// 		val := (part1Len - minute - pth.len()) * v.rate

// 		// Give some value to that nodes children
// 		var max int
// 		for _, nxt := range v.tunnels {
// 			if !pth.includes(nxt) && !isOpen(open, nxt) {
// 				v := (part1Len - minute - pth.len() - 2) * valves[nxt].rate
// 				if v > max {
// 					max = v
// 				}
// 			}
// 		}
// 		val += max

// 		// For longer paths the value should diminish by the value of shorter
// 		// paths

// 		fmt.Printf("%s -> %s (%d) = %d (%s)\n", location, v.name, pth.len(), val, pth.nextMove())
// 		if val > suggestion.cost {
// 			suggestion.t = pth.nextMove()
// 			suggestion.cost = val
// 			suggestion.dest = v.name
// 		}
// 	}

// 	if suggestion.t == "" {
// 		if !isOpen(open, suggestion.v.name) {
// 			fmt.Printf("Opening valve %s\n", suggestion.v.name)
// 			open = append(open, suggestion.v)
// 		}
// 		return open, location
// 	}

// 	fmt.Printf("Move to %s (destination %s)\n", suggestion.t, suggestion.dest)
// 	return open, suggestion.t
// }

// type tunnel struct {
// 	valve
// 	previous *tunnel
// }

// func (t tunnel) nextMove() string {
// 	this := &t
// 	var out string
// 	for this.previous != nil {
// 		out = this.name
// 		this = this.previous
// 	}
// 	return out
// }

// func (t tunnel) totalValue() int {
// 	var out int
// 	var minutes = 1
// 	this := &t
// 	for this != nil {
// 		out += (minutes * this.valve.rate)
// 		minutes++
// 		this = this.previous
// 	}
// 	return out
// }

// func (t tunnel) len() int {
// 	var out int
// 	this := &t
// 	for this.previous != nil {
// 		out++
// 		this = this.previous
// 	}
// 	return out
// }

// func (t tunnel) includes(name string) bool {
// 	this := &t
// 	for this != nil {
// 		if this.name == name {
// 			return true
// 		}
// 		this = this.previous
// 	}
// 	return false
// }

// func path(valves map[string]valve, location, destination string) tunnel {
// 	at, ok := valves[location]
// 	if !ok {
// 		panic("location '" + location + "' not found")
// 	}
// 	var moves = []tunnel{{valve: at}}
// 	var visited = map[string]struct{}{location: {}}

// 	for len(moves) > 0 {
// 		curr := moves[0]
// 		moves = moves[1:]

// 		if curr.name == destination {
// 			return curr
// 		}

// 		for _, t := range curr.tunnels {
// 			if _, ok := visited[t]; ok {
// 				continue
// 			}

// 			prev := curr
// 			moves = append(moves, tunnel{valve: valves[t], previous: &prev})
// 			visited[t] = struct{}{}
// 		}
// 	}

// 	panic("No path found between " + location + " and " + destination)
// }

func parse(input *bufio.Scanner) (map[string]int, map[string][]string) {
	var valves = map[string]int{}
	var tunnels = map[string][]string{}

	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, ";")
		valveMatches := valveParser.FindSubmatch([]byte(parts[0]))
		rate, _ := strconv.Atoi(string(valveMatches[2]))

		rawT := strings.TrimPrefix(parts[1], " tunnels lead to valves ")
		rawT = strings.TrimPrefix(rawT, " tunnel leads to valve ")

		valves[string(valveMatches[1])] = rate
		tunnels[string(valveMatches[1])] = strings.Split(rawT, ", ")
	}

	return valves, tunnels
}

// func isOpen(open []valve, name string) bool {
// 	for i := range open {
// 		if open[i].name == name {
// 			return true
// 		}
// 	}
// 	return false
// }
