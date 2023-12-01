package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	scanner := bufio.NewScanner(f)
	switch *part {
	case 1:
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

var numberParser = regexp.MustCompile(`([a-z]+): ([0-9]+)`)
var calcParser = regexp.MustCompile(`([a-z]+): ([a-z]+) ([-+\*/]) ([a-z]+)`)

func Part1(input *bufio.Scanner) int {
	var index = make(map[string]tree)
	for input.Scan() {
		line := input.Bytes()

		numberMatches := numberParser.FindSubmatch(line)
		if len(numberMatches) > 0 {
			num, _ := strconv.Atoi(string(numberMatches[2]))
			index[string(numberMatches[1])] = &leaf{
				name:  string(numberMatches[1]),
				shout: num,
			}
			continue
		}

		calcMatches := calcParser.FindSubmatch(line)
		if len(calcMatches) > 0 {
			index[string(calcMatches[1])] = &calculation{
				name:     string(calcMatches[1]),
				op:       rune(string(calcMatches[3])[0]),
				leftKey:  string(calcMatches[2]),
				rightKey: string(calcMatches[4]),
			}
			continue
		}

		panic("No match: " + string(line))
	}

	root := buildTree(index)

	return root.Value()
}

func Part2(input *bufio.Scanner) int {
	// var index = make(map[string]tree)
	// for input.Scan() {
	// 	line := input.Bytes()

	// 	numberMatches := numberParser.FindSubmatch(line)
	// 	if len(numberMatches) > 0 {
	// 		if string(numberMatches[1]) == humanKey {
	// 			index[string(numberMatches[1])] = &human{
	// 				name:  string(numberMatches[1]),
	// 				shout: 0,
	// 			}
	// 			continue
	// 		}
	// 		num, _ := strconv.Atoi(string(numberMatches[2]))
	// 		index[string(numberMatches[1])] = &leaf{
	// 			name:  string(numberMatches[1]),
	// 			shout: num,
	// 		}
	// 		continue
	// 	}

	// 	calcMatches := calcParser.FindSubmatch(line)
	// 	if len(calcMatches) > 0 {
	// 		op := rune(string(calcMatches[3])[0])
	// 		if string(calcMatches[1]) == rootKey {
	// 			op = '='
	// 		}
	// 		index[string(calcMatches[1])] = &calculation{
	// 			name:     string(calcMatches[1]),
	// 			op:       op,
	// 			leftKey:  string(calcMatches[2]),
	// 			rightKey: string(calcMatches[4]),
	// 		}
	// 		continue
	// 	}

	// 	panic("No match: " + string(line))
	// }

	// root := buildTree(index)

	// target := root.Target(0)
	// fmt.Println(target)
	// fmt.Println(root.Value())
	// return target
	var index = make(map[string]tree)
	for input.Scan() {
		line := input.Bytes()

		numberMatches := numberParser.FindSubmatch(line)
		if len(numberMatches) > 0 {
			if _, ok := index[string(numberMatches[1])]; ok {
				panic("not unique")
			}
			if string(numberMatches[1]) == humanKey {
				index[string(numberMatches[1])] = &human{
					name:  string(numberMatches[1]),
					shout: 0,
				}
				continue
			}
			num, _ := strconv.Atoi(string(numberMatches[2]))
			index[string(numberMatches[1])] = &leaf{
				name:  string(numberMatches[1]),
				shout: num,
			}
			continue
		}

		calcMatches := calcParser.FindSubmatch(line)
		if len(calcMatches) > 0 {
			if _, ok := index[string(calcMatches[1])]; ok {
				panic("not unique")
			}
			op := rune(string(calcMatches[3])[0])
			if string(calcMatches[1]) == rootKey {
				op = '='
			}
			index[string(calcMatches[1])] = &calculation{
				name:     string(calcMatches[1]),
				op:       op,
				leftKey:  string(calcMatches[2]),
				rightKey: string(calcMatches[4]),
			}
			continue
		}

		panic("No match: " + string(line))
	}

	root := buildTree(index)
	return root.Target(0)
}

const (
	rootKey  = "root"
	humanKey = "humn"

	plus     = '+'
	minus    = '-'
	divide   = '/'
	multiply = '*'
	equals   = '='
)

type tree interface {
	Value() int
	Name() string
	Target(int) int
	ContainsHuman() bool
}

type human struct {
	name  string
	shout int
}

func (l human) Value() int {
	return l.shout
}

func (l human) Name() string {
	return l.name
}

func (l *human) Target(target int) int {
	l.shout = target
	return l.shout
}

func (l human) ContainsHuman() bool {
	return true
}

type leaf struct {
	name  string
	shout int
}

func (l leaf) Value() int {
	return l.shout
}

func (l leaf) Name() string {
	return l.name
}

func (l *leaf) Target(target int) int {
	return l.shout
}

func (l leaf) ContainsHuman() bool {
	return false
}

type calculation struct {
	name     string
	op       rune
	leftKey  string
	rightKey string
	left     tree
	right    tree
}

func (c calculation) Name() string {
	return c.name
}

func (c calculation) Value() int {
	switch c.op {
	case equals:
		return c.left.Value() - c.right.Value()
	case plus:
		return c.left.Value() + c.right.Value()
	case minus:
		return c.left.Value() - c.right.Value()
	case divide:
		return c.left.Value() / c.right.Value()
	case multiply:
		return c.left.Value() * c.right.Value()
	default:
		panic("Unknown op " + string(c.op))
	}
}

func (c calculation) ContainsHuman() bool {
	return c.left.ContainsHuman() || c.right.ContainsHuman()
}

func (c calculation) Target(t int) int {
	var v, x int
	var side tree
	var leftSide = c.left.ContainsHuman()
	if leftSide {
		v = c.right.Value()
		side = c.left
	} else {
		v = c.left.Value()
		side = c.right
	}

	switch c.op {
	case equals:
		x = v
	case plus:
		// if leftSide {
		x = t - v
		// } else {
		// x = t - v
		// }
	case minus:
		if leftSide {
			x = t + v
		} else {
			x = v - t
		}
	case divide:
		if leftSide {
			x = t * v
		} else {
			x = v / t
		}
	case multiply:
		x = t / v
	default:
		panic("Unknown op " + string(c.op))
	}

	// if leftSide {
	// 	fmt.Printf("%s: (%d) %c %d\n", c.name, x, c.op, v)
	// } else {
	// 	fmt.Printf("%s: %d %c (%d)\n", c.name, v, c.op, x)
	// }

	return side.Target(x)
}

// buildTree builds a tree and returns the root node
func buildTree(index map[string]tree) tree {
	var root tree

	for k, n := range index {
		if k == rootKey {
			root = n
		}

		switch nn := n.(type) {
		case *calculation:
			nn.left = index[nn.leftKey]
			nn.right = index[nn.rightKey]
		case *leaf:
		}
	}

	return root
}
