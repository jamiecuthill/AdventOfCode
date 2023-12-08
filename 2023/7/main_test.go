package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

const example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 6440, Part1(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 5905, Part2(bufio.NewScanner(strings.NewReader(example))))
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 241344943, Part1(input))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewScanner(f)

	assert.Equal(t, 243101568, Part2(input))
}

func TestJokerEdgeCases(t *testing.T) {
	assert.Equal(t, 5, handTypeRankWithJokers(&Hand{[]int{10, 5, 5, 1, 5}, 0}))
	assert.Equal(t, 5, handTypeRankWithJokers(&Hand{[]int{13, 10, 1, 1, 10}, 0}))
	assert.Equal(t, 5, handTypeRankWithJokers(&Hand{[]int{12, 12, 12, 1, 14}, 0}))
	assert.Equal(t, 5, handTypeRankWithJokers(&Hand{[]int{13, 10, 1, 1, 10}, 0}))
	assert.Equal(t, 3, handTypeRankWithJokers(&Hand{[]int{2, 3, 1, 1, 4}, 0}))

	hands := []Hand{
		{[]int{13, 10, 1, 1, 10}, 2},
		{[]int{12, 12, 12, 1, 14}, 1},
	}
	OrderedBy(handTypeWithJokers, handCards).Sort(hands)
	assert.Equal(t, 1, hands[0].score)
	assert.Equal(t, 2, hands[1].score)

	// Try and figure this out
	assert.Equal(t, 1, handTypeRankWithJokers(&Hand{[]int{1, 2, 3, 4, 5}, 0}), "One Joker, is one pair")
	assert.Equal(t, 3, handTypeRankWithJokers(&Hand{[]int{1, 2, 2, 3, 4}, 0}), "One Joker with one pair, is three of a kind")
	assert.Equal(t, 5, handTypeRankWithJokers(&Hand{[]int{1, 2, 2, 2, 3}, 0}), "One Joker with three of a kind, is four of a kind")
	assert.Equal(t, 6, handTypeRankWithJokers(&Hand{[]int{1, 2, 2, 2, 2}, 0}), "One Joker with four of a kind, is five of a kind")

	assert.Equal(t, 3, handTypeRankWithJokers(&Hand{[]int{1, 1, 2, 3, 4}, 0}), "Two Joker, is three of a kind")
	assert.Equal(t, 5, handTypeRankWithJokers(&Hand{[]int{1, 1, 2, 2, 3}, 0}), "Two Joker with one pair, is four of a kind")
	assert.Equal(t, 6, handTypeRankWithJokers(&Hand{[]int{1, 1, 2, 2, 2}, 0}), "Two Joker with three of a kind, is five of a kind")
}
