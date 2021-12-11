package day11

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

type Cave struct {
	Octopuses [][]int
}

func (c *Cave) Each(f func(dx int, dy int)) {
	for y := 0; y < len(c.Octopuses); y++ {
		for x := 0; x < len(c.Octopuses[y]); x++ {
			f(x, y)
		}
	}
}

func (c *Cave) Increment(x int, y int) {
	c.Octopuses[y][x] = c.Octopuses[y][x] + 1
}

func (c *Cave) Reset(x int, y int) {
	c.Octopuses[y][x] = 0
}

func (c *Cave) ToString() string {
	result := ""
	for y := 0; y < len(c.Octopuses); y++ {
		for x := 0; x < len(c.Octopuses[y]); x++ {
			result += fmt.Sprintf("%d", c.Octopuses[y][x])
		}
		result += "\n"
	}
	return result
}

func (c *Cave) FindFirstWithFilter(filter func(x int, y int) bool, f func(x int, y int)) {
	for y := 0; y < len(c.Octopuses); y++ {
		for x := 0; x < len(c.Octopuses[y]); x++ {
			if filter(x, y) {
				f(x, y)
				return
			}
		}
	}
}

func (c *Cave) Adjacents(dx int, dy int, f func(x int, y int)) {
	for _, y := range []int{dy - 1, dy, dy + 1} {
		for _, x := range []int{dx - 1, dx, dx + 1} {
			if y == dy && x == dx {
				continue
			}
			if y < 0 || y > len(c.Octopuses)-1 {
				continue
			}
			if x < 0 || x > len(c.Octopuses[0])-1 {
				continue
			}
			f(x, y)
		}
	}
}

func solve1(lines []string) error {
	cave := Cave{
		Octopuses: make([][]int, 0),
	}
	for _, line := range lines {
		cave.Octopuses = append(cave.Octopuses, aoc.ParseInts(line, ""))
	}
	//fmt.Printf("Before any steps: \n%s\n", cave.ToString())
	steps := 0
	flashes := 0
	for steps < 100 {
		steps++
		cave.Each(func(x int, y int) {
			cave.Increment(x, y)
		})
		flashing := make(map[string]bool)
		moreThan9Filter := func(x int, y int) bool {
			if flashed := flashing[fmt.Sprintf("%d/%d", x, y)]; flashed {
				return false
			}
			return cave.Octopuses[y][x] > 9
		}
		for {
			found := false
			cave.FindFirstWithFilter(moreThan9Filter, func(x int, y int) {
				flashing[fmt.Sprintf("%d/%d", x, y)] = true
				found = true
				//fmt.Printf("adjacents from %d/%d\n", x, y)
				cave.Adjacents(x, y, func(ax int, ay int) {
					//fmt.Printf("adjacents from %d/%d -> %d/%d\n", x, y, ax, ay)
					if flashed := flashing[fmt.Sprintf("%d/%d", ax, ay)]; flashed {
						return
					}
					cave.Increment(ax, ay)
				})
			})
			if !found {
				break
			}
		}
		for key := range flashing {
			part := strings.Split(key, "/")
			cave.Reset(aoc.ParseInt(part[0]), aoc.ParseInt(part[1]))
		}
		flashes += len(flashing)
		//fmt.Printf("After step %d: \n%s\n", steps, cave.ToString())
	}
	aoc.PrintSolution(fmt.Sprintf("After %d steps, there habve been a total of %d flashes", steps, flashes))
	return nil
}

func solve2(lines []string) error {
	cave := Cave{
		Octopuses: make([][]int, 0),
	}
	for _, line := range lines {
		cave.Octopuses = append(cave.Octopuses, aoc.ParseInts(line, ""))
	}
	//fmt.Printf("Before any steps: \n%s\n", cave.ToString())
	steps := 0
	flashes := 0
	totalOctopuses := len(cave.Octopuses) * len(cave.Octopuses[0])
	for {
		steps++
		cave.Each(func(x int, y int) {
			cave.Increment(x, y)
		})
		flashing := make(map[string]bool)
		moreThan9Filter := func(x int, y int) bool {
			if flashed := flashing[fmt.Sprintf("%d/%d", x, y)]; flashed {
				return false
			}
			return cave.Octopuses[y][x] > 9
		}
		for {
			found := false
			cave.FindFirstWithFilter(moreThan9Filter, func(x int, y int) {
				flashing[fmt.Sprintf("%d/%d", x, y)] = true
				found = true
				//fmt.Printf("adjacents from %d/%d\n", x, y)
				cave.Adjacents(x, y, func(ax int, ay int) {
					//fmt.Printf("adjacents from %d/%d -> %d/%d\n", x, y, ax, ay)
					if flashed := flashing[fmt.Sprintf("%d/%d", ax, ay)]; flashed {
						return
					}
					cave.Increment(ax, ay)
				})
			})
			if !found {
				break
			}
		}
		for key := range flashing {
			part := strings.Split(key, "/")
			cave.Reset(aoc.ParseInt(part[0]), aoc.ParseInt(part[1]))
		}
		flashes += len(flashing)
		//fmt.Printf("After step %d: \n%s\n", steps, cave.ToString())
		if len(flashing) == totalOctopuses {
			break
		}
	}
	aoc.PrintSolution(fmt.Sprintf("After %d steps, there have been a total of %d flashes (first time simultaneously)", steps, flashes))
	return nil
}
