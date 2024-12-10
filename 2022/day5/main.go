package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var stacks = make(map[int][]string)

	stacks[1] = []string{"B", "P", "N", "Q", "H", "D", "R", "T"}
	stacks[2] = []string{"W", "G", "B", "J", "T", "V"}
	stacks[3] = []string{"N", "R", "H", "D", "S", "V", "M", "Q"}
	stacks[4] = []string{"P", "Z", "N", "M", "C"}
	stacks[5] = []string{"D", "Z", "B"}
	stacks[6] = []string{"V", "C", "W", "Z"}
	stacks[7] = []string{"G", "Z", "N", "C", "V", "Q", "L", "S"}
	stacks[8] = []string{"L", "G", "J", "M", "D", "N", "V"}
	stacks[9] = []string{"T", "P", "M", "F", "Z", "C", "G"}

	// [T]     [Q]             [S]
	// [R]     [M]             [L] [V] [G]
	// [D] [V] [V]             [Q] [N] [C]
	// [H] [T] [S] [C]         [V] [D] [Z]
	// [Q] [J] [D] [M]     [Z] [C] [M] [F]
	// [N] [B] [H] [N] [B] [W] [N] [J] [M]
	// [P] [G] [R] [Z] [Z] [C] [Z] [G] [P]
	// [B] [W] [N] [P] [D] [V] [G] [L] [T]
	//  1   2   3   4   5   6   7   8   9

	var matcher = regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")

	for scanner.Scan() {
		line := scanner.Bytes()

		matches := matcher.FindSubmatch(line)

		// fmt.Printf("%q\n", matches)

		numCrates, _ := strconv.Atoi(string(matches[1]))
		from, _ := strconv.Atoi(string(matches[2]))
		to, _ := strconv.Atoi(string(matches[3]))

		// Part 1
		// for i := 0; i < numCrates; i++ {
		// c := stacks[from][len(stacks[from])-1]
		// stacks[to] = append(stacks[to], c)
		// stacks[from] = stacks[from][:len(stacks[from])-1]
		// }

		// Part 2
		var tmp []string = stacks[from][len(stacks[from])-numCrates:]
		fmt.Printf("%d creates %q\n", numCrates, tmp)
		stacks[to] = append(stacks[to], tmp...)
		stacks[from] = stacks[from][0 : len(stacks[from])-numCrates]
	}

	for i := 1; i <= 9; i++ {
		fmt.Printf("%s", stacks[i][len(stacks[i])-1])
	}
	fmt.Println("")
}
