package day19

import (
	"fmt"
	"github.com/knalli/aoc"
)

type Point struct {
	X int
	Y int
	Z int
}

func (p *Point) ToString() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func (p *Point) Plus(o Point) Point {
	return Point{
		X: p.X + o.X,
		Y: p.Y + o.Y,
		Z: p.Z + o.Z,
	}
}

func (p *Point) Substrate(o Point) Point {
	return Point{
		X: p.X - o.X,
		Y: p.Y - o.Y,
		Z: p.Z - o.Z,
	}
}

func (p *Point) RotateX() {
	// cos(90) = 0
	// sin(90) = 1
	y := p.Y*0 - p.Z*1
	z := p.Y*1 + p.Z*0
	p.Y = y
	p.Z = z
}

func (p *Point) RotateY() {
	x := p.X*0 + p.Z*1
	z := p.Z*0 - p.X*1
	p.X = x
	p.Z = z
}

func (p *Point) RotateZ() {
	x := p.X*0 - p.Y*1
	y := p.X*1 + p.Y*0
	p.X = x
	p.Y = y
}

func (p *Point) Clone() Point {
	return Point{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
	}
}

func (p *Point) Rotated(rule string) Point {
	clone := p.Clone()
	for _, r := range rule {
		if r == 'x' {
			clone.RotateX()
		} else if r == 'y' {
			clone.RotateY()
		} else if r == 'z' {
			clone.RotateZ()
		}
	}
	return clone
}

func (p *Point) RotateBack(rule string) {
	for i := len(rule) - 1; i >= 0; i-- {
		r := rule[i]
		if r == 'x' {
			p.RotateX()
			p.RotateX()
			p.RotateX()
		} else if r == 'y' {
			p.RotateY()
			p.RotateY()
			p.RotateY()
		} else if r == 'z' {
			p.RotateZ()
			p.RotateZ()
			p.RotateZ()
		}
	}
}

func (p *Point) ManhattenDistance(o Point) int {
	return aoc.AbsInt(p.X-o.X) + aoc.AbsInt(p.Y-o.Y) + aoc.AbsInt(p.Z-o.Z)
}
