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

const CostA int = 3
const CostB int = 1

type P struct {
	X, Y uint64
}

type Machine struct {
	Prize   P
	ButtonA P
	ButtonB P
}

func MachineCost(prize, a, b P) uint64 {
	// Try applying cramers rule

	// We have two equations with two unknowns
	// a.X * x + b.X * y = prize.X
	// b.X * x + b.Y * y = prize.Y

	// If the determinant is 0, then the system is unsolvable
	var det int = int(a.X*b.Y) - int(a.Y*b.X)

	// If the determinant is not 0, then we can solve for x and y
	var x int = (int(prize.X*b.Y) - int(prize.Y*b.X))
	var y int = (int(a.X*prize.Y) - int(a.Y*prize.X))

	// Ensure we've found an integer solution
	if x%det != 0 || y%det != 0 {
		return 0
	}

	// Can't have a negative number of button presses
	if x/det < 0 || y/det < 0 {
		return 0
	}

	// Return the number of each button press
	return uint64(CostA*(x/det) + CostB*(y/det))
}

func parse(input *bufio.Scanner) []Machine {
	var machines []Machine
	var m Machine

	for input.Scan() {
		line := input.Text()
		if line == "" {
			machines = append(machines, m)
			m = Machine{}
			continue
		}

		// Parse into current machine state
		if strings.HasPrefix(line, "Button ") {
			switch line[7] {
			case 'A':
				m.ButtonA = toCoordinates(line[10:])
			case 'B':
				m.ButtonB = toCoordinates(line[10:])
			}
			continue
		}

		if strings.HasPrefix(line, "Prize: ") {
			m.Prize = toCoordinates(line[7:])
			continue
		}
	}

	// Catch last machine if no newline
	if m.Prize.X > 0 {
		machines = append(machines, m)
	}
	return machines
}

func Part1(input *bufio.Scanner) uint64 {
	machines := parse(input)

	var sum uint64
	for _, m := range machines {
		sum += MachineCost(m.Prize, m.ButtonA, m.ButtonB)
	}
	return sum
}

func toCoordinates(coordinates string) P {
	coordParts := strings.Split(coordinates, ", ")
	x, _ := strconv.Atoi(coordParts[0][2:])
	y, _ := strconv.Atoi(coordParts[1][2:])
	return P{X: uint64(x), Y: uint64(y)}
}

func Part2(input *bufio.Scanner) uint64 {
	machines := parse(input)

	var sum = uint64(0)
	for _, m := range machines {
		m.Prize.X += 10_000_000_000_000
		m.Prize.Y += 10_000_000_000_000
		sum += MachineCost(m.Prize, m.ButtonA, m.ButtonB)
	}
	return sum
}
