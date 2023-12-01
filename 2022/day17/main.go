package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

var blocks = [][]byte{
	{0b00111100},
	{0b00010000, 0b00111000, 0b00010000},
	{0b00111000, 0b00001000, 0b00001000},
	{0b00100000, 0b00100000, 0b00100000, 0b00100000},
	{0b00110000, 0b00110000},
}

const (
	chamberWidth = 7
	appearX      = 2
	appearY      = 3
	part1Rocks   = 2022
	part2Rocks   = 1000000000000
	// part2Rocks       = 100
	empty            = 0b00000000
	tetris           = 0b11111110
	againstRightWall = 0b00000010
	debug            = false
)

var (
	left  = byte('<')
	right = byte('>')
)

func Part1(input *bufio.Scanner) int {
	input.Scan()
	operations := input.Bytes()
	var board = make([]byte, 0, 1000)

	var nextBlock int
	var nextOp int

	var blockBuff = make([]byte, 4)

	for numBlocks := 1; numBlocks <= part1Rocks; numBlocks++ {
		if debug {
			fmt.Println("A new rock begins falling")
		}
		blockY := len(board) + appearY

		currentBlock := nextBlock

		// Make a copy of the blockBuff so that we can move it around
		blockBuff = blockBuff[0:len(blocks[currentBlock])]
		copy(blockBuff, blocks[currentBlock])

		// render(board, blockBuff, blockY)

		var landed bool

		for !landed {
			// fmt.Println("Block at row", blockY)

			// jet operation
			shift(operations[nextOp], blockY, blockBuff, board)

			// increment operation
			nextOp++
			if nextOp == len(operations) {
				nextOp = 0
			}

			// render(board, blockBuff, blockY)

			// fall
			if debug {
				fmt.Println("Rock falls 1 unit")
			}
			blockY--
			var shouldLand bool

			if blockY < 0 {
				shouldLand = true
			}

			if !shouldLand && blockY < len(board) {
				// fmt.Println("Checking collisions on row", blockY)

				for i := 0; i < len(blockBuff) && i+blockY < len(board); i++ {
					if board[blockY+i]&blockBuff[i] != empty {
						// fmt.Println("Detected collision at", blockY)
						shouldLand = true
						break
					}
				}
			}

			if shouldLand {
				blockY++
				if debug {
					fmt.Println(", causing it to come to rest")
				}
				neededLen := (blockY) + (len(blockBuff))
				if neededLen > len(board) {
					tmp := make([]byte, neededLen-len(board))
					board = append(board, tmp...)
				}
				for i := 0; i < len(blockBuff); i++ {
					board[blockY+i] = board[blockY+i] | blockBuff[i]
				}
				landed = true
				// render(board, nil, 0)
			} else {
				// render(board, blockBuff, blockY)
			}
		}

		// increment next block
		nextBlock++
		if nextBlock == len(blocks) {
			nextBlock = 0
		}
	}

	return len(board)
}

func Part2(input *bufio.Scanner) uint64 {
	input.Scan()
	operations := input.Bytes()
	var board []byte
	var optimizedLen uint64

	var nextBlock int
	var nextOp int

	var blockBuff = make([]byte, 4)

	for numBlocks := 1; numBlocks <= part2Rocks; numBlocks++ {
		if debug {
			fmt.Println("A new rock begins falling")
		}
		blockY := len(board) + appearY

		currentBlock := nextBlock

		// Make a copy of the blockBuff so that we can move it around
		blockBuff = blockBuff[0:len(blocks[currentBlock])]
		copy(blockBuff, blocks[currentBlock])

		// render(board, blockBuff, blockY)

		var landed bool

		for !landed {
			// fmt.Println("Block at row", blockY)

			// jet operation
			shift(operations[nextOp], blockY, blockBuff, board)
			// increment operation
			nextOp++
			if nextOp == len(operations) {
				nextOp = 0
			}

			// render(board, blockBuff, blockY)

			// fall
			if debug {
				fmt.Println("Rock falls 1 unit")
			}
			blockY--
			var shouldLand bool

			if blockY < 0 {
				shouldLand = true
			}

			if !shouldLand && blockY < len(board) {
				// fmt.Println("Checking collisions on row", blockY)

				for i := 0; i < len(blockBuff) && i+blockY < len(board); i++ {
					if board[blockY+i]&blockBuff[i] != empty {
						// fmt.Println("Detected collision at", blockY)
						shouldLand = true
						break
					}
				}
			}

			if shouldLand {
				blockY++
				if debug {
					fmt.Println(", causing it to come to rest")
				}
				neededLen := (blockY) + (len(blockBuff))
				if neededLen > len(board) {
					tmp := make([]byte, neededLen-len(board))
					board = append(board, tmp...)
				}
				for i := 0; i < len(blockBuff); i++ {
					board[blockY+i] = board[blockY+i] | blockBuff[i]
				}
				landed = true
				// render(board, nil, 0)
			} else {
				// render(board, blockBuff, blockY)
			}
		}

		// increment next block
		nextBlock++
		if nextBlock == len(blocks) {
			nextBlock = 0
		}

		// Optimise
		var row byte
		for i := len(board) - 1; i >= 0; i-- {
			row = row | board[i]
			if row == tetris {
				optimizedLen += uint64(i + 1)
				board = board[i:]
				break
			}
		}
	}

	return optimizedLen + uint64(len(board))
}

