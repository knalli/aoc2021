package day02

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

func solve1(lines []string) error {
	h := 0
	d := 0
	if err := parseCommands(lines, func(cmd string, val int) {
		if cmd == "forward" {
			h += val
		} else if cmd == "down" {
			d += val
		} else if cmd == "up" {
			d -= val
		}
	}); err != nil {
		return err
	}
	aoc.PrintSolution(fmt.Sprintf("Submarine has a horizontal position of %d and a depth of %d.", h, d))
	aoc.PrintSolution(fmt.Sprintf("Check solution = %d", h*d))
	return nil
}

func solve2(lines []string) error {
	h := 0
	d := 0
	aim := 0
	if err := parseCommands(lines, func(cmd string, val int) {
		if cmd == "forward" {
			h += val
			d += aim * val
		} else if cmd == "down" {
			aim += val
		} else if cmd == "up" {
			aim -= val
		}
	}); err != nil {
		return err
	}
	aoc.PrintSolution(fmt.Sprintf("Submarine has a horizontal position of %d and a depth of %d.", h, d))
	aoc.PrintSolution(fmt.Sprintf("Check solution = %d", h*d))
	return nil
}

func parseCommands(lines []string, handler func(cmd string, val int)) error {
	for _, line := range lines {
		split := strings.Split(line, " ")
		v := aoc.ParseInt(split[1])
		handler(split[0], v)
	}
	return nil
}
