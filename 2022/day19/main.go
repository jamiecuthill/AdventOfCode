package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	scanner := bufio.NewScanner(f)
	switch *part {
	case 1:
		fmt.Println(Part1(scanner))
	case 2:
		fmt.Println(Part2(scanner))
	}
}

type blueprint struct {
	plan [4]resources
	max  resources
}

var oreParser = regexp.MustCompile(`Each ore robot costs (\d+) ore`)
var clayParser = regexp.MustCompile(`Each clay robot costs (\d+) ore`)
var obsidianParser = regexp.MustCompile(`Each obsidian robot costs (\d+) ore and (\d+) clay`)
var geodeParser = regexp.MustCompile(`Each geode robot costs (\d+) ore and (\d+) obsidian`)

func parse(input *bufio.Scanner) []blueprint {
	var bps []blueprint
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, ": ")

		bp := blueprint{}
		matches := oreParser.FindSubmatch([]byte(parts[1]))
		oreCost, _ := strconv.Atoi(string(matches[1]))
		bp.plan[ore] = resources{ore: oreCost}
		bp.max[ore] = max(bp.max[ore], oreCost)

		matches = clayParser.FindSubmatch([]byte(parts[1]))
		clayOreCost, _ := strconv.Atoi(string(matches[1]))
		bp.plan[clay] = resources{ore: clayOreCost}
		bp.max[ore] = max(bp.max[ore], clayOreCost)

		matches = obsidianParser.FindSubmatch([]byte(parts[1]))
		obsidianOreCost, _ := strconv.Atoi(string(matches[1]))
		obsidianClayCost, _ := strconv.Atoi(string(matches[2]))
		bp.plan[obsidian] = resources{ore: obsidianOreCost, clay: obsidianClayCost}
		bp.max[ore] = max(bp.max[ore], obsidianOreCost)
		bp.max[clay] = max(bp.max[clay], obsidianClayCost)

		matches = geodeParser.FindSubmatch([]byte(parts[1]))
		geodeOreCost, _ := strconv.Atoi(string(matches[1]))
		geodeObsidianCost, _ := strconv.Atoi(string(matches[2]))
		bp.plan[geode] = resources{ore: geodeOreCost, obsidian: geodeObsidianCost}
		bp.max[ore] = max(bp.max[ore], geodeOreCost)
		bp.max[obsidian] = max(bp.max[obsidian], geodeObsidianCost)

		bps = append(bps, bp)
	}

	return bps
}

func Part1(input *bufio.Scanner) int {
	blueprints := parse(input)

	var total int
	for i, bp := range blueprints {
		v := dfs(24, bp, resources{ore: 1}, resources{}, map[key]int{})

		total += (i + 1) * v
	}

	return total
}

func Part2(input *bufio.Scanner) int {
	return 0
}

func dfs(time int, bp blueprint, robots resources, stash resources, cache map[key]int) int {
	if time == 0 {
		return stash[geode]
	}

	k := key{t: time, b: robots, s: stash}
	if maxval, ok := cache[k]; ok {
		return maxval
	}

	maxval := stash[geode] + robots[geode]*time

	for robot, costs := range bp.plan {
		if robot != geode && robots[robot] >= bp.max[robot] {
			continue
		}

		var timecost int
		var hasBot = true
		for r, v := range costs {
			if v == 0 {
				continue
			}
			// don't wait for resources for which we have no robots
			if robots[r] == 0 {
				hasBot = false
				break
			}
			timecost = max(timecost, -(-(v - stash[r]) / robots[r]))
		}
		if hasBot {
			remaining := time - timecost - 1
			if remaining <= 0 {
				continue
			}
			bots := robots
			amt := stash
			for typ := range bots {
				amt[typ] += bots[typ] * (timecost + 1)
			}
			for r, v := range costs {
				amt[r] -= v
			}
			bots[robot] += 1
			for i := 0; i < 3; i++ {
				amt[i] = min(amt[i], bp.max[i]*remaining)
			}
			maxval = max(maxval, dfs(remaining, bp, bots, amt, cache))
		}
	}
	cache[k] = maxval
	return maxval
}

type key struct {
	t int
	b resources
	s resources
}

type resources [4]int

const (
	ore int = iota
	clay
	obsidian
	geode
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
