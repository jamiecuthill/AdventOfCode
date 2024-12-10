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

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

func Part1(input *bufio.Scanner) int {
	return count(1, input)
}

func Part2(input *bufio.Scanner) int {
	return count(3, input)
}

func count(n int, input *bufio.Scanner) int {
	var topN = top{d: make([]int, n)}
	var cur int

	for input.Scan() {
		if input.Text() == "" {
			topN.take(cur)

			cur = 0
			continue
		}

		cals, _ := strconv.Atoi(input.Text())
		cur += cals
	}

	topN.take(cur)

	return topN.sum()
}

type top struct {
	d []int
}

func (t *top) take(new int) {
	for i, c := range t.d {
		if new > c {
			t.d[i] = new
			new = c
		}
	}
}

func (t top) sum() int {
	var total int
	for _, c := range t.d {
		total += c
	}
	return total
}
