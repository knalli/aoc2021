package day19

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

func calcUniqueRotationRules() []string {
	p := Point{5, 11, 19}
	m := make(map[Point]int)
	rotations := make([]string, 0)
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				c := p.Clone()
				rule := ""
				for ix := 0; ix < x; ix++ {
					c.RotateX()
					rule += "x"
				}
				for iy := 0; iy < y; iy++ {
					c.RotateY()
					rule += "y"
				}
				for iz := 0; iz < z; iz++ {
					c.RotateZ()
					rule += "z"
				}
				if _, exist := m[c]; exist {
					continue
				}
				m[c]++
				rotations = append(rotations, rule)
			}
		}
	}
	return rotations
}

func ensureSamePositionInDifferentOrientations() {
	o := Scanner{
		Position: Point{},
		beacons: []*Point{
			&Point{-1, -1, 1},
			&Point{-2, -2, 2},
			&Point{-3, -3, 3},
			&Point{-2, -3, 1},
			&Point{5, 6, -4},
			&Point{8, 0, 7},
		},
		orientation: Point{},
	}
	s := o
	fmt.Println(s.ToString() + "\n=====")
	s = *o.Clone()
	s = *s.GetRotated("xyy")
	fmt.Println(s.ToString() + "\n=====")
	s = *o.Clone()
	s = *s.GetRotated("yyy")
	fmt.Println(s.ToString() + "\n=====")
	s = *o.Clone()
	s = *s.GetRotated("yxx")
	fmt.Println(s.ToString() + "\n=====")
	s = *o.Clone()
	s = *s.GetRotated("zxxx")
	fmt.Println(s.ToString() + "\n=====")
}

func computePointDeltaMap(a []Point, b []Point) map[Point]int {
	result := make(map[Point]int)
	for _, bp := range b {
		for _, ap := range a {
			dp := bp.Substrate(ap)
			result[dp]++
		}
	}
	return result
}

func solve1(lines []string) error {

	uniqueRotations := calcUniqueRotationRules()
	scanners := parseScanners(lines)

	scannerPositions := make(map[int]Point)
	// first scanner is fixed
	scannerPositions[0] = scanners[0].Position

	for len(scannerPositions) < len(scanners) {
		restart := false
		for n, nScanner := range scanners {
			if _, exist := scannerPositions[n]; exist {
				// already positioned
				continue
			}
			for o, oScanner := range scanners {
				if _, exist := scannerPositions[o]; !exist {
					// not yet positioned
					continue
				}

				oBeacons := oScanner.AbsoluteBeacons()
				for _, rule := range uniqueRotations {
					nBeacons := nScanner.Clone().GetRotated(rule).AbsoluteBeacons()
					deltaMap := computePointDeltaMap(nBeacons, oBeacons)
					for deltaPoint, value := range deltaMap {
						if value >= 12 {
							// found
							fmt.Printf("Overlapping found for scanner %d (via %d, rotated=%s)\n", n, o, rule)
							nScanner.Position = deltaPoint
							scannerPositions[n] = deltaPoint
							nScanner.RotateBack(rule)
							oScanner.AddBeacons(nBeacons, deltaPoint)
							restart = true
							break
						}
					}
					if restart {
						break
					}
				}
				if restart {
					break
				}
			}
			if restart {
				break
			}
		}
	}

	beacons := make(map[Point]int)
	for _, b := range scanners[0].AbsoluteBeacons() {
		beacons[b]++
	}
	for i, scanner := range scanners {
		fmt.Printf("Scanner %d, position %s\n", i, scanner.Position.ToString())
	}
	aoc.PrintSolution(fmt.Sprintf("Total beacons = %d", len(beacons)))

	// part2
	max := 0
	for a, aPos := range scannerPositions {
		for b, bPos := range scannerPositions {
			if a == b {
				continue
			}
			max = aoc.MaxInt(max, aPos.ManhattenDistance(bPos))
		}
	}
	aoc.PrintSolution(fmt.Sprintf("Largest manhatten distance = %d", max))

	return nil
}

func solve2(lines []string) error {
	aoc.PrintSolution("see part1")
	return nil
}

func parseScanners(lines []string) []*Scanner {
	scanners := make([]*Scanner, 0)
	input := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "---") {
			if len(input) > 0 {
				scanners = append(scanners, parseScanner(input))
			}
			input = make([]string, 0)
		} else {
			input = append(input, line)
		}
	}
	if len(input) > 0 {
		scanners = append(scanners, parseScanner(input))
	}
	return scanners
}
