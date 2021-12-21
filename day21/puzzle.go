package day21

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

// part1

func solve1(lines []string) error {
	positions := make([]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "position: ")
		positions = append(positions, aoc.ParseInt(parts[1]))
	}
	totals, turns := play(positions, deterministicDiceGenerator(), func(score int) bool {
		return score >= 1000
	})
	loosingPlayerScore := aoc.MinInt(totals[0], totals[1])
	aoc.PrintSolution(fmt.Sprintf("turns=%d, rolls=%d, loosingScore=%d, result=%d", turns, turns*3, loosingPlayerScore, loosingPlayerScore*turns*3))
	return nil
}

func play(positions []int, generator func() int, isFinished func(score int) bool) ([]int, int) {
	totals := make([]int, len(positions))
	for i := range positions {
		totals[i] = 0
		positions[i]-- // change positions as index (-1)
	}
	turns := 0
	for {
		for p := range positions {
			turns++
			move := generator() + generator() + generator()
			positions[p] = (positions[p] + move) % 10
			totals[p] += positions[p] + 1
			if isFinished(totals[p]) {
				return totals, turns
			}
		}
	}
}

func deterministicDiceGenerator() func() int {
	next := 1
	return func() int {
		current := next
		next++
		if next > 100 {
			next = 1
		}
		return current
	}
}

// part2

func solve2(lines []string) error {
	positions := make([]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "position: ")
		positions = append(positions, aoc.ParseInt(parts[1]))
	}

	for i := range positions {
		positions[i]-- // change positions as index (-1)
	}

	// 3x3 = 27 possibilities, however, this can be reduced
	rolls := make(map[int]int)
	for _, i := range []int{1, 2, 3} {
		for _, j := range []int{1, 2, 3} {
			for _, k := range []int{1, 2, 3} {
				rolls[i+j+k]++
			}
		}
	}

	results := play2(rolls, positions, 0, []int{0, 0}, func(score int) bool {
		return score >= 21
	})
	aoc.PrintSolution(fmt.Sprintf("player 1 = %d, player 2 = %d", results[0], results[1]))
	aoc.PrintSolution(fmt.Sprintf("max = %d", aoc.MaxInt(results[0], results[1])))

	return nil
}

// thank you @ https://github.com/kolonialno/adventofcode/commit/dd6a632b5e6781f9325af067e1e851f6485b79b2#diff-bd0ebfd104228323882b0f7a44db2c0f2ab2e590f634abae953e7aa04d2d57ebR38
func play2(rolls map[int]int, positions []int, p int, totals []int, isFinished func(score int) bool) []int {
	result := make([]int, len(positions))
	for move, count := range rolls {
		newPositions := make([]int, len(positions))
		for i := range positions {
			newPositions[i] = positions[i]
		}
		newTotals := make([]int, len(totals))
		for i := range totals {
			newTotals[i] = totals[i]
		}
		newPositions[p] = (positions[p] + move) % 10
		newTotals[p] += newPositions[p] + 1
		if isFinished(newTotals[p]) {
			result[p] += count
		} else {
			for p2, result2 := range play2(rolls, newPositions, (p+1)%len(positions), newTotals, isFinished) {
				result[p2] += result2 * count
			}
		}
	}
	return result
}
