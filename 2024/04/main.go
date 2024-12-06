package main

import (
	"bufio"
	"flag"
	"fmt"
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

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

const xmas = "XMAS"
const samx = "SAMX"

func Part1(input *bufio.Scanner) int {
	var sum int

	var lines []string

	for input.Scan() {
		lines = append(lines, input.Text())
	}

	// ROWS
	for _, line := range lines {
		sum += strings.Count(line, xmas)
		sum += strings.Count(line, samx)
	}

	// COLUMNS
	for i := 0; i < len(lines[0]); i++ {
		var column = make([]byte, 0, len(lines))
		for j := range lines {
			column = append(column, lines[j][i])
		}

		sum += strings.Count(string(column), xmas)
		sum += strings.Count(string(column), samx)
	}

	// DIAGONAL NW -> SE
	var startingPoints = []struct{ x, y int }{}
	for x := 1; x < len(lines[0]); x++ {
		startingPoints = append(startingPoints, struct {
			x int
			y int
		}{x, 0})
	}
	for y := 0; y < len(lines); y++ {
		startingPoints = append(startingPoints, struct {
			x int
			y int
		}{0, y})
	}

	for o := range startingPoints {
		var i = startingPoints[o]
		var line []byte
		for {
			line = append(line, lines[i.x][i.y])
			i.x++
			i.y++
			if i.x == len(lines[0]) || i.y == len(lines) {
				break
			}
		}
		linestr := string(line)
		sum += strings.Count(linestr, xmas)
		sum += strings.Count(linestr, samx)
	}

	// DIAGONAL NE -> SW
	startingPoints = []struct{ x, y int }{}
	for x := 0; x < len(lines[0]); x++ {
		startingPoints = append(startingPoints, struct {
			x int
			y int
		}{x, 0})
	}
	for y := 1; y < len(lines); y++ {
		startingPoints = append(startingPoints, struct {
			x int
			y int
		}{len(lines[0]) - 1, y})
	}
	for o := range startingPoints {
		var i = startingPoints[o]
		var line []byte
		for {
			line = append(line, lines[i.x][i.y])
			i.x--
			i.y++
			if i.x < 0 || i.y >= len(lines) {
				break
			}
		}
		linestr := string(line)
		sum += strings.Count(linestr, xmas)
		sum += strings.Count(linestr, samx)
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	var sum int

	var lines []string
	for input.Scan() {
		lines = append(lines, input.Text())
	}

	// Look for the following
	var patterns = map[string]struct{}{
		"MSSM": {},
		"SMMS": {},
		"SMSM": {},
		"MSMS": {},
	}
	var corners = make([]byte, 4)

	for x := 1; x < len(lines[0])-1; x++ {
		for y := 1; y < len(lines)-1; y++ {
			// Have we found the right center char?
			if lines[y][x] != 'A' {
				continue
			}

			corners[0] = lines[y-1][x-1]
			corners[1] = lines[y+1][x+1]
			corners[2] = lines[y-1][x+1]
			corners[3] = lines[y+1][x-1]

			if _, ok := patterns[string(corners)]; ok {
				sum++
			}
		}
	}

	return sum
}
