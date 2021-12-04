package day04

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

func parseBoards(lines []string) []*BingoBoard {
	result := make([]*BingoBoard, 0)
	offset := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			result = append(result, parseBoard(lines[offset:i]))
			offset = i + 1
		}
	}
	// if input ends with last line (no newline)
	if lines[len(lines)-1] != "" {
		result = append(result, parseBoard(lines[offset:]))
	}
	return result
}

func parseBoard(lines []string) *BingoBoard {
	nums := make([][]int, 0)
	for _, line := range lines {
		nums = append(nums, ParseInts(line, " "))
	}
	return NewBingoBoard(nums)
}

func ParseInts(str string, delim string) []int {
	result := make([]int, 0)
	for _, s := range strings.Split(str, delim) {
		s2 := strings.TrimSpace(s)
		if len(s2) > 0 {
			result = append(result, aoc.ParseInt(s2))
		}
	}
	return result
}

func solve1(lines []string) error {
	numbers := aoc.ParseInts(lines[0], ",")
	boards := parseBoards(lines[2:])

	winnerFound := false
	for _, n := range numbers {
		// fmt.Printf("Check the number: %d\n", n)
		for _, board := range boards {
			board.Draw(n)
		}

		for b, board := range boards {
			//fmt.Printf("Board %d has a marked row: %t\n", b, board.HasMarkedRow())
			//fmt.Printf("Board %d has a marked col: %t\n", b, board.HasMarkedCol())
			if board.HasMarkedLine() {
				fmt.Printf("BINGO! Board #%d is the winner\n", b)
				unmarked := board.SumUnmarkedValues()
				aoc.PrintSolution(fmt.Sprintf("board's sum of unmarked = %d, num = %d, score = %d", unmarked, n, unmarked*n))
				winnerFound = true
			}
		}

		if winnerFound {
			break
		}
	}

	return nil
}

func solve2(lines []string) error {
	numbers := aoc.ParseInts(lines[0], ",")
	boards := parseBoards(lines[2:])

	boardsWinner := make(map[int]bool)
	for _, n := range numbers {
		for _, board := range boards {
			board.Draw(n)
		}

		for b, board := range boards {
			if won := boardsWinner[b]; won {
				// skip boards already won
				continue
			}
			if board.HasMarkedLine() {
				boardsWinner[b] = true
				if len(boardsWinner) == len(boards) {
					fmt.Printf("BINGO! Board #%d is the last winner\n", b)
					unmarked := board.SumUnmarkedValues()
					aoc.PrintSolution(fmt.Sprintf("board's sum of unmarked = %d, num = %d, score = %d", unmarked, n, unmarked*n))
				}
			}
		}
	}

	return nil
}
