package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
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

	input := bufio.NewScanner(f)

	switch *part {
	case 1:
		fmt.Println(Part1(input))
	case 2:
		fmt.Println(Part2(input))
	}
}

func parse(input *bufio.Scanner) []int {
	var id int

	input.Scan()
	line := input.Text()

	var disk = make([]int, 0, len(line))

	var isFile = true
	for i := range line {
		v, _ := strconv.Atoi(string(line[i]))
		var blockValue = -1
		if isFile {
			blockValue = id
			id++
		}
		for j := 0; j < v; j++ {
			disk = append(disk, blockValue)
		}
		isFile = !isFile
	}

	return disk
}

func Part1(input *bufio.Scanner) int {
	disk := parse(input)

	var target int

	for i := len(disk) - 1; i >= 0; i-- {
		// Continue until we find a non-empty block
		if disk[i] == -1 {
			continue
		}

		// Scan to next free space
		for disk[target] >= 0 {
			target++
		}

		// Done when target is at or beyond the current block
		if target >= i {
			break
		}

		// Move the block to the target
		disk[target] = disk[i]
		disk[i] = -1
	}

	return checksum(disk, true)
}

func checksum(disk []int, contiguous bool) int {
	var sum int
	for i := range disk {
		if disk[i] == -1 {
			if contiguous {
				break
			}
			continue
		}
		sum += i * disk[i]

	}
	return sum
}

func Part2(input *bufio.Scanner) int {
	disk := parse(input)

	fileID := disk[len(disk)-1]

	for fileID > 0 {
		// Find the file and its size
		filePosition := slices.Index(disk, fileID)
		blockSize := 1
		for filePosition+blockSize < len(disk) && disk[filePosition+blockSize] == fileID {
			blockSize++
		}

		var i int

		// only search left of the file index
		for i < filePosition {
			// Scan forward to a free space
			if disk[i] != -1 {
				i++
				continue
			}

			// Look ahead for enough free space for the current block
			var haveCapacity bool
			for n := 0; n < blockSize; n++ {
				if disk[i+n] != -1 {
					haveCapacity = false
					break
				}
				haveCapacity = true
			}

			if !haveCapacity {
				i++
				continue
			}

			// Move the file if we found enough space
			for j := 0; j < blockSize; j++ {
				disk[i+j] = disk[filePosition+j]
				disk[filePosition+j] = -1
			}
			break
		}

		fileID--
	}

	return checksum(disk, false)
}
