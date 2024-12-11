package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

func parse(line string) (int, []int) {
	parts := strings.Split(line, ": ")
	intTestValue, _ := strconv.Atoi(parts[0])
	testValue := int(intTestValue)
	var numbers []int
	for _, v := range strings.Split(parts[1], " ") {
		num, _ := strconv.Atoi(v)
		numbers = append(numbers, int(num))
	}

	return testValue, numbers
}

const (
	Add = iota
	Multiply
	Concat
)

func Part1(input *bufio.Scanner) int {
	var sum int

	for input.Scan() {
		targetValue, numbers := parse(input.Text())

		if slices.Contains(CalcWith(Add, Multiply)(numbers, targetValue), targetValue) {
			sum += int(targetValue)
		}
	}
	return sum
}

func CalcWith(operators ...int) (fn func([]int, int) []int) {
	fn = func(numbers []int, targetValue int) []int {
		if len(numbers) == 1 {
			return numbers
		}

		var options []int

		for _, op := range operators {
			var value int
			switch op {
			case Add:
				value = numbers[0] + numbers[1]
			case Multiply:
				value = numbers[0] * numbers[1]
			case Concat:
				value = (numbers[0] * int(math.Pow(10, float64(len(strconv.Itoa(int(numbers[1]))))))) + numbers[1]
			}

			// Exclude value if exceeds target value
			if value > targetValue {
				continue
			}
			options = append(options, fn(append([]int{value}, numbers[2:]...), targetValue)...)
		}

		return options
	}
	return fn

}

func Part2(input *bufio.Scanner) int {
	var sum int

	for input.Scan() {
		testValue, numbers := parse(input.Text())

		if slices.Contains(CalcWith(Add, Multiply, Concat)(numbers, testValue), testValue) {
			sum += testValue
		}
	}
	return sum
}
