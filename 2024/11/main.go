package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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

const factor = 2024
const maxBlinks = 100

func parse(input *bufio.Scanner) []int {
	input.Scan()
	line := input.Text()

	var stones []int
	for _, v := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(v)
		stones = append(stones, n)
	}

	return stones
}

var cache = new(sync.Map)

func Blink(value int, times int) int {
	if times > maxBlinks {
		panic("Too many Blinks requested")
	}

	var cacheKey = maxBlinks*value + times

	if v, ok := cache.Load(cacheKey); ok {
		return v.(int)
	}

	if times == 1 {
		if value == 0 {
			return 1
		}
		var digits = int(math.Log10(float64(value)) + 1)
		if digits%2 == 0 {
			return 2
		}
		return 1
	}

	if value == 0 {
		len := Blink(1, times-1)
		cache.Store(cacheKey, len)
		return len
	}

	var len int
	var digits = int(math.Log10(float64(value)) + 1)
	if digits%2 == 0 {
		var div = 1

		for i := 0; i < digits/2; i++ {
			div *= 10
		}

		len += Blink(value/div, times-1)
		len += Blink(value%div, times-1)
		cache.Store(cacheKey, len)
		return len
	}

	len += Blink(value*factor, times-1)
	cache.Store(cacheKey, len)
	return len
}

func Run(stones []int, times int) int {
	var res = make(chan int, len(stones))
	wg := sync.WaitGroup{}
	wg.Add(len(stones))

	var length int
	for i := range stones {
		go func(stone int) {
			defer wg.Done()
			res <- Blink(stone, times)
		}(stones[i])
	}
	wg.Wait()
	close(res)
	for r := range res {
		length += r
	}
	return length
}

func Part1(input *bufio.Scanner) int {
	return Run(parse(input), 25)
}

func Part2(input *bufio.Scanner) int {
	return Run(parse(input), 75)
}
