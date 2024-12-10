package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
		lineParts := strings.Split(line, ": ")
		cardParts := strings.Split(lineParts[1], " | ")
		winningNumbers := strings.Split(cardParts[0], " ")
		haveNumbers := strings.Split(cardParts[1], " ")

		sum += score(haveNumbers, winningNumbers)
	}

	return sum
}

func score(numbers []string, winners []string) int {
	var points int

	if i := numMatches(numbers, winners); i > 0 {
		points = 1 << (i - 1)
	}

	return points
}

func numMatches(numbers []string, winners []string) int {
	var i int
	for _, n := range numbers {
		n = strings.TrimSpace(n)
		for _, w := range winners {
			w = strings.TrimSpace(w)
			if n == w && n != "" {
				i++
			}
		}
	}
	return i
}

func Part2(input *bufio.Scanner) int {
	var sum, i int
	var copies = map[int]int{}

	for input.Scan() {
		i++
		line := input.Text()
		lineParts := strings.Split(line, ": ")
		cardParts := strings.Split(lineParts[1], " | ")
		winningNumbers := strings.Split(cardParts[0], " ")
		haveNumbers := strings.Split(cardParts[1], " ")

		n := numMatches(haveNumbers, winningNumbers)
		for j := 1; j <= n; j++ {
			copies[i+j] += (1 + copies[i])
		}
	}

	for _, n := range copies {
		sum += n
	}

	return sum + i
}