func shift(op byte, blockY int, blockBuff []byte, board []byte) {
	switch op {
	case left:
		if debug {
			fmt.Print("Jets pushing block left")
		}
		var collision bool
		for i := 0; i < len(blockBuff); i++ {
			row := blockBuff[i]
			// detect overflow
			if (row<<1)>>1 != row {
				collision = true
				break
			}
			if blockY+i >= len(board) {
				continue
			}
			rowmoved := row << 1
			boardrow := board[blockY+i]
			if boardrow&rowmoved != empty {
				collision = true
				break
			}
		}
		if collision {
			if debug {
				fmt.Println(", but nothing happens")
			}
			break
		}
		if debug {
			fmt.Println("")
		}
		// Shift the block
		for i := 0; i < len(blockBuff); i++ {
			blockBuff[i] = blockBuff[i] << 1
		}

	case right:
		if debug {
			fmt.Print("Jets pushing block right")
		}
		var collision bool
		for i := 0; i < len(blockBuff); i++ {
			row := blockBuff[i]
			if row&againstRightWall != empty {
				collision = true
				break
			}
			if blockY+i >= len(board) {
				continue
			}
			rowmoved := row >> 1
			boardrow := board[blockY+i]
			if boardrow&rowmoved != empty {
				collision = true
				break
			}
		}
		if collision {
			if debug {
				fmt.Println(", but nothing happens")
			}
			break
		}
		if debug {
			fmt.Println("")
		}
		for i := 0; i < len(blockBuff); i++ {
			blockBuff[i] = blockBuff[i] >> 1
		}
	}
}

func render(board []byte, falling []byte, at int) {
	start := len(board) - 1
	if at+len(falling)-1 > start {
		start = at + len(falling) - 1
	}

	for y := start; y >= 0; y-- {
		var row byte
		var fallRow byte
		if y >= len(board) {
			// no board to render
		} else {
			row = board[y]
		}

		if y-at >= 0 && (y-at) < len(falling) {
			fallRow = falling[y-at]
		}

		fmt.Print("|")
		for pixel := byte(0b10000000); pixel != 0b00000001; pixel = pixel >> 1 {
			if pixel&fallRow != empty {
				fmt.Print("@")
			} else if pixel&row != empty {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}

	// if len(falling) > 0 {
	// 	for i := len(falling) - 1; i >= 0; i-- {
	// 		row := falling[i]
	// 		fmt.Print("|")
	// 		for pixel := byte(0b10000000); pixel != 0b00000001; pixel = pixel >> 1 {
	// 			if pixel&row != empty {
	// 				fmt.Print("@")
	// 			} else {
	// 				fmt.Print(".")
	// 			}
	// 		}
	// 		fmt.Println("|")
	// 	}
	// 	delta := at - len(board)
	// 	for i := 0; i < delta; i++ {
	// 		fmt.Println("|.......|")
	// 	}
	// }

	// for i := len(board) - 1; i >= 0; i-- {
	// 	row := board[i]
	// 	fmt.Print("|")
	// 	for pixel := byte(0b10000000); pixel != 0b00000001; pixel = pixel >> 1 {
	// 		if pixel&row != empty {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println("|")
	// }

	fmt.Println("+-------+")
}
