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
	sum := 0

	for input.Scan() {
		line := input.Text()
		first := first(line)
		last := last(line)
		val, _ := strconv.Atoi(first + last)
		sum += val
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	sum := 0

	for input.Scan() {
		line := input.Text()
		first := firstWithSpelling(line)
		last := lastWithSpelling(line)
		val, _ := strconv.Atoi(first + last)
		sum += val
	}

	return sum
}

var isDigit = regexp.MustCompile("[0-9]")

func first(line string) string {
	for i := 0; i < len(line); i++ {
		if isDigit.MatchString(string(line[i])) {
			return string(line[i])
		}
	}
	return ""
}

func last(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if isDigit.MatchString(string(line[i])) {
			return string(line[i])
		}
	}
	return ""
}

var numbers = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

func firstWithSpelling(line string) string {
	for i := 0; i < len(line); i++ {
		if isDigit.MatchString(string(line[i])) {
			return string(line[i])
		}
		for val, word := range numbers {
			if strings.HasPrefix(line[i:], word) {
				return val
			}
		}
	}
	return ""
}

func lastWithSpelling(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if isDigit.MatchString(string(line[i])) {
			return string(line[i])
		}
		for val, word := range numbers {
			if strings.HasSuffix(line[:i+1], word) {
				return val
			}
		}
	}
	return ""
}
