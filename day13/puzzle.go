package day13

import (
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Grid struct {
	//points []Point
	refs map[Point]bool
	minX int
	maxX int
	minY int
	maxY int
}

func NewGrid() *Grid {
	return &Grid{
		minX: math.MaxInt64,
		maxX: math.MinInt64,
		minY: math.MaxInt64,
		maxY: math.MinInt64,
	}
}

func (g *Grid) pointId(p Point) string {
	return fmt.Sprintf("%d/%d", p.X, p.Y)
}

func (g *Grid) Add(p Point) {
	if g.refs == nil {
		g.refs = make(map[Point]bool)
	}
	if _, exist := g.refs[p]; !exist {
		//g.points = append(g.points, p)
		g.refs[p] = true
		g.minX = aoc.MinInt(g.minX, p.X)
		g.maxX = aoc.MaxInt(g.maxX, p.X)
		g.minY = aoc.MinInt(g.minY, p.Y)
		g.maxY = aoc.MaxInt(g.maxY, p.Y)
	}
}

func (g *Grid) ToString() string {
	res := ""
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if _, exist := g.refs[Point{X: x, Y: y}]; exist {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func (g *Grid) FoldUp(atY int) {
	for y := atY + 1; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			p := Point{X: x, Y: y}
			if _, exist := g.refs[p]; exist {
				delete(g.refs, p)
				newP := Point{
					X: x,
					Y: atY - (y - atY),
				}
				g.Add(newP)
			}
		}
	}
	g.maxY = atY - 1
}

func (g *Grid) FoldLeft(atX int) {
	for y := g.minY; y <= g.maxY; y++ {
		for x := atX + 1; x <= g.maxX; x++ {
			p := Point{X: x, Y: y}
			if _, exist := g.refs[p]; exist {
				delete(g.refs, p)
				newP := Point{
					X: atX - (x - atX),
					Y: y,
				}
				g.Add(newP)
			}
		}
	}
	g.maxX = atX - 1
}

type FoldRule struct {
	At     string
	Amount int
}

func solve1(lines []string) error {
	grid := NewGrid()
	rules := make([]FoldRule, 0)
	for _, line := range lines {
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			grid.Add(Point{
				X: aoc.ParseInt(parts[0]),
				Y: aoc.ParseInt(parts[1]),
			})
		} else if strings.Contains(line, "fold along x=") {
			rules = append(rules, FoldRule{
				At:     "x",
				Amount: aoc.ParseInt(line[13:]),
			})
		} else if strings.Contains(line, "fold along y=") {
			rules = append(rules, FoldRule{
				At:     "y",
				Amount: aoc.ParseInt(line[13:]),
			})
		}
	}
	//fmt.Println(grid.ToString())
	//fmt.Println("============")
	for _, rule := range rules {
		if rule.At == "x" {
			grid.FoldLeft(rule.Amount)
		} else if rule.At == "y" {
			grid.FoldUp(rule.Amount)
		}
		break
	}
	//fmt.Println(grid.ToString())
	count := 0
	for _, c := range grid.ToString() {
		if c == '#' {
			count++
		}
	}
	aoc.PrintSolution(fmt.Sprintf("There are %d dots after the first fold instruction", count))
	return nil
}

func solve2(lines []string) error {
	grid := NewGrid()
	rules := make([]FoldRule, 0)
	for _, line := range lines {
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			grid.Add(Point{
				X: aoc.ParseInt(parts[0]),
				Y: aoc.ParseInt(parts[1]),
			})
		} else if strings.Contains(line, "fold along x=") {
			rules = append(rules, FoldRule{
				At:     "x",
				Amount: aoc.ParseInt(line[13:]),
			})
		} else if strings.Contains(line, "fold along y=") {
			rules = append(rules, FoldRule{
				At:     "y",
				Amount: aoc.ParseInt(line[13:]),
			})
		}
	}
	for _, rule := range rules {
		fmt.Printf("Applying fold..\n")
		if rule.At == "x" {
			grid.FoldLeft(rule.Amount)
		} else if rule.At == "y" {
			grid.FoldUp(rule.Amount)
		}
	}
	fmt.Println(grid.ToString())
	return nil
}
