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

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

var multiplyPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var enablePattern = regexp.MustCompile(`do\(\)`)
var disablePattern = regexp.MustCompile(`don't\(\)`)

func Part1(input *bufio.Scanner) int {
	var sum int

	var line string
	for input.Scan() {
		line += input.Text()
	}
	matches := multiplyPattern.FindAllStringSubmatch(line, -1)

	for i := range matches {
		a, _ := strconv.Atoi(matches[i][1])
		b, _ := strconv.Atoi(matches[i][2])
		sum += a * b
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int

	var line string
	for input.Scan() {
		line += input.Text()
	}

	multiplyInstructions := multiplyPattern.FindAllStringSubmatchIndex(line, -1)
	enableInstructions := enablePattern.FindAllStringIndex(line, -1)
	disableInstructions := disablePattern.FindAllStringIndex(line, -1)

	for i := range multiplyInstructions {
		a, _ := strconv.Atoi(line[multiplyInstructions[i][2]:multiplyInstructions[i][3]])
		b, _ := strconv.Atoi(line[multiplyInstructions[i][4]:multiplyInstructions[i][5]])
		if areWeEnabled(multiplyInstructions[i][0], enableInstructions, disableInstructions) {
			sum += a * b
		}
	}

	return sum
}

type p struct {
	i       int
	enabled bool
}

func areWeEnabled(i int, enabled, disabled [][]int) bool {
	var changes []p

	for i := range enabled {
		changes = append(changes, p{i: enabled[i][0], enabled: true})
	}

	for i := range disabled {
		changes = append(changes, p{i: disabled[i][0], enabled: false})
	}

	sort.SliceStable(changes, func(i, j int) bool { return changes[i].i < changes[j].i })

	if i < changes[0].i {
		return true
	}

	for j := 1; j < len(changes); j++ {
		if i >= changes[j-1].i && i <= changes[j].i {
			return changes[j-1].enabled
		}
	}

	return changes[len(changes)-1].enabled
}
