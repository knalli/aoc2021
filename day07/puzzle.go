package day07

import (
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"sort"
)

func computeDistances(nums []int, pos int) int {
	result := 0
	for _, n := range nums {
		if pos > n {
			result += pos - n
		} else {
			result += n - pos
		}
	}
	return result
}

func computeDistancesGauss(nums []int, pos int) int {
	result := 0
	for _, n := range nums {
		d := 0
		if pos > n {
			d = pos - n
		} else {
			d = n - pos
		}
		f := ((d * d) + d) / 2
		result += f
	}
	return result
}

func solve1(lines []string) error {
	nums := aoc.ParseInts(lines[0], ",")
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	mid := (max - min) / 2

	minFuel := math.MaxInt64
	minPos := -1
	// initial idea was optimizing starting from the mid, but not followed
	for i := 0; i < mid; i++ {
		left := min + mid - i
		right := min + mid + i

		fuel := computeDistances(nums, left)
		if fuel < minFuel {
			minFuel = fuel
			minPos = left
		}

		if right != left {
			fuel := computeDistances(nums, right)
			if fuel < minFuel {
				minFuel = fuel
				minPos = right
			}
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Crabs will use %d fuel to get to position %d!", minFuel, minPos))
	return nil
}

func solve2(lines []string) error {
	nums := aoc.ParseInts(lines[0], ",")
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	mid := (max - min) / 2

	minFuel := math.MaxInt64
	minPos := -1
	// initial idea was optimizing starting from the mid, but not followed
	for i := 0; i < mid; i++ {
		left := min + mid - i
		right := min + mid + i

		fuel := computeDistancesGauss(nums, left)
		if fuel < minFuel {
			minFuel = fuel
			minPos = left
		}

		if right != left {
			fuel := computeDistancesGauss(nums, right)
			if fuel < minFuel {
				minFuel = fuel
				minPos = right
			}
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Crabs will use %d fuel to get to position %d!", minFuel, minPos))
	return nil
}
