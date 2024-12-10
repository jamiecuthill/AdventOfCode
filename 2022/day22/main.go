package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

func Part1(input *bufio.Scanner) int {
	var rows []span
	var wallRows = make(map[int][]int)
	var wallColumns = make(map[int][]int)
	var max_x int

	var y int
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		var curr *span
		for x, c := range line {
			if (c == '.' || c == '#') && curr == nil {
				curr = &span{left: x}
			}

			if c == '#' {
				wallRows[y] = append(wallRows[y], x)
				wallColumns[x] = append(wallColumns[x], y)
			}

			if c == ' ' && curr != nil {
				curr.right = x
				curr.right = len(line) - 1
				if curr.right > max_x {
					max_x = curr.right
				}
				rows = append(rows, *curr)
				curr = nil
			}
		}
		if curr != nil {
			curr.right = len(line) - 1
			if curr.right > max_x {
				max_x = curr.right
			}
			rows = append(rows, *curr)
		}

		y++
	}

	var columns = make([]span, max_x+1)
	for i := range columns {
		columns[i].left = math.MaxInt
	}

	for y, row := range rows {
		for x := row.left; x <= row.right; x++ {
			if columns[x].left > y {
				columns[x].left = y
			}
			if columns[x].right < y {
				columns[x].right = y
			}
		}
	}

	// next line
	input.Scan()

	var facing, x int
	y = 0
	x = rows[0].left

	var distance strings.Builder
	for _, v := range input.Text() {
		if v >= '0' && v <= '9' {
			distance.WriteRune(v)
		} else if distance.Len() > 0 {
			// Move
			d, _ := strconv.Atoi(distance.String())
			switch facing {
			case 0:
				x = move(x, d, wallRows[y], rows[y], 1)
			case 1:
				y = move(y, d, wallColumns[x], columns[x], 1)
			case 2:
				x = move(x, d, wallRows[y], rows[y], -1)
			case 3:
				y = move(y, d, wallColumns[x], columns[x], -1)
			}
			distance.Reset()
		}

		if v == 'R' {
			facing++
			if facing == 4 {
				facing = 0
			}
		}

		if v == 'L' {
			facing--
			if facing == -1 {
				facing = 3
			}
		}
	}
	if distance.Len() > 0 {
		// Move
		d, _ := strconv.Atoi(distance.String())
		switch facing {
		case 0:
			x = move(x, d, wallRows[y], rows[y], 1)
		case 1:
			y = move(y, d, wallColumns[x], columns[x], 1)
		case 2:
			x = move(x, d, wallRows[y], rows[y], -1)
		case 3:
			y = move(y, d, wallColumns[x], columns[x], -1)
		}
		distance.Reset()
	}

	return (1000 * (y + 1)) + (4 * (x + 1)) + facing
}

func Part2(input *bufio.Scanner) int {
	return 0
}

type span struct {
	left  int
	right int
}

func findWall(walls []int, tiles span, from, to int, rev bool) (int, bool) {
	if rev {
		walls = walls[:]
		sort.Sort(sort.Reverse(sort.IntSlice(walls)))
	}
	for _, w := range walls {
		if w >= from && w <= to {
			if w == 0 && !rev {
				return tiles.right, true
			}
			if w == tiles.right && rev {
				return 0, true
			}
			if rev {
				return w + 1, true
			} else {
				return w - 1, true
			}
		}
	}
	return 0, false
}

func move(pos, d int, walls []int, tiles span, direction int) int {
	for i := 0; i < d; i++ {
		next := pos + direction
		if next > tiles.right {
			next = tiles.left
		}
		if next < tiles.left {
			next = tiles.right
		}
		if inArray(walls, next) {
			return pos
		}
		pos = next
	}
	return pos

	// if w, ok := findWall(walls, tiles, pos, pos+d, false); ok {
	// 	return w
	// }
	// if pos+d > tiles.right {
	// 	remainder := d - (tiles.right - pos)
	// 	l := tiles.right - tiles.left
	// 	if remainder/l > 1 {
	// 		if w, ok := findWall(walls, tiles, tiles.left, tiles.right, false); ok {
	// 			return w
	// 		}
	// 		remainder = remainder % l
	// 	}
	// 	if w, ok := findWall(walls, tiles, tiles.left, tiles.left+remainder, false); ok {
	// 		return w
	// 	}

	// 	return tiles.left + remainder - 1
	// }
	// return pos + d
}

func moveNeg(pos, d int, walls []int, tiles span) int {
	if w, ok := findWall(walls, tiles, pos-d, pos, true); ok {
		return w
	}
	if pos-d < tiles.left {
		remainder := d - (pos - tiles.left)
		l := tiles.right - tiles.left
		if remainder/l > 1 {
			if w, ok := findWall(walls, tiles, tiles.left, tiles.right, true); ok {
				return w
			}
			remainder = remainder % l
		}
		if w, ok := findWall(walls, tiles, tiles.right-remainder, tiles.right, true); ok {
			return w
		}

		return tiles.right - remainder + 1
	}
	return pos - d
}

func inArray[T comparable](list []T, x T) bool {
	for i := range list {
		if list[i] == x {
			return true
		}
	}
	return false
}
