package day17

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

type Area struct {
	X Interval
	Y Interval
}

type Interval struct {
	Start int
	End   int
}

func (i *Interval) Includes(v int) bool {
	return i.Start <= v && v <= i.End
}

type Point struct {
	X int
	Y int
}

func parseInterval(str string) *Interval {
	parts := strings.Split(str, "..")
	return &Interval{
		Start: aoc.ParseInt(parts[0]),
		End:   aoc.ParseInt(parts[1]),
	}
}

func parseTargetArea(lines []string) Area {
	targetArea := Area{
		X: *parseInterval(lines[0][strings.Index(lines[0], "x=")+2 : strings.Index(lines[0], ", y=")]),
		Y: *parseInterval(lines[0][strings.Index(lines[0], "y=")+2:]),
	}
	return targetArea
}

func solve1(lines []string) error {
	targetArea := parseTargetArea(lines)
	// bad day: easy for gauss sum of y; but need several hours to find out (reddit) the off by one...
	aoc.PrintSolution(fmt.Sprintf("max y = %d", gaussSum(-targetArea.Y.Start-1)))
	return nil
}

func solve2(lines []string) error {
	targetArea := parseTargetArea(lines)

	count := 0
	for vx := 1; vx <= targetArea.X.End; vx++ {
		for vy := targetArea.Y.Start; vy <= -targetArea.Y.Start; vy++ {
			v := Point{X: vx, Y: vy}
			p := Point{X: 0, Y: 0}
			for p.Y >= targetArea.Y.Start {
				p.X += v.X
				p.Y += v.Y
				if v.X > 0 {
					v.X--
				} else if v.X < 0 {
					v.X++
				}
				v.Y--
				if targetArea.X.Includes(p.X) && targetArea.Y.Includes(p.Y) {
					//fmt.Printf("%d/%d\n", p.X, p.Y)
					count++
					break
				}
			}
		}
	}
	aoc.PrintSolution(fmt.Sprintf("total = %d", count))
	return nil
}

func gaussSum(x int) int {
	return ((x * x) + x) / 2
}
