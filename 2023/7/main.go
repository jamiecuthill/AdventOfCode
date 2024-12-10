package main

import (
	"bufio"
	"flag"
	"fmt"
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

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

type Hand struct {
	cards []int
	score int
}

func Part1(input *bufio.Scanner) int {
	var hands []Hand

	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, " ")
		cards := make([]int, len(parts[0]))
		for i, c := range parts[0] {
			cards[i] = int(c) - '0'
			if c == 'T' {
				cards[i] = 10
			} else if c == 'J' {
				cards[i] = 11
			} else if c == 'Q' {
				cards[i] = 12
			} else if c == 'K' {
				cards[i] = 13
			} else if c == 'A' {
				cards[i] = 14
			}
		}
		s, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{cards, s})
	}

	OrderedBy(handType, handCards).Sort(hands)

	return score(hands)
}

func Part2(input *bufio.Scanner) int {
	var hands []Hand

	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, " ")
		cards := make([]int, len(parts[0]))
		for i, c := range parts[0] {
			cards[i] = int(c) - '0'
			if c == 'T' {
				cards[i] = 10
			} else if c == 'J' {
				cards[i] = 1
			} else if c == 'Q' {
				cards[i] = 12
			} else if c == 'K' {
				cards[i] = 13
			} else if c == 'A' {
				cards[i] = 14
			}
		}
		s, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{cards, s})
	}

	OrderedBy(handTypeWithJokers, handCards).Sort(hands)

	return score(hands)
}

func score(hands []Hand) int {
	var sum int
	for i, h := range hands {
		sum += (i + 1) * h.score
	}
	return sum
}

func handTypeWithJokers(p, q *Hand) int {
	r1, r2 := handTypeRankWithJokers(p), handTypeRankWithJokers(q)
	if r1 > r2 {
		return 1
	} else if r1 < r2 {
		return -1
	}
	return 0
}

func handType(p, q *Hand) int {
	r1, r2 := handTypeRank(p), handTypeRank(q)
	if r1 > r2 {
		return 1
	} else if r1 < r2 {
		return -1
	}
	return 0
}

func handTypeRank(h *Hand) int {
	var cSet = map[int]int{}

	for _, c := range h.cards {
		cSet[c]++
	}

	return rankFromSet(cSet)
}

func handTypeRankWithJokers(h *Hand) int {
	var cSet = map[int]int{}

	for _, c := range h.cards {
		cSet[c]++
	}

	r := rankFromSet(cSet)

	if jokerCount, ok := cSet[1]; ok {
		// Has at least one joker
		if jokerCount == 1 {
			switch r {
			case 0:
				return 1
			case 5:
				return 6
			default:
				return r + 2
			}
		} else if jokerCount == 2 {
			switch r {
			case 1:
				return 3
			case 2:
				return 5
			case 4:
				return 6
			}
		} else if jokerCount == 3 {
			switch r {
			case 3:
				return 5
			case 4:
				return 6
			}
		} else if jokerCount == 4 {
			switch r {
			case 5:
				return 6
			}
		}
	}

	return r
}

func rankFromSet(cSet map[int]int) int {
	switch len(cSet) {
	case 5: // High card
		return 0
	case 4: // One pair
		return 1
	case 3: // Two pair or three of a kind
		for _, v := range cSet {
			if v == 2 {
				return 2
			}
		}
		// Three of a kind
		return 3
	case 2: // Full house or Four of a kind
		for _, v := range cSet {
			if v == 3 {
				return 4
			}
		}
		// Four of a kind
		return 5
	case 1: // Five of a kind
		return 6
	}
	return 0
}

func handCards(p, q *Hand) int {
	for i := 0; i < len(p.cards); i++ {
		if p.cards[i] > q.cards[i] {
			return 1
		} else if p.cards[i] < q.cards[i] {
			return -1
		}
	}
	return 0
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{less: less}
}

type lessFunc func(p1, p2 *Hand) int

type multiSorter struct {
	hands []Hand
	less  []lessFunc
}

func (ms *multiSorter) Sort(hands []Hand) {
	ms.hands = hands
	sort.Sort(ms)
}

func (ms *multiSorter) Len() int {
	return len(ms.hands)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.hands[i], ms.hands[j] = ms.hands[j], ms.hands[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.hands[i], &ms.hands[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch less(p, q) {
		case -1:
			return true
		case 1:
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q) < 0
}
