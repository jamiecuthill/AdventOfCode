package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

	scanner := bufio.NewScanner(f)
	switch *part {
	case 1:
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

func Part1(input *bufio.Scanner) string {
	var total int
	for input.Scan() {
		snafu := input.Text()
		total += ToDec(snafu)
	}
	return ToSnafu(total)
}

func Part2(input *bufio.Scanner) int {
	return 0
}

func encode(l []int) string {
	s := strings.Builder{}

	for i := len(l) - 1; i >= 0; i-- {
		s.WriteByte(charMap[l[i]])
	}

	return s.String()
}

func ToSnafu(d int) string {
	t := d
	out := make([]int, 0)

	for t > 0 {
		rem := t % 5
		t /= 5

		if rem <= 2 {
			out = append(out, rem)
		} else {
			if rem == 3 {
				out = append(out, -2)
			}
			if rem == 4 {
				out = append(out, -1)
			}
			t += 1
		}
	}

	return encode(out)
}

func ToDec(snafu string) int {
	var acc int
	for i := 0; i < len(snafu); i++ {
		b := base(len(snafu) - 1 - i)
		acc += b * valueMap[snafu[i]]
	}
	return acc
}

var valueMap = map[byte]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

var charMap = map[int]byte{
	-2: '=',
	-1: '-',
	0:  '0',
	1:  '1',
	2:  '2',
}

func base(i int) int {
	if i == 0 {
		return 1
	}
	return int(math.Pow(5, float64(i)))
}
