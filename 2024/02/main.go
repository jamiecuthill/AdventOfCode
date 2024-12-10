package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

const safeThreshold = 3

func Part1(input *bufio.Scanner) int {
	var sum int

	for input.Scan() {
		line := input.Text()

		var levels []int
		parts := strings.Split(line, " ")

		for i := range parts {
			l, _ := strconv.Atoi(parts[i])
			levels = append(levels, l)
		}

		if isSafe(levels) {
			sum++
		}
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int

	for input.Scan() {
		line := input.Text()

		var levels []int
		parts := strings.Split(line, " ")

		for i := range parts {
			l, _ := strconv.Atoi(parts[i])
			levels = append(levels, l)
		}

		if isSafe(levels) {
			sum++
		} else {
			// Expand by removing
			for i := 0; i < len(levels); i++ {
				b := copyWithout(levels, i)
				if isSafe(b) {
					sum++
					break
				}
			}
		}
	}

	return sum
}

func copyWithout(levels []int, i int) []int {
	b := make([]int, len(levels)-1)
	copy(b[:i], levels[:i])
	copy(b[i:], levels[i+1:])
	return b
}

func isSafe(levels []int) bool {
	var isSafe = true
	var increasing *bool
	var t = true
	var f = false
	for i := 1; i < len(levels); i++ {
		if levels[i-1]-levels[i] < 0 {
			if increasing == nil {
				increasing = &t
			} else {
				if !*increasing {
					return false
				}
			}
		} else {
			if increasing == nil {
				increasing = &f
			} else {
				if *increasing {
					return false
				}
			}
		}
		d := int(math.Abs(float64(levels[i-1] - levels[i])))
		if d > safeThreshold || d < 1 {
			return false
		}
	}
	return isSafe
}
