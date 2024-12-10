package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var example = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 1651, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 1707, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 4886370, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 11374534948438, Part2(input))
}

// func TestTunnel(t *testing.T) {
// 	p := tunnel{
// 		valve: valve{
// 			name: "C",
// 			rate: 10,
// 		},
// 		previous: &tunnel{
// 			valve: valve{
// 				name:    "B",
// 				rate:    20,
// 				tunnels: []string{"C"},
// 			},
// 			previous: &tunnel{
// 				valve: valve{
// 					name:    "A",
// 					rate:    50,
// 					tunnels: []string{"B"},
// 				},
// 			},
// 		},
// 	}

// 	assert.Equal(t, 3, p.len())
// 	assert.Equal(t, 200, p.totalValue())
// 	assert.Equal(t, "B", p.nextMove())
// }
