package day20

import (
	"fmt"
	"github.com/knalli/aoc"
	"strconv"
)

const (
	DARK  int32 = '.'
	LIGHT int32 = '#'
)

func enhanceImage(input *Grid, enhancementAlgorithm string, outside int32) *Grid {
	output := NewGrid(input.Width()+2, input.Height()+2)
	for y := 0; y < output.Height(); y++ {
		for x := 0; x < output.Width(); x++ {
			inBasePoint := Point{X: x - 1, Y: y - 1}
			encoded := ""
			for _, inPoint := range inBasePoint.Block9() {
				if input.Valid(inPoint) {
					if input.GetValue(inPoint) == 1 {
						encoded += "1"
					} else {
						encoded += "0"
					}
				} else {
					// infinite (dark/light)
					// well, that trick required a look into reddit.
					// the variable "outside" is required because the "infinite default" flipped each round
					if outside == LIGHT {
						encoded += "1"
					} else {
						encoded += "0"
					}
				}
			}
			// binary
			decoded := binary2Int(encoded)
			if enhancementAlgorithm[decoded] == uint8(LIGHT) {
				output.SetValue(Point{x, y}, 1)
			}
		}
	}
	return output
}

func binary2Int(s string) int64 {
	if v, err := strconv.ParseInt(s, 2, 64); err != nil {
		return 0
	} else {
		return v
	}
}

func parseInput(lines []string) (string, *Grid) {
	enhancementAlgorithm := lines[0]

	grid := NewGrid(
		len(lines[2]),
		len(lines[2:]),
	)
	for y, line := range lines[2:] {
		for x, c := range line {
			if c == LIGHT {
				grid.SetValue(Point{x, y}, 1)
			}
		}
	}
	return enhancementAlgorithm, grid
}

func gridRenderer(x int, y int, v int) string {
	if v == 1 {
		return string(LIGHT)
	}
	return string(DARK)
}

func solve1(lines []string) error {
	enhancementAlgorithm, grid := parseInput(lines)

	output := grid
	for i := 0; i < 2; i++ {
		outside := LIGHT
		if i%2 == 0 {
			outside = DARK
		}
		output = enhanceImage(output, enhancementAlgorithm, outside)
	}

	//fmt.Println("Output:")
	//fmt.Println(output.ToString(gridRenderer))

	result := output.Count(func(x int, y int, v int) bool {
		return v == 1
	})
	aoc.PrintSolution(fmt.Sprintf("%d pixels are lit", result))

	return nil
}

func solve2(lines []string) error {
	enhancementAlgorithm, grid := parseInput(lines)

	output := grid
	for i := 0; i < 50; i++ {
		outside := LIGHT
		if i%2 == 0 {
			outside = DARK
		}
		output = enhanceImage(output, enhancementAlgorithm, outside)
	}

	//fmt.Println("Output:")
	//fmt.Println(output.ToString(gridRenderer))

	result := output.Count(func(x int, y int, v int) bool {
		return v == 1
	})
	aoc.PrintSolution(fmt.Sprintf("%d pixels are lit", result))

	return nil
}
