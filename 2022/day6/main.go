package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed input
var input []byte

// const markerSize = 4
const markerSize = 14

func main() {
	var buf []rune

	for i, c := range string(input) {
		buf = appendChar(c, buf)
		if isUnique(buf) {
			fmt.Println(i + 1)
			os.Exit(0)
		}
	}
}

func isUnique(in []rune) bool {
	var set = make(map[rune]struct{})

	for _, c := range in {
		set[c] = struct{}{}
	}
	return len(set) == markerSize
}

func appendChar(c rune, buf []rune) []rune {
	if len(buf) < markerSize {
		return append(buf, c)
	}

	return append(buf, c)[1:]
}
