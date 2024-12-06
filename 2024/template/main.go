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

func Part1(input *bufio.Scanner) int {
	var sum int

	var line string
	for input.Scan() {
		line += input.Text()
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int

	var line string
	for input.Scan() {
		line += input.Text()
	}

	return sum
}
