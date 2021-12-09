package main

import (
	_ "aoc2021/day01"
	_ "aoc2021/day02"
	_ "aoc2021/day03"
	"github.com/fatih/color"
	"time"

	_ "aoc2021/day04"
	_ "aoc2021/day05"
	_ "aoc2021/day06"
	_ "aoc2021/day07"
	_ "aoc2021/day09"
	//_ "aoc2021/dayXX"
	"errors"
	"github.com/knalli/aoc"
	"os"
	"strconv"
)

func init() {
	aoc.AocYear = 2021
}

func main() {
	start := time.Now()
	defer printTimeUsed(start)

	if err := invoke(os.Args); err != nil {
		aoc.PrintError(err)
		os.Exit(1)
	}
}
func invoke(args []string) error {
	if len(args) < 2 {
		return errors.New("first argument must be the day (e.g. 1)")
	}
	if day, err := strconv.Atoi(args[1]); err == nil {
		return aoc.Registry.Invoke(day, args[2:])
	} else {
		return err
	}
}

//noinspection GoUnhandledErrorResult
func printTimeUsed(start time.Time) {
	elapsed := time.Since(start)
	c := color.New(color.Bold, color.FgGreen)
	c.Printf("⏱ Took: %s\n", elapsed)
	c.Println()
}
