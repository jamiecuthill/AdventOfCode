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

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

type def struct {
	dest, src, len int
}

func Part1(input *bufio.Scanner) int {
	var seeds []string
	var mapType string
	var almanac = map[string][]def{}

	for input.Scan() {
		line := input.Text()

		if line == "" {
			mapType = ""
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			seeds = strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
		}

		if strings.HasSuffix(line, " map:") {
			// seed-to-soil map:
			mapType = strings.TrimSuffix(line, " map:")
			continue
		}

		if mapType != "" {
			var d def
			vals := strings.Split(line, " ")
			d.dest, _ = strconv.Atoi(vals[0])
			d.src, _ = strconv.Atoi(vals[1])
			d.len, _ = strconv.Atoi(vals[2])
			almanac[mapType] = append(almanac[mapType], d)
		}
	}

	var locationMin = math.MaxInt

	for _, seed := range seeds {
		seedVal, _ := strconv.Atoi(seed)

		soil := traverse(almanac["seed-to-soil"], seedVal)
		fertilizer := traverse(almanac["soil-to-fertilizer"], soil)
		water := traverse(almanac["fertilizer-to-water"], fertilizer)
		light := traverse(almanac["water-to-light"], water)
		temperature := traverse(almanac["light-to-temperature"], light)
		humidity := traverse(almanac["temperature-to-humidity"], temperature)
		location := traverse(almanac["humidity-to-location"], humidity)
		// fmt.Printf("Seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d.\n", seedVal, soil, fertilizer, water, light, temperature, humidity, location)
		if location < locationMin {
			locationMin = location
		}
	}

	return locationMin
}

func traverse(via []def, src int) int {
	var dest = src
	for _, d := range via {
		if src >= d.src && src < d.src+d.len {
			i := src - d.src
			dest = d.dest + i
			break
		}
	}
	return dest
}

func Part2(input *bufio.Scanner) int {
	var seeds []string
	var mapType string
	var almanac = map[string][]def{}

	for input.Scan() {
		line := input.Text()

		if line == "" {
			mapType = ""
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			seeds = strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
		}

		if strings.HasSuffix(line, " map:") {
			// seed-to-soil map:
			mapType = strings.TrimSuffix(line, " map:")
			continue
		}

		if mapType != "" {
			var d def
			vals := strings.Split(line, " ")
			d.dest, _ = strconv.Atoi(vals[0])
			d.src, _ = strconv.Atoi(vals[1])
			d.len, _ = strconv.Atoi(vals[2])
			almanac[mapType] = append(almanac[mapType], d)
		}
	}

	var locationMin = math.MaxInt

	seedInts := make([]int, len(seeds))
	for i, s := range seeds {
		seedInts[i], _ = strconv.Atoi(s)
	}

	for i := 0; i < len(seedInts); i += 2 {
		seed := seedInts[i]
		len := seedInts[i+1]
		for s := seed; s < seed+len; s++ {
			soil := traverse(almanac["seed-to-soil"], s)
			fertilizer := traverse(almanac["soil-to-fertilizer"], soil)
			water := traverse(almanac["fertilizer-to-water"], fertilizer)
			light := traverse(almanac["water-to-light"], water)
			temperature := traverse(almanac["light-to-temperature"], light)
			humidity := traverse(almanac["temperature-to-humidity"], temperature)
			location := traverse(almanac["humidity-to-location"], humidity)
			// fmt.Printf("Seed %d, soil %d, fertilizer %d, water %d, light %d, temperature %d, humidity %d, location %d.\n", seed, soil, fertilizer, water, light, temperature, humidity, location)
			if location < locationMin {
				locationMin = location
			}
		}
	}

	return locationMin
}

// 79-92
// 50-97
// 52-99
// split the input range into ranges of the source mappings
// for each source range convert to a destination range
// return the destination ranges
// pass into the next traverse
func traverseRng(via []def, src, len int) (int, int) {
	var dest = src
	// var destLen = len
	for _, d := range via {
		if src >= d.src && src < d.src+d.len {
			i := src - d.src
			dest = d.dest + i
			break
		}
	}
	return dest, 0
}
