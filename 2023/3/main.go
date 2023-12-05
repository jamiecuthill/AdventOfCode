package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

const dot = '.'
const first = '0'
const last = '9'
const star = '*'

func Part1(input *bufio.Scanner) int {
	var world = [][]rune{}

	var x, y int

	for input.Scan() {
		line := input.Text()
		world = append(world, make([]rune, len(line)))
		for x = 0; x < len(line); x++ {
			world[y][x] = rune(line[x])
		}
		y++
	}
	var sum int
	var num string
	var neighbors int

	for y, row := range world {
		for x = 0; x < len(row); x++ {
			cell := row[x]
			if cell >= first && cell <= last {
				num += string(cell)
				if y > 0 {
					if x > 0 {
						if isSymbol(world[y-1][x-1]) {
							neighbors++
						}
					}
					if isSymbol(world[y-1][x]) {
						neighbors++
					}
					if x < len(row)-1 {
						if isSymbol(world[y-1][x+1]) {
							neighbors++
						}
					}
				}
				if x > 0 {
					if isSymbol(world[y][x-1]) {
						neighbors++
					}
				}

				if x < len(row)-1 {
					if isSymbol(world[y][x+1]) {
						neighbors++
					}
				}

				if y < len(world)-1 {
					if x > 0 {
						if isSymbol(world[y+1][x-1]) {
							neighbors++
						}
					}
					if isSymbol(world[y+1][x]) {
						neighbors++
					}
					if x < len(row)-1 {
						if isSymbol(world[y+1][x+1]) {
							neighbors++
						}
					}
				}

				continue
			} else {
				if len(num) > 0 {
					if neighbors > 0 {
						numInt, _ := strconv.Atoi(num)
						sum += numInt
					}
					// reset
					num = ""
					neighbors = 0
				}
			}
		}
		if len(num) > 0 {
			if neighbors > 0 {
				numInt, _ := strconv.Atoi(num)
				sum += numInt
			}
			// reset
			num = ""
			neighbors = 0
		}
	}

	return sum
}

func isSymbol(r rune) bool {
	if r >= first && r <= last {
		return false
	}
	if r == dot {
		return false
	}
	return true
}

func isGear(r rune) bool {
	return r == star
}

type pos struct {
	x, y int
	num  string
}

func Part2(input *bufio.Scanner) int {
	var numbers []pos
	var gears []pos

	var x, xx, y, yy int
	var num string

	for input.Scan() {
		line := input.Text()
		for x = 0; x < len(line); x++ {
			if isGear(rune(line[x])) {
				gears = append(gears, pos{x, y, string(line[x])})
				if len(num) > 0 {
					numbers = append(numbers, pos{xx, yy, num})
					num = ""
				}
				continue
			}
			cell := rune(line[x])
			if cell >= first && cell <= last {
				if len(num) == 0 {
					xx = x
					yy = y
				}
				num += string(cell)
				continue
			} else {
				if len(num) > 0 {
					numbers = append(numbers, pos{xx, yy, num})
					num = ""
				}
			}
		}
		if len(num) > 0 {
			numbers = append(numbers, pos{xx, yy, num})
			num = ""
		}
		y++
	}

	var sum int

	for _, gear := range gears {
		var adjacent []pos
		for _, number := range numbers {
			if touches(number, gear) {
				adjacent = append(adjacent, number)
			}

			if len(adjacent) > 2 {
				break
			}
		}

		if len(adjacent) == 2 {
			r1, _ := strconv.Atoi(adjacent[0].num)
			r2, _ := strconv.Atoi(adjacent[1].num)
			sum += (r1 * r2)
		}
	}

	return sum
}

func touches(number, gear pos) bool {
	if gear.x >= number.x-1 &&
		gear.x <= number.x+len(number.num) &&
		gear.y >= number.y-1 &&
		gear.y <= number.y+1 {
		return true
	}
	return false
}
