package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
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

func Part1(input *bufio.Scanner) int {
	var index int
	var sum int

	for input.Scan() {
		index++

		var left, right []any
		err := json.Unmarshal(input.Bytes(), &left)
		if err != nil {
			panic(err)
		}
		input.Scan()
		err = json.Unmarshal(input.Bytes(), &right)
		if err != nil {
			panic(err)
		}
		// Blank line
		input.Scan()

		if ordered, _ := compare(left, right); ordered {
			sum += index
		}
	}

	return sum
}

func compare(left, right []any) (bool, bool) {
	// fmt.Printf("Compare %+v vs %+v\n", left, right)

	for i := 0; i < max(len(left), len(right)); i++ {
		// left ran out of items first so ordered
		if i >= len(left) && i < len(right) {
			return true, false
		}

		// right ran out of items so not ordered
		if i < len(left) && i >= len(right) {
			return false, false
		}

		l, lok := left[i].(float64)
		r, rok := right[i].(float64)

		// both values are integers
		if lok && rok {
			ordered, equal := compareInts(int(l), int(r))
			if equal {
				continue
			}
			return ordered, false
		}

		if l, ok := left[i].([]any); ok {
			var r []any
			switch t := right[i].(type) {
			case float64:
				r = []any{t}
			case []any:
				r = t
			}
			ordered, cont := compare(l, r)
			if !cont {
				return ordered, false
			}
			continue
		}

		if r, ok := right[i].([]any); ok {
			var l []any
			switch t := left[i].(type) {
			case float64:
				l = []any{t}
			case []any:
				l = t
			}
			ordered, cont := compare(l, r)
			if !cont {
				return ordered, false
			}
			continue
		}
	}

	return false, true
}

func compareInts(left, right int) (bool, bool) {
	return left < right, left == right
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Part2(input *bufio.Scanner) int {
	var packets = [][]any{}

	for input.Scan() {
		if input.Text() == "" {
			continue
		}

		var pkt []any
		err := json.Unmarshal(input.Bytes(), &pkt)
		if err != nil {
			panic(err)
		}
		packets = append(packets, pkt)
	}

	packets = append(packets, []any{[]any{float64(2)}})
	packets = append(packets, []any{[]any{float64(6)}})

	sort.Sort(ByPacket(packets))

	key := 1
	for i, pkt := range packets {
		p := str(pkt)
		// fmt.Printf("%d: %v\n", i+1, p)
		if p == "[[2]]" || p == "[[6]]" {
			key *= (i + 1)
		}
	}
	return key
}

func str(pkt []any) string {
	return fmt.Sprintf("%+v", pkt)
}

type ByPacket [][]any

func (a ByPacket) Len() int      { return len(a) }
func (a ByPacket) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPacket) Less(i, j int) bool {
	ordered, _ := compare(a[i], a[j])
	return ordered
}
