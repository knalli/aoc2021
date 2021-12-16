package day15

import (
	"fmt"
	"github.com/knalli/aoc"
)

type Point struct {
	X int
	Y int
}

func solve1(lines []string) error {
	cave := Cave{}
	cave.Init(len(lines[0]), len(lines))

	for y, line := range lines {
		for x, c := range line {
			v := int(c) - 48
			cave.Set(x, y, v)
		}
	}

	maxX, maxY := cave.GetMax()
	cost := minCost(maxX, maxY, cave)
	aoc.PrintSolution(fmt.Sprintf("Total risk = %d", cost))

	return nil
}

func solve2(lines []string) error {
	cave := Cave{}
	cave.Init(len(lines[0]), len(lines))

	for y, line := range lines {
		for x, c := range line {
			v := int(c) - 48
			cave.Set(x, y, v)
		}
	}

	cave.Expand(5)

	maxX, maxY := cave.GetMax()
	cost := minCost(maxX, maxY, cave)
	aoc.PrintSolution(fmt.Sprintf("Total risk = %d", cost))

	return nil
}

func minCost(targetX int, targetY int, cave Cave) int {
	costs := make(map[Point]int)
	costs[Point{X: 0, Y: 0}] = 0

	q := aoc.NewQueue()
	q.Add(Point{X: 0, Y: 0})

	for !q.IsEmpty() {
		p := q.Head().(Point)
		cave.AdjacentsHV(p.X, p.Y, func(ax int, ay int) {
			ac := cave.Get(ax, ay)
			ap := Point{X: ax, Y: ay}
			if existingCost, exist := costs[ap]; exist {
				if existingCost <= costs[p]+ac {
					return // won't be better
				}
			}
			costs[ap] = costs[p] + ac
			q.Add(ap)
		})
	}

	return costs[Point{X: targetX, Y: targetY}]
}
