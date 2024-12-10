package main

import (
	"bufio"
	"flag"
	"fmt"
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

func Part1(input *bufio.Scanner) int {
	var sum int
	for input.Scan() {
		line := input.Text()
		numbers := strings.Split(line, " ")
		sum += nextInSequence(toInts(numbers))
	}
	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int
	for input.Scan() {
		line := input.Text()
		numbers := strings.Split(line, " ")
		sum += prevInSequence(toInts(numbers))
	}
	return sum
}

func toInts(numbers []string) []int {
	var ints = make([]int, len(numbers))
	for i, n := range numbers {
		ints[i], _ = strconv.Atoi(n)
	}
	return ints
}

func nextInSequence(sequence []int) int {
	if allZero(sequence) {
		return 0
	}

	diffs := differences(sequence)
	return sequence[len(sequence)-1] + nextInSequence(diffs)
}

func prevInSequence(sequence []int) int {
	if allZero(sequence) {
		return 0
	}

	diffs := differences(sequence)
	return sequence[0] - prevInSequence(diffs)
}

func differences(sequence []int) []int {
	var diffs = make([]int, len(sequence)-1)
	for i := 1; i < len(sequence); i++ {
		diffs[i-1] = sequence[i] - sequence[i-1]
	}
	return diffs
}

func allZero(sequence []int) bool {
	for _, n := range sequence {
		if n != 0 {
			return false
		}
	}
	return true
}
