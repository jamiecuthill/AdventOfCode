package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
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

	rules, allPages := parse(input)
	fn := sortByRules(rules)

	for _, pages := range allPages {
		if slices.IsSortedFunc(pages, fn) {
			sum += pages[len(pages)/2]
		}
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int

	rules, allPages := parse(input)
	fn := sortByRules(rules)

	for _, pages := range allPages {
		if !slices.IsSortedFunc(pages, fn) {
			slices.SortStableFunc(pages, fn)
			sum += pages[len(pages)/2]
		}
	}

	return sum
}

func parse(input *bufio.Scanner) (rules [][2]int, allPages [][]int) {
	var rulesMode = true

	for input.Scan() {
		line := input.Text()
		if line == "" {
			rulesMode = false
			continue
		}

		if rulesMode {
			ruleParts := strings.Split(line, "|")
			a, _ := strconv.Atoi(ruleParts[0])
			b, _ := strconv.Atoi(ruleParts[1])
			rules = append(rules, [2]int{a, b})
		} else {
			var pages []int
			updateParts := strings.Split(line, ",")
			for i := range updateParts {
				page, _ := strconv.Atoi(updateParts[i])
				pages = append(pages, page)
			}
			allPages = append(allPages, pages)
		}
	}

	return
}

func sortByRules(rules [][2]int) func(int, int) int {
	return func(a, b int) int {
		for _, r := range rules {
			if r[0] == a && r[1] == b {
				return -1
			}
			if r[0] == b && r[1] == a {
				return 1
			}
		}
		return 0
	}
}
