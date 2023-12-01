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

	scanner := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Solve(scanner))
	case 2:
		screen := Solve2(scanner)

		for y := range screen {
			for x := range screen[y] {
				if screen[y][x] {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println("")
		}
	}
}

type command struct {
	op     string
	length uint
	fn     func(registers map[string]int)
}

func (c *command) tick(registers map[string]int) {
	c.length--

	if c.length == 0 && c.fn != nil {
		c.fn(registers)
	}
}

func (c command) done() bool {
	return c.length == 0
}

func Solve(scanner *bufio.Scanner) int {
	commands := []command{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "noop":
			commands = append(commands, command{
				op:     parts[0],
				length: 1,
			})
		case "addx":
			operand, _ := strconv.Atoi(parts[1])
			commands = append(commands, command{
				op:     parts[0],
				length: 2,
				fn: func(registers map[string]int) {
					registers["X"] += operand
				},
			})
		}
	}
	var registers = map[string]int{
		"X": 1,
	}
	var current *command
	var cycle uint = 1
	var signals int
	var capture = 20
	for {
		if cycle == uint(capture) {
			signals += (int(cycle) * registers["X"])
			capture += 40
		}

		// Empty
		if current == nil {
			if len(commands) == 0 {
				// Execution complete
				break
			}

			// Load command
			current = &commands[0]
			commands = commands[1:]
		}

		// Run command
		current.tick(registers)

		// Unload completed commands
		if current.done() {
			current = nil
		}

		cycle++
	}

	return signals
}

const (
	crtWidth    = 40
	crtHeight   = 6
	spriteWidth = 3
)

func Solve2(scanner *bufio.Scanner) [][]bool {
	commands := []command{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "noop":
			commands = append(commands, command{
				op:     parts[0],
				length: 1,
			})
		case "addx":
			operand, _ := strconv.Atoi(parts[1])
			commands = append(commands, command{
				op:     parts[0],
				length: 2,
				fn: func(registers map[string]int) {
					registers["X"] += operand
				},
			})
		}
	}
	var registers = map[string]int{
		"X": 1,
	}
	var current *command
	var cycle uint = 1

	// Init screen
	var screen = make([][]bool, crtHeight)
	for y := 0; y < crtHeight; y++ {
		screen[y] = make([]bool, crtWidth)
	}

	for {
		draw(screen, cycle, registers)

		// Empty
		if current == nil {
			if len(commands) == 0 {
				// Execution complete
				break
			}

			// Load command
			current = &commands[0]
			commands = commands[1:]
		}

		// Run command
		current.tick(registers)

		// Unload completed commands
		if current.done() {
			current = nil
		}

		cycle++
	}

	return screen
}

func draw(screen [][]bool, cycle uint, registers map[string]int) {
	pos := int(cycle) - 1
	y := pos / crtWidth
	x := pos % crtWidth
	if x >= registers["X"]-1 && x <= registers["X"]+1 {
		screen[y][x] = true
	}
}
