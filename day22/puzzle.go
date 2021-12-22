package day22

import (
	"fmt"
	"github.com/knalli/aoc"
	"regexp"
)

type Point struct {
	X int
	Y int
	Z int
}

func (p *Point) ToString() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

// Range […[
type Range struct {
	Begin int
	End   int
}

func (r *Range) Each(f func(v int)) {
	for v := r.Begin; v < r.End; v++ {
		f(v)
	}
}

func (r *Range) IncludesRange(other Range) bool {
	return r.Begin <= other.Begin && r.End >= other.End
}
func (r *Range) Length() int {
	return aoc.AbsInt(r.End - r.Begin)
}

func (r *Range) ToString() string {
	return fmt.Sprintf("%d..%d", r.Begin, r.End-1)
}

func (r *Range) Intersection(other Range) Range {
	if r.End <= other.Begin || r.Begin >= other.End {
		return Range{Begin: 0, End: 0}
	} else {
		return Range{
			Begin: aoc.MaxInt(r.Begin, other.Begin),
			End:   aoc.MinInt(r.End, other.End),
		}
	}
}

type Instruction struct {
	State  bool
	RangeX Range
	RangeY Range
	RangeZ Range
}

type Cube struct {
	Points map[Point]bool
	rangeX Range
	rangeY Range
	rangeZ Range
}

type Cuboid struct {
	RangeX Range
	RangeY Range
	RangeZ Range
}

func (c *Cuboid) ToString() string {
	return fmt.Sprintf("{x=%s,y=%s,z=%s}", c.RangeX.ToString(), c.RangeY.ToString(), c.RangeZ.ToString())
}

func (c *Cuboid) Includes(other *Cuboid) bool {
	return c.RangeX.IncludesRange(other.RangeX) && c.RangeY.IncludesRange(other.RangeY) && c.RangeZ.IncludesRange(other.RangeZ)
}

func (c *Cuboid) Volume() int {
	return c.RangeX.Length() * c.RangeY.Length() * c.RangeZ.Length()
}

type Intersection struct {
	Includes bool
	Cuboid   *Cuboid
}

func (c *Cuboid) Intersection(other *Cuboid) *Cuboid {
	return &Cuboid{
		RangeX: c.RangeX.Intersection(other.RangeX),
		RangeY: c.RangeY.Intersection(other.RangeY),
		RangeZ: c.RangeZ.Intersection(other.RangeZ),
	}
}

func (c *Cube) Set(p Point, newValue bool) {
	if oldValue, exist := c.Points[p]; !exist || oldValue != newValue {
		c.Points[p] = newValue

		if !exist {
			c.rangeX.Begin = aoc.MinInt(c.rangeX.Begin, p.X)
			c.rangeX.End = aoc.MaxInt(c.rangeX.End, p.X)
			c.rangeY.Begin = aoc.MinInt(c.rangeY.Begin, p.Y)
			c.rangeY.End = aoc.MaxInt(c.rangeY.End, p.Y)
			c.rangeZ.Begin = aoc.MinInt(c.rangeZ.Begin, p.Z)
			c.rangeZ.End = aoc.MaxInt(c.rangeZ.End, p.Z)
		}
	}
}

func (c *Cube) CountOn() int {
	result := 0
	for _, value := range c.Points {
		if value {
			result++
		}
	}
	return result
}

func NewCube() *Cube {
	return &Cube{
		Points: make(map[Point]bool),
	}
}

func parseLine(line string) Instruction {

	// on x=10..12,y=10..12,z=10..12
	pattern, _ := regexp.Compile("(on|off) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)")
	m := pattern.FindStringSubmatch(line)
	return Instruction{
		State: m[1] == "on",
		RangeX: Range{
			Begin: aoc.ParseInt(m[2]),
			End:   aoc.ParseInt(m[3]) + 1,
		},
		RangeY: Range{
			Begin: aoc.ParseInt(m[4]),
			End:   aoc.ParseInt(m[5]) + 1,
		},
		RangeZ: Range{
			Begin: aoc.ParseInt(m[6]),
			End:   aoc.ParseInt(m[7]) + 1,
		},
	}
}

func solveViaIntersection(steps []*Instruction) int {

	// helper
	type CubeVolume struct {
		Cube   Cuboid
		Weight int
	}

	// we only hold cubes with state=on
	cubes := make([]*CubeVolume, 0)
	for _, step := range steps {
		nextCube := Cuboid{
			RangeX: step.RangeX,
			RangeY: step.RangeY,
			RangeZ: step.RangeZ,
		}

		// reduce always the existing cubes for this new cuboid (even if state=on)
		{
			toCreate := make(map[string]*CubeVolume)
			toRemove := make([]int, 0)
			for cvIdx, cv := range cubes {
				intersection := nextCube.Intersection(&cv.Cube)
				intersectionVolume := intersection.Volume()
				if intersection.ToString() == cv.Cube.ToString() {
					toRemove = append(toRemove, cvIdx)
				} else if intersectionVolume > 0 {
					id := cv.Cube.ToString()
					if val, exist := toCreate[id]; exist {
						val.Weight -= cv.Weight
					} else {
						toCreate[id] = &CubeVolume{
							Cube:   *intersection,
							Weight: -cv.Weight,
						}
					}
				}
			}

			if len(toRemove) > 0 {
				removed := make([]*CubeVolume, 0)
				for cvIdx, cv := range cubes {
					contains := false
					for _, idx := range toRemove {
						if cvIdx == idx {
							contains = true
							break
						}
					}
					if !contains {
						removed = append(removed, cv)
					}
				}
				cubes = removed
			}
			for _, cv := range toCreate {
				if cv.Weight != 0 {
					cubes = append(cubes, cv)
				}
			}
		}

		if step.State {
			cubes = append(cubes, &CubeVolume{
				Cube:   nextCube,
				Weight: 1,
			})
		}
	}

	result := 0
	for _, cv := range cubes {
		result += cv.Weight * cv.Cube.Volume()
	}

	return result
}

func solve1(lines []string) error {
	steps := make([]*Instruction, 0)
	for _, line := range lines {
		step := parseLine(line)
		steps = append(steps, &step)
	}

	// part1 limit to -50..50
	newSteps := make([]*Instruction, 0)
	for _, step := range steps {
		skip := false
		for _, r := range []*Range{&step.RangeX, &step.RangeY, &step.RangeZ} {
			if r.Begin < -50 {
				r.Begin = 50
			}
			if r.End > 50 {
				r.End = 51
			}
			if r.Begin >= r.End {
				skip = true
				break
			}
		}
		if !skip {
			newSteps = append(newSteps, step)
		}
	}
	steps = newSteps

	cube := NewCube()
	for _, step := range steps {
		step.RangeX.Each(func(x int) {
			step.RangeY.Each(func(y int) {
				step.RangeZ.Each(func(z int) {
					p := Point{X: x, Y: y, Z: z}
					//fmt.Printf("%s to %c\n", p.ToString(), step.State)
					cube.Set(p, step.State)
				})
			})
		})
	}
	aoc.PrintSolution(fmt.Sprintf("-50..50; Cubes on = %d", cube.CountOn()))
	// of part2, verify
	aoc.PrintSolution(fmt.Sprintf("-50..50; Cubes on = %d", solveViaIntersection(steps)))
	return nil
}

func solve2(lines []string) error {
	steps := make([]*Instruction, 0)
	for _, line := range lines {
		step := parseLine(line)
		steps = append(steps, &step)
	}

	aoc.PrintSolution(fmt.Sprintf("Cubes on = %d", solveViaIntersection(steps)))
	return nil
}
