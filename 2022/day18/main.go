package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

func Part1(input *bufio.Scanner) int {
	s := make(shape)
	for input.Scan() {
		coords := input.Text()
		coordParts := strings.Split(coords, ",")
		s.Build(cube{
			x: toInt(coordParts[0]),
			y: toInt(coordParts[1]),
			z: toInt(coordParts[2]),
		})
	}
	return s.SurfaceArea()
}

func Part2(input *bufio.Scanner) int {
	s := make(shape)
	for input.Scan() {
		coords := input.Text()
		coordParts := strings.Split(coords, ",")
		s.Build(cube{
			x: toInt(coordParts[0]),
			y: toInt(coordParts[1]),
			z: toInt(coordParts[2]),
		})
	}
	return s.OuterArea()
}

func toInt(in string) int {
	i, _ := strconv.Atoi(in)
	return i
}

type cube struct {
	x int
	y int
	z int
}

func (c cube) apply(d cube) cube {
	return cube{
		x: c.x + d.x,
		y: c.y + d.y,
		z: c.z + d.z,
	}
}

type shape map[cube]byte

func (s shape) Build(c cube) {
	s[c] = 1
}

var directions = []cube{
	{x: 1},
	{x: -1},
	{y: 1},
	{y: -1},
	{z: 1},
	{z: -1},
}

func (s shape) SurfaceArea() int {
	var sum int
	for c := range s {
		toAdd := 6
		for i := range directions {
			if _, touching := s[c.apply(directions[i])]; touching {
				toAdd--
			}
		}
		if toAdd == 0 {
			_ = toAdd
		}
		sum += toAdd
	}
	return sum
}

const threshold = 2

func (s shape) OuterArea() int {
	var min, max cube
	min.x = math.MaxInt
	min.y = math.MaxInt
	min.z = math.MaxInt
	for c := range s {
		if c.x < min.x {
			min.x = c.x
		}
		if c.y < min.y {
			min.y = c.y
		}
		if c.z < min.z {
			min.z = c.z
		}
		if c.x > max.x {
			max.x = c.x
		}
		if c.y > max.y {
			max.y = c.y
		}
		if c.z > max.z {
			max.z = c.z
		}
	}

	min.x -= threshold
	min.y -= threshold
	min.z -= threshold

	max.x += threshold
	max.y += threshold
	max.z += threshold

	var visited = make(map[cube]struct{})
	var Q []cube
	Q = append(Q, min)
	visited[min] = struct{}{}

	var surfaceArea int
	for len(Q) > 0 {
		n := Q[0]
		Q = Q[1:]

		for i := range directions {
			nn := n.apply(directions[i])
			if _, ok := visited[nn]; ok {
				continue
			}

			if _, part := s[nn]; part {
				surfaceArea++
			} else {
				if nn.x > max.x || nn.y > max.y || nn.z > max.z {
					continue
				}

				if nn.x < min.x || nn.y < min.y || nn.z < min.z {
					continue
				}
				Q = append(Q, nn)
				visited[nn] = struct{}{}
			}
		}
	}

	return surfaceArea
}

func (s shape) Inside(c cube) bool {
	// If it's part of the shape we are not inside
	if _, has := s[c]; has {
		return false
	}

	var min, max cube
	for c := range s {
		if c.x < min.x {
			min.x = c.x - 1
		}
		if c.y < min.y {
			min.y = c.y - 1
		}
		if c.z < min.z {
			min.z = c.z - 1
		}
		if c.x > max.x {
			max.x = c.x + 1
		}
		if c.y > max.y {
			max.y = c.y + 1
		}
		if c.z > max.z {
			max.z = c.z + 1
		}
	}

	for i := range directions {
		var count int
		cc := c
		for {
			cc = cc.apply(directions[i])
			// Reached the edge
			if cc.x > max.x || cc.y > max.y || cc.z > max.z ||
				cc.x < min.x || cc.y < min.y || cc.z < min.z {
				break
			}

			if _, has := s[cc]; has {
				count++
				break
			}
		}
		if count == 0 {
			return false
		}
	}
	return true
}
