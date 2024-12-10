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

func Part1(input *bufio.Scanner) int {
	var margin int = 1

	var timeLine, recordLine string

	for input.Scan() {
		timeLine = strings.TrimPrefix(input.Text(), "Time: ")
		input.Scan()
		recordLine = strings.TrimPrefix(input.Text(), "Distance: ")
	}

	timeValues := toIntValues(strings.Split(timeLine, " "))
	recordValues := toIntValues(strings.Split(recordLine, " "))

	for i := 0; i < len(timeValues); i++ {
		n := waysToBeat(timeValues[i], recordValues[i])
		margin *= n
	}

	return margin
}

func Part2(input *bufio.Scanner) int {
	var timeLine, recordLine string

	for input.Scan() {
		timeLine = strings.TrimPrefix(input.Text(), "Time: ")
		input.Scan()
		recordLine = strings.TrimPrefix(input.Text(), "Distance: ")
	}

	timeValue, _ := strconv.Atoi(strings.ReplaceAll(timeLine, " ", ""))
	recordValue, _ := strconv.Atoi(strings.ReplaceAll(recordLine, " ", ""))

	return waysToBeat(timeValue, recordValue)
}

func toIntValues(input []string) []int {
	var out []int
	for _, v := range input {
		if v != "" {
			t, _ := strconv.Atoi(v)
			out = append(out, t)
		}
	}
	return out
}

func waysToBeat(t, record int) int {
	b := math.Pow(float64(t), 2.0)
	c := math.Sqrt(b - (-4.0 * -float64(record)))
	from := (-float64(t) + c) / (-2.0)
	to := (-float64(t) - c) / (-2.0)
	return int(math.Ceil(to)) - int(math.Floor(from)) - 1
}
