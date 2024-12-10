package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var values map[rune]int
	values = make(map[rune]int)
	var char rune = 97
	for i := 0; i < 26; i++ {
		values[rune(int(char)+i)] = i + 1
	}

	char = 65
	for i := 0; i < 26; i++ {
		values[rune(int(char)+i)] = 26 + i + 1
	}

	var prioritySum int
	var badgePrioritySum int

	var lset map[rune]struct{}
	var trioset map[rune]struct{} = make(map[rune]struct{})

	var groupNum = 0

	for scanner.Scan() {
		lset = make(map[rune]struct{})
		line := scanner.Text()

		midpoint := len(line) / 2

		for i := range line[0:midpoint] {
			lset[rune(line[i])] = struct{}{}
		}

		for i := range line[midpoint:] {
			if _, ok := lset[rune(line[midpoint+i])]; ok {
				item := rune(line[midpoint+i])
				prioritySum += values[item]
				break
			}
		}

		// Part 2

		groupNum++
		if groupNum > 3 {
			trioset = make(map[rune]struct{})
			groupNum = 1
		}

		var badge rune

		switch groupNum {
		case 1:
			fmt.Println("Elf 1:", line)
			for _, r := range line {
				trioset[rune(r)] = struct{}{}
			}
		case 2:
			fmt.Println("Elf 2:", line)
			var new = make(map[rune]struct{})
			for _, r := range line {
				if _, ok := trioset[rune(r)]; ok {
					new[rune(r)] = struct{}{}
				}
			}
			trioset = new
		case 3:
			fmt.Println("Elf 3:", line)
			for _, r := range line {
				if _, ok := trioset[rune(r)]; ok {
					badge = rune(r)
				}
			}

			badgePrioritySum += values[badge]
		}
	}

	fmt.Println("Priorities Sum:", prioritySum)
	fmt.Println("Badhe Priorities Sum:", badgePrioritySum)
}
