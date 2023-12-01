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

	scanner := bufio.NewScanner(f)
	switch *part {
	case 1:
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

const (
	decryptionKey = 811_589_153
)

func Part1(input *bufio.Scanner) int {
	var source = make([]*node, 0, 100)

	for input.Scan() {
		line := input.Text()
		n, _ := strconv.Atoi(line)
		source = append(source, &node{v: n})
	}

	for i := range source {
		source[i].next = source[(i+1)%len(source)]
		if i == 0 {
			source[i].prev = source[len(source)-1]
		} else {
			source[i].prev = source[(i-1)%len(source)]
		}
	}

	var z *node
	var m = len(source) - 1

	for _, k := range source {
		if k.v == 0 {
			z = k
			continue
		}
		p := k
		if k.v > 0 {
			for i := 0; i < (k.v % m); i++ {
				p = p.next
			}
			if k == p {
				continue
			}
			k.next.prev = k.prev
			k.prev.next = k.next
			p.next.prev = k
			k.next = p.next
			p.next = k
			k.prev = p
		} else {
			for i := 0; i < (-k.v % m); i++ {
				p = p.prev
			}
			if k == p {
				continue
			}
			k.prev.next = k.next
			k.next.prev = k.prev
			p.prev.next = k
			k.prev = p.prev
			p.prev = k
			k.next = p
		}
	}

	var sum int

	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			z = z.next
		}
		sum += z.v
	}

	return sum
}

func Part2(input *bufio.Scanner) int {
	var source = make([]*node, 0, 100)

	for input.Scan() {
		line := input.Text()
		n, _ := strconv.Atoi(line)
		source = append(source, &node{v: n * decryptionKey})
	}

	for i := range source {
		source[i].next = source[(i+1)%len(source)]
		if i == 0 {
			source[i].prev = source[len(source)-1]
		} else {
			source[i].prev = source[(i-1)%len(source)]
		}
	}

	var z *node
	var m = len(source) - 1

	for sh := 0; sh < 10; sh++ {
		for i := range source {
			k := source[i]
			if k.v == 0 {
				z = k
				continue
			}
			p := k
			if k.v > 0 {
				for jj := 0; jj < (k.v % m); jj++ {
					p = p.next
				}
				if k == p {
					continue
				}
				k.next.prev = k.prev
				k.prev.next = k.next
				p.next.prev = k
				k.next = p.next
				p.next = k
				k.prev = p
			} else {
				for jj := 0; jj < (-k.v % m); jj++ {
					p = p.prev
				}
				if k == p {
					continue
				}
				k.prev.next = k.next
				k.next.prev = k.prev
				p.prev.next = k
				k.prev = p.prev
				p.prev = k
				k.next = p
			}
		}
	}

	var sum int

	for i := 0; i < 3; i++ {
		for j := 0; i < 1000; j++ {
			z = z.next
		}
		sum += z.v
	}

	return sum
}

type node struct {
	v    int
	next *node
	prev *node
}
