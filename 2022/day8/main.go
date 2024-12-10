package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const gridSize = 99

var part = flag.Int("part", 1, "Part 1 or part 2?")

func main() {
	flag.Parse()

	start := time.Now()

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var forest = make([][]int, 99) // X, Y, H

	var y int

	for scanner.Scan() {
		line := scanner.Text()

		for x := range line {
			if forest[x] == nil {
				forest[x] = make([]int, 99)
			}
			height, _ := strconv.Atoi(string(line[x]))
			forest[x][y] = height
		}

		y++
	}

	if *part == 1 {
		visibleTrees := 0

		for x := range forest {
			for y := range forest[x] {
				t := tree{x: x, y: y, height: forest[x][y]}

				if isVisible(t, forest) {
					visibleTrees++
				}
			}
		}

		fmt.Println(visibleTrees)
		fmt.Println(time.Since(start))
		os.Exit(0)
	}

	if *part == 2 {
		scenicScore := 0
		for x := range forest {
			for y := range forest[x] {
				if this := score(coordinate{x, y}, forest); this > scenicScore {
					// fmt.Printf("New high score %d for tree %+v\n", this, t)
					scenicScore = this
				}
			}
		}

		fmt.Println(scenicScore)
		fmt.Println(time.Since(start))
		os.Exit(0)
	}
}

type tree struct {
	x      int
	y      int
	height int
}

type coordinate struct {
	x int
	y int
}

func move(p, direction coordinate) coordinate {
	return coordinate{
		x: p.x + direction.x,
		y: p.y + direction.y,
	}
}

func inForest(p coordinate, size int) bool {
	if p.x < 0 || p.y < 0 {
		return false
	}

	if p.x >= size || p.y >= size {
		return false
	}

	return true
}

func countTrees(t coordinate, forest [][]int, direction coordinate) int {
	var trees int
	var p coordinate = move(t, direction)
	for inForest(p, len(forest)) {
		if forest[p.x][p.y] >= forest[t.x][t.y] {
			trees++
			break
		}
		trees++
		p = move(p, direction)
	}
	return trees
}

func score(t coordinate, forest [][]int) int {
	return countTrees(t, forest, coordinate{x: -1}) *
		countTrees(t, forest, coordinate{x: +1}) *
		countTrees(t, forest, coordinate{y: -1}) *
		countTrees(t, forest, coordinate{y: +1})
}

func isVisible(t tree, forest [][]int) (visible bool) {
	// Edge tree?
	if t.x == 0 || t.y == 0 || t.x == gridSize-1 || t.y == gridSize-1 {
		return true
	}

	// visible left?
	visible = true
	for x := t.x - 1; x >= 0; x-- {
		if forest[x][t.y] < t.height {
			continue
		}
		visible = false
		break
	}

	if visible {
		return true
	}

	// visible right
	visible = true
	for x := t.x + 1; x < gridSize; x++ {
		if forest[x][t.y] < t.height {
			continue
		}
		visible = false
		break
	}
	if visible {
		return true
	}

	// visible top
	visible = true
	for y := t.y - 1; y >= 0; y-- {
		if forest[t.x][y] < t.height {
			continue
		}
		visible = false
		break
	}
	if visible {
		return true
	}

	// visible bottom
	visible = true
	for y := t.y + 1; y < gridSize; y++ {
		if forest[t.x][y] < t.height {
			continue
		}
		visible = false
		break
	}

	return visible
}
