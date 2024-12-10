package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	limit      = 100000
	diskTotal  = 70000000
	diskTarget = 30000000
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var root = &node{name: "/"}
	var cwd *node

	var mode = cmdNull

	for scanner.Scan() {
		line := scanner.Text()
		switch string(line[0]) {
		// Command
		case "$":
			mode = cmdRun
			if strings.HasPrefix(line, "$ cd") {
				switch strings.TrimPrefix(line, "$ cd ") {
				case "/":
					cwd = root
				case "..":
					cwd = cwd.parent
				default:

					cwd = cwd.Cd(strings.TrimPrefix(line, "$ cd "))
				}
			}
			if strings.HasPrefix(line, "$ ls") {
				// TODO expect output
				mode = cmdList
			}
		// Output
		default:
			if mode != cmdList {
				panic("Not in list mode but got output")
			}
			if strings.HasPrefix(line, "dir") {
				dir := &node{name: strings.TrimPrefix(line, "dir ")}
				cwd.HasDir(dir)
			} else {
				parts := strings.Split(line, " ")
				sizeInt, _ := strconv.Atoi(parts[0])
				cwd.HasFile(file{name: parts[1], size: sizeInt})
			}
		}
	}

	root.Print("")

	fmt.Println("Sum at most: ", root.SumAtMost(limit, 0))

	fmt.Println("Space used:", root.totalSize())
	diskUnused := diskTotal - root.totalSize()
	fmt.Println("Unused space:", diskUnused)
	fmt.Println("Space to find:", diskTarget-diskUnused)

	fmt.Println("Delete to free:", root.FindTarget(diskTarget-diskUnused, diskTotal))
}

const (
	cmdNull int = iota
	cmdRun
	cmdList
)

type node struct {
	name   string
	parent *node
	dirs   []*node
	files  []file
}

func (n *node) HasDir(d *node) {
	n.dirs = append(n.dirs, d)
	d.parent = n
}

func (n *node) Cd(d string) *node {
	for _, c := range n.dirs {
		if c.name == d {
			return c
		}
	}
	panic("dir not found: " + d)
	// return nil
}

func (n *node) HasFile(f file) {
	n.files = append(n.files, f)
}

func (n node) Print(indent string) {
	fmt.Printf(indent+"%s\n", n.name)

	for _, c := range n.dirs {
		c.Print(indent + " ")
	}

	for _, f := range n.files {
		f.Print(indent)
	}
}

func (n node) SumAtMost(limit int, sum int) int {
	if thisSum := n.totalSize(); thisSum <= limit {
		sum += thisSum
	}
	for _, d := range n.dirs {
		sum = d.SumAtMost(limit, sum)
	}
	return sum
}

func (n node) totalSize() int {
	var sum int
	for _, f := range n.files {
		sum += f.size
	}
	for _, d := range n.dirs {
		sum += d.totalSize()
	}
	return sum
}

func (n node) FindTarget(target int, current int) int {
	for _, d := range n.dirs {
		size := d.totalSize()
		if size >= target && size < current {
			current = size
		}
		current = d.FindTarget(target, current)
	}
	return current
}

type file struct {
	name string
	size int
}

func (f file) Print(indent string) {
	fmt.Printf(indent+"%s: %d\n", f.name, f.size)
}
