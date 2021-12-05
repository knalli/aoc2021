package day05

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type Vent struct {
	Begin Point
	End   Point
}

func NewVent(begin Point, end Point) *Vent {
	return &Vent{
		Begin: begin,
		End:   end,
	}
}

func parseVents(lines []string) []Vent {
	result := make([]Vent, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")
		t := strings.Split(split[0], ",")
		begin := *NewPoint(
			aoc.ParseInt(t[0]),
			aoc.ParseInt(t[1]),
		)
		t = strings.Split(split[2], ",")
		end := *NewPoint(
			aoc.ParseInt(t[0]),
			aoc.ParseInt(t[1]),
		)
		result = append(result, *NewVent(begin, end))
	}
	return result
}

// solution for horizontal and vertical lines
func solve1(lines []string) error {
	mem := make(map[Point]int, 0)
	for _, vent := range parseVents(lines) {
		if vent.Begin.X == vent.End.X {
			x := vent.Begin.X
			y1 := aoc.MinInt(vent.Begin.Y, vent.End.Y)
			y2 := aoc.MaxInt(vent.Begin.Y, vent.End.Y)
			for y := y1; y <= y2; y++ {
				p := *NewPoint(x, y)
				mem[p]++
			}
		} else if vent.Begin.Y == vent.End.Y {
			y := vent.Begin.Y
			x1 := aoc.MinInt(vent.Begin.X, vent.End.X)
			x2 := aoc.MaxInt(vent.Begin.X, vent.End.X)
			for x := x1; x <= x2; x++ {
				p := *NewPoint(x, y)
				mem[p]++
			}
		} else {
			// ignored
		}
	}
	// count at least two overlaps
	count := 0
	for _, c := range mem {
		if c > 1 {
			count++
		}
	}
	aoc.PrintSolution(fmt.Sprintf("At %d points having at least two lines overlap", count))
	return nil
}

// solution for horizontal, vertical and diagonal (45°) lines
func solve2(lines []string) error {
	mem := make(map[Point]int, 0)
	for _, vent := range parseVents(lines) {
		x1 := vent.Begin.X
		x2 := vent.End.X
		xd := 0
		if x1 < x2 {
			xd = 1
		} else if x1 > x2 {
			xd = -1
		}
		y1 := vent.Begin.Y
		y2 := vent.End.Y
		yd := 0
		if y1 < y2 {
			yd = 1
		} else if y1 > y2 {
			yd = -1
		}
		end := *NewPoint(x2, y2)
		p := *NewPoint(x1, y1)
		for p != end {
			mem[p]++
			p = *NewPoint(p.X+xd, p.Y+yd)
		}
		mem[p]++ // end also
	}
	// count at least two overlaps
	count := 0
	for _, c := range mem {
		if c > 1 {
			count++
		}
	}
	aoc.PrintSolution(fmt.Sprintf("At %d points having at least two lines overlap", count))
	return nil
}
