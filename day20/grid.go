package day20

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) Block9() []Point {
	x := p.X
	y := p.Y
	return []Point{
		{X: x - 1, Y: y - 1},
		{X: x, Y: y - 1},
		{X: x + 1, Y: y - 1},
		{X: x - 1, Y: y},
		{X: x, Y: y},
		{X: x + 1, Y: y},
		{X: x - 1, Y: y + 1},
		{X: x, Y: y + 1},
		{X: x + 1, Y: y + 1},
	}
}

type Grid struct {
	data [][]int
}

func NewGrid(width int, height int) *Grid {
	data := make([][]int, height)
	for i := 0; i < height; i++ {
		data[i] = make([]int, width)
	}
	return &Grid{
		data: data,
	}
}

func (g *Grid) Height() int {
	return len(g.data)
}

func (g *Grid) Width() int {
	return len(g.data[0])
}

func (g *Grid) Valid(p Point) bool {
	x := p.X
	y := p.Y
	return -1 < y && y < len(g.data) && -1 < x && x < len(g.data[0])
}

func (g *Grid) SetValue(p Point, v int) {
	if !g.Valid(p) {
		panic("invalid p")
	}
	g.data[p.Y][p.X] = v
}

func (g *Grid) GetValue(p Point) int {
	if !g.Valid(p) {
		panic("invalid p")
	}
	return g.data[p.Y][p.X]
}

func (g *Grid) ToString(f func(x int, y int, v int) string) string {
	result := ""
	for y := 0; y < len(g.data); y++ {
		for x := 0; x < len(g.data[y]); x++ {
			result += fmt.Sprintf("%s", f(x, y, g.data[y][x]))
		}
		result += "\n"
	}
	return result
}

func (g *Grid) Count(f func(x int, y int, v int) bool) int {
	result := 0
	for y := 0; y < len(g.data); y++ {
		for x := 0; x < len(g.data[y]); x++ {
			if f(x, y, g.data[y][x]) {
				result++
			}
		}
	}
	return result
}
