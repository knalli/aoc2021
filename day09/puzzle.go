package day09

import (
	"fmt"
	"github.com/knalli/aoc"
	"sort"
)

func parseIntMatrix(lines []string) [][]int {
	result := make([][]int, 0)
	for _, line := range lines {
		result = append(result, aoc.ParseInts(line, ""))
	}
	return result
}

type LowPoint struct {
	X     int
	Y     int
	Value int
	Risk  int
}

func findLowPoints(heightmap [][]int) []LowPoint {
	lowPoints := make([]LowPoint, 0)
	for y, row := range heightmap {
		for x, v := range row {
			if y > 0 && heightmap[y-1][x] <= v {
				continue
			}
			if x < len(row)-1 && heightmap[y][x+1] <= v {
				continue
			}
			if y < len(heightmap)-1 && heightmap[y+1][x] <= v {
				continue
			}
			if x > 0 && heightmap[y][x-1] <= v {
				continue
			}
			lowPoints = append(lowPoints, LowPoint{
				X:     x,
				Y:     y,
				Value: v,
				Risk:  v + 1,
			})
		}
	}
	return lowPoints
}

func solve1(lines []string) error {
	heightmap := parseIntMatrix(lines)
	lowPoints := findLowPoints(heightmap)
	//fmt.Printf("Found %d points\n", len(lowPoints))
	sum := 0
	for _, p := range lowPoints {
		sum += p.Risk
	}
	aoc.PrintSolution(fmt.Sprintf("Sum of risks of all low points = %d", sum))
	return nil
}

func solve2(lines []string) error {
	heightmap := parseIntMatrix(lines)
	lowPoints := findLowPoints(heightmap)
	basins := make([]int, 0)
	visits := make(map[string]bool)
	for _, lowpoint := range lowPoints {
		basinSize := 0
		q := aoc.NewQueue()
		q.Add(lowpoint)
		for !q.IsEmpty() {
			p := q.Head().(LowPoint)
			if visited := visits[fmt.Sprintf("%d/%d", p.X, p.Y)]; visited {
				continue
			}
			visits[fmt.Sprintf("%d/%d", p.X, p.Y)] = true
			basinSize++
			//fmt.Printf("Start at %d/%d\n", p.X, p.Y)
			for _, y := range []int{p.Y - 1, p.Y, p.Y + 1} {
				for _, x := range []int{p.X - 1, p.X, p.X + 1} {
					// avoid out-of-bound
					if y < 0 || y == len(heightmap) || x < 0 || x == len(heightmap[0]) {
						continue
					}
					// avoid self
					if x == p.X && y == p.Y {
						continue
					}
					// avoid diagonal
					if x != p.X && y != p.Y {
						continue
					}
					if heightmap[y][x] != 9 && heightmap[y][x] > p.Value {
						//fmt.Printf("Found next level point %d/%d (%d)\n", x, y, nextV)
						q.Add(LowPoint{
							X:     x,
							Y:     y,
							Value: heightmap[y][x],
						})
					}
				}
			}
		}
		basins = append(basins, basinSize)
	}
	sort.Ints(basins)
	t := len(basins)
	aoc.PrintSolution(fmt.Sprintf("Product of three largest basins = %d", basins[t-3]*basins[t-2]*basins[t-1]))
	return nil
}
