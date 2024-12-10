package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"time"
)

var part = flag.Int("part", 1, "Run part 1 or part 2?")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	cap := 36

	monkeys := []*monkey{
		{
			items: (&ring{d: make([]int, cap)}).append(72).append(97), // []int{72, 97},
			op: func(i int) int {
				return i * 13
			},
			modulo: 19,
		},
		{
			items: (&ring{d: make([]int, cap)}).append(55).append(70).append(90).append(74).append(95), //[]int{55, 70, 90, 74, 95},
			op: func(i int) int {
				return i * i
			},
			modulo: 7,
		},
		{
			items: (&ring{d: make([]int, cap)}).append(74).append(97).append(66).append(57), //[]int{74, 97, 66, 57},
			op: func(i int) int {
				return i + 6
			},
			modulo: 17,
		},
		{
			items: (&ring{d: make([]int, cap)}).append(86).append(54).append(53), //[]int{86, 54, 53},
			op: func(i int) int {
				return i + 2
			},
			modulo: 13,
		},
		{
			items: (&ring{d: make([]int, cap)}).append(50).append(65).append(78).append(50).append(62).append(99), //[]int{50, 65, 78, 50, 62, 99},
			op: func(i int) int {
				return i + 3
			},
			modulo: 11,
		},
		{
			items: (&ring{d: make([]int, cap)}).append(90), //[]int{90},
			op: func(i int) int {
				return i + 4
			},
			modulo: 2,
		},
		{
			items: (&ring{d: make([]int, cap)}).append(88).append(92).append(63).append(94).append(96).append(82).append(53).append(53), //[]int{88, 92, 63, 94, 96, 82, 53, 53},
			op: func(i int) int {
				return i + 8
			},
			modulo: 5,
		},
		{
			items: (&ring{d: make([]int, cap)}).append(70).append(60).append(71).append(69).append(77).append(70).append(98), //[]int{70, 60, 71, 69, 77, 70, 98},
			op: func(i int) int {
				return i * 7
			},
			modulo: 3,
		},
	}

	monkeys[0].a = monkeys[5]
	monkeys[0].b = monkeys[6]

	monkeys[1].a = monkeys[5]
	monkeys[1].b = monkeys[0]

	monkeys[2].a = monkeys[1]
	monkeys[2].b = monkeys[0]

	monkeys[3].a = monkeys[1]
	monkeys[3].b = monkeys[2]

	monkeys[4].a = monkeys[3]
	monkeys[4].b = monkeys[7]

	monkeys[5].a = monkeys[4]
	monkeys[5].b = monkeys[6]

	monkeys[6].a = monkeys[4]
	monkeys[6].b = monkeys[7]

	monkeys[7].a = monkeys[2]
	monkeys[7].b = monkeys[3]

	switch *part {
	case 1:
		fmt.Println(Solve(monkeys, func(i int) int { return i / 3 }))
	case 2:
		var modulo = 1
		for _, m := range monkeys {
			modulo *= m.modulo
		}
		fmt.Println(Solve2(monkeys, func(i int) int {
			if i < modulo {
				return i
			}
			return i % modulo
		}))
	}
}

func Solve(monkeys []*monkey, reducer func(int) int) int {
	for round := 1; round <= 20; round++ {
		for i := range monkeys {
			monkeys[i].inspect(reducer)
		}
	}

	return monkeyBusinessLevel(monkeys)
}

func Solve2(monkeys []*monkey, reducer func(int) int) int {
	for round := 1; round <= 10000; round++ {
		for i := range monkeys {
			monkeys[i].inspect(reducer)
		}
	}

	return monkeyBusinessLevel(monkeys)
}

func monkeyBusinessLevel(monkeys []*monkey) int {
	var inspections []int

	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	return inspections[0] * inspections[1]
}

type monkey struct {
	items       *ring
	op          func(int) int
	inspections int
	modulo      int
	a           *monkey
	b           *monkey
}

func (m *monkey) inspect(relief func(int) int) {
	for m.items.len() > 0 {
		item := relief(m.op(m.items.pop()))
		m.throwTo(item).catch(item)
		m.inspections++
	}
}

func (m *monkey) throwTo(item int) *monkey {
	if item%m.modulo == 0 {
		return m.a
	}
	return m.b
}

func (m *monkey) catch(item int) {
	m.items.append(item)
}

type ring struct {
	d []int
	h int
	t int
	l int
}

func (r *ring) pop() int {
	i := r.d[r.h]
	r.h++
	if r.h >= len(r.d) {
		r.h = r.h % len(r.d)
	}
	r.l--
	return i
}

func (r *ring) append(i int) *ring {
	if r.l == len(r.d) {
		panic("reached capacity")
	}
	r.d[r.t] = i
	r.l++
	r.t++
	if r.t >= len(r.d) {
		r.t = r.t % len(r.d)
	}
	return r
}

func (r ring) len() int {
	return r.l
}
