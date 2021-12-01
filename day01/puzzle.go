package day01

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	prev := 0
	incs := 0
	for _, num := range aoc.ParseStringToIntArray(lines) {
		if prev > 0 && num > prev {
			incs++
		}
		prev = num
	}
	aoc.PrintSolution(fmt.Sprintf("There are %d that are larger than the previous measurement", incs))
	return nil
}

func solve2(lines []string) error {
	nums := aoc.ParseStringToIntArray(lines)
	prev := 0
	incs := 0
	for i := 0; i < len(nums)-2; i++ {
		sl := nums[i] + nums[i+1] + nums[i+2]
		if prev > 0 && sl > prev {
			incs++
		}
		prev = sl
	}
	aoc.PrintSolution(fmt.Sprintf("There are %d sums that are larger than the previous sum", incs))
	return nil
}
