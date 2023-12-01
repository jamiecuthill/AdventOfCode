package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var countFullyContained int
	var countOverlapping int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		left, right := parts[0], parts[1]
		if within(left, right) || within(right, left) {
			countFullyContained++
		}
		if overlapping(left, right) {
			fmt.Println("✅")
			countOverlapping++
		} else {
			fmt.Println("❌")
		}

	}

	fmt.Println("Fully Contained Count:", countFullyContained)
	fmt.Println("Overlapping Count:", countOverlapping)
}

func within(left, right string) bool {
	llower, lupper := parseRange(left)
	rlower, rupper := parseRange(right)

	return rlower >= llower && rupper <= lupper
}

func overlapping(left, right string) bool {
	llower, lupper := parseRange(left)
	rlower, rupper := parseRange(right)

	if rlower < llower {
		// fmt.Printf("RL %d - %d ==== %d - %d ", rlower, rupper, llower, lupper)

		return rupper >= llower
	}

	// fmt.Printf("LR %d - %d ==== %d - %d ", llower, lupper, rlower, rupper)
	return lupper >= rlower

}

func parseRange(in string) (int, int) {
	parts := strings.Split(in, "-")
	lower, _ := strconv.Atoi(parts[0])
	upper, _ := strconv.Atoi(parts[1])
	return lower, upper
}
