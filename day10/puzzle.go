package day10

import (
	"fmt"
	"github.com/knalli/aoc"
	"sort"
	"strings"
)

const OPENING_CHARS = "([{<"
const CLOSING_CHARS = ")]}>"

var ERROR_SCORES = map[int32]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var CORRECTION_SCORES = map[int32]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func parseLine(line string) (bool, int, string, int) {
	s := aoc.NewStack()
	// error?
	for _, c := range line {
		idx := strings.Index(OPENING_CHARS, string(c))
		if idx > -1 {
			s.Add(c)
			continue
		}
		idx = strings.Index(CLOSING_CHARS, string(c))
		if idx > -1 {
			last := s.Head().(int32)
			expected := int32(CLOSING_CHARS[strings.Index(OPENING_CHARS, string(last))])
			if c != expected {
				//fmt.Printf("Syntax error, expected %s, but found %s instead\n", string(expected), string(c))
				return false, ERROR_SCORES[c], "", 0
			}
		}
	}
	// correction
	result := line
	score := 0
	for !s.IsEmpty() {
		last := s.Head().(int32)
		expected := int32(CLOSING_CHARS[strings.Index(OPENING_CHARS, string(last))])
		result += string(expected)
		score *= 5
		score += CORRECTION_SCORES[expected]
	}
	return true, 0, result, score
}

func solve1(lines []string) error {
	totalErrorScore := 0
	for _, line := range lines {
		valid, score, _, _ := parseLine(line)
		if !valid {
			totalErrorScore += score
		}
	}
	aoc.PrintSolution(fmt.Sprintf("total syntax error score = %d", totalErrorScore))
	return nil
}

func solve2(lines []string) error {
	nums := make([]int, 0)
	for _, line := range lines {
		valid, _, _, score := parseLine(line)
		if !valid {
			continue
		}
		nums = append(nums, score)
	}
	sort.Ints(nums)
	aoc.PrintSolution(fmt.Sprintf("middle score (corrections) = %d", nums[len(nums)/2]))
	return nil
}
