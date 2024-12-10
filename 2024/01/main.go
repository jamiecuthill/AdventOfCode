package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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

func Part1(input *bufio.Scanner) int {
	var sum int

	left, right := parse(input)

	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int

	left, right := parse(input)

	var rightCounts = make(map[int]int)

	for _, v := range right {
		rightCounts[v]++
	}

	for _, v := range left {
		sum += v * rightCounts[v]
	}

	return sum
}

func parse(input *bufio.Scanner) ([]int, []int) {
	var left, right []int

	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}
