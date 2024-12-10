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

const l = 0
const r = 1

type node [2]string

func parse(input *bufio.Scanner) (string, map[string]node) {
	var path string
	var nodes = map[string]node{}

	for input.Scan() {
		line := input.Text()
		if path == "" {
			path = line
			continue
		}

		if line == "" {
			continue
		}

		n := line[0:3]
		l := line[7:10]
		r := line[12:15]

		nodes[n] = node{l, r}
	}
	return path, nodes
}

func Part1(input *bufio.Scanner) int {
	path, nodes := parse(input)

	node := nodes["AAA"]

	var i, dir int
	for {
		if path[i%len(path)] == 'L' {
			dir = l
		} else {
			dir = r
		}

		if node[dir] == "ZZZ" {
			return i + 1
		}

		node = nodes[node[dir]]

		i++
	}
}

func Part2(input *bufio.Scanner) int {
	path, nodes := parse(input)

	var collect []int

	for k := range nodes {
		if k[2] == 'A' {
			incrs := findIncrement(path, nodes, k)
			collect = append(collect, incrs)
		}
	}

	return LCM(collect[0], collect[1], collect[2:]...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func findIncrement(path string, nodes map[string]node, node string) int {
	var i, dir int

	for {
		if path[i%len(path)] == 'L' {
			dir = l
		} else {
			dir = r
		}

		node = nodes[node][dir]
		if node[2] == 'Z' {
			return i + 1
		}

		i++
	}
}
