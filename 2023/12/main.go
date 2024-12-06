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
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		dims := toInts(strings.Split(parts[1], ","))
		c := CalcPositions(parts[0], dims)
		sum += c
	}
	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		dims := toInts(strings.Split(strings.Join(expand(parts[1], 5), ","), ","))
		reports := expand(parts[0], 5)
		c := CalcPositions(strings.Join(reports, "?"), dims)
		sum += c
	}
	return sum
}

// func Configurations(report string, dims []int) int {
// 	c := combinations(report)
// 	var valid int
// 	for _, combo := range c {
// 		if validate(&combo, dims) {
// 			valid++
// 		}
// 	}
// 	return valid
// }

// TODO Rewrite this as an iterative solution so
// that we do not need all combinations in memory
// at once
// func combinations(report string) []string {
// 	if len(report) == 0 {
// 		return []string{""}
// 	}

// 	var combos []string
// 	c := string(report[0])
// 	rest := report[1:]

// 	switch c {
// 	case ".":
// 		fallthrough
// 	case "#":
// 		for _, combo := range combinations(rest) {
// 			combos = append(combos, c+combo)
// 		}
// 	case "?":
// 		for _, combo := range combinations(rest) {
// 			combos = append(combos, "#"+combo, "."+combo)
// 		}
// 	}

// 	return combos
// }

// func validate(report *string, dims []int) bool {
// 	if len(dims) == 0 && !strings.Contains(*report, "#") {
// 		return true
// 	}

// 	if *report == "#" {
// 		return len(dims) == 1 && dims[0] == 1
// 	}

// 	var parts = make([]int, len(dims))
// 	var i int
// 	var lastWasBroken bool

// 	for _, c := range *report {
// 		if c == '#' {
// 			if i == len(parts) {
// 				parts = append(parts, 0)
// 			}
// 			lastWasBroken = true
// 			parts[i]++
// 		} else {
// 			if lastWasBroken && parts[i] > 0 {
// 				i++
// 				lastWasBroken = false
// 			}
// 		}
// 	}

// 	return reflect.DeepEqual(parts, dims)
// }

func toInts(s []string) []int {
	var ints = make([]int, len(s))
	for i, v := range s {
		ints[i], _ = strconv.Atoi(v)
	}
	return ints
}

func expand(str string, n int) []string {
	var out []string
	for i := 0; i < 5; i++ {
		out = append(out, str)
	}
	return out
}

func CalcPositions(input string, lengths []int) int {
	chars := "." + input + "."
	springs := Base(lengths)

	table := make([][]int, len(chars)+1)
	for i := range table {
		table[i] = make([]int, len(springs)+1)
	}

	table[len(chars)][len(springs)] = 1

	for c := len(chars) - 1; c >= 0; c-- {
		for s := len(springs) - 1; s >= 0; s-- {
			if chars[c] != '.' && springs[s] == '#' {
				table[c][s] = table[c+1][s+1]
			} else if chars[c] != '#' && springs[s] == '.' {
				table[c][s] = table[c+1][s+1] + table[c+1][s]
			} else {
				table[c][s] = 0
			}
		}
	}

	return table[0][0]
}

func Base(lengths []int) string {
	out := make([]byte, 0, sum(lengths)+len(lengths)+1)
	out = append(out, '.')

	for _, length := range lengths {
		for i := 0; i < length; i++ {
			out = append(out, '#')
		}
		out = append(out, '.')
	}

	return string(out)
}

func sum(ints []int) int {
	var t int
	for _, i := range ints {
		t += i
	}
	return t
}
