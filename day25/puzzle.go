package day25

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
)

const (
	EMPTY    int = '.'
	CC_EAST  int = '>'
	CC_SOUTH int = 'v'
)

type Grid struct {
	data [][]int
}

func (g *Grid) Each(f func(x int, y int, v int)) {
	for y := 0; y < len(g.data); y++ {
		for x := 0; x < len(g.data[0]); x++ {
			f(x, y, g.data[y][x])
		}
	}
}

func (g *Grid) Height() int {
	return len(g.data)
}

func (g *Grid) Width() int {
	return len(g.data[0])
}

func (g *Grid) Get(x int, y int) (int, error) {
	if 0 <= x && x < g.Width() && 0 <= y && y < g.Height() {
		return g.data[y][x], nil
	}
	return 0, errors.New("invalid x,y")
}

func (g *Grid) ToString() string {
	result := ""
	for y := range g.data {
		for _, c := range g.data[y] {
			result += string(c)
		}
		result += "\n"
	}
	return result
}

func (g *Grid) Set(x int, y int, v int) {
	g.data[y][x] = v
}

func NewGrid(height int, width int) *Grid {
	data := make([][]int, height)
	for i := range data {
		data[i] = make([]int, width)
	}
	return &Grid{
		data: data,
	}
}

func parseGrid(lines []string) *Grid {
	grid := NewGrid(len(lines), len(lines[0]))
	for y, line := range lines {
		for x, c := range line {
			grid.data[y][x] = int(c)
		}
	}
	return grid
}

type Point struct {
	x, y int
}

func solve1(lines []string) error {
	grid := parseGrid(lines)
	fmt.Printf("Initial:\n%s\n", grid.ToString())
	steps := 0
	moved := 0
	for {
		anyMoved := false

		// find movables
		points := make([][]Point, 0)
		grid.Each(func(x int, y int, v int) {
			if v == CC_EAST {
				ax := (x + 1) % grid.Width()
				ay := y
				if av, err := grid.Get(ax, ay); err == nil && av == EMPTY {
					points = append(points, []Point{{x, y}, {ax, ay}})
				}
			}
		})
		for _, p := range points {
			anyMoved = true
			moved++
			from := p[0]
			to := p[1]
			grid.Set(from.x, from.y, EMPTY)
			grid.Set(to.x, to.y, CC_EAST)
		}

		points = make([][]Point, 0)
		grid.Each(func(x int, y int, v int) {
			if v == CC_SOUTH {
				ax := x
				ay := (y + 1) % grid.Height()
				if av, err := grid.Get(ax, ay); err == nil && av == EMPTY {
					points = append(points, []Point{{x, y}, {ax, ay}})
				}
			}
		})
		for _, p := range points {
			anyMoved = true
			moved++
			from := p[0]
			to := p[1]
			grid.Set(from.x, from.y, EMPTY)
			grid.Set(to.x, to.y, CC_SOUTH)
		}

		if anyMoved {
			steps++
		} else {
			break
		}

		//fmt.Printf("After %d steps:\n%s\n", steps, grid.ToString())
	}
	fmt.Printf("After %d steps:\n%s\n", steps, grid.ToString())
	aoc.PrintSolution(fmt.Sprintf("Nothing moved after %d steps", steps+1))
	return nil
}

func solve2(lines []string) error {
	return nil
}
