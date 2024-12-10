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

type rng struct {
	start int
	end   int
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

	seedInts := make([]int, len(seeds))
	for i, s := range seeds {
		seedInts[i], _ = strconv.Atoi(s)
	}

	var ranges = []rng{}

	for i := 0; i < len(seedInts); i += 2 {
		ranges = append(ranges, rng{start: seedInts[i], end: seedInts[i] + seedInts[i+1] - 1})
	}

	ranges = traverseRange(almanac["seed-to-soil"], ranges)
	ranges = traverseRange(almanac["soil-to-fertilizer"], ranges)
	ranges = traverseRange(almanac["fertilizer-to-water"], ranges)
	ranges = traverseRange(almanac["water-to-light"], ranges)
	ranges = traverseRange(almanac["light-to-temperature"], ranges)
	ranges = traverseRange(almanac["temperature-to-humidity"], ranges)
	ranges = traverseRange(almanac["humidity-to-location"], ranges)

	var locationMin = math.MaxInt

	for _, r := range ranges {
		if r.start < locationMin {
			locationMin = r.start
		}
	}

	return locationMin
}

func traverseRange(via []def, ranges []rng) []rng {
	var new = []rng{}

	// detect breaks in ranges
	for i := range ranges {
		for _, d := range via {
			r := ranges[i]
			if r.start >= d.src && r.start < d.src+d.len {
				if d.src+d.len-1 < r.end {
					ranges[i].end = d.src + d.len - 1
					ranges = append(ranges,
						rng{start: d.src + d.len, end: r.end},
					)
				}
			}
		}
	}

	// map ranges to new values
	for _, r := range ranges {
		var found bool
		for _, d := range via {
			if r.start >= d.src && r.start < d.src+d.len {
				new = append(new, rng{
					start: r.start + (d.dest - d.src),
					end:   r.end + (d.dest - d.src),
				})
				found = true
				break
			}
		}
		if !found {
			new = append(new, rng{start: r.start, end: r.end})
		}
	}

	return new
}
