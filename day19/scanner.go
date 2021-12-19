package day19

import (
	"github.com/knalli/aoc"
	"strings"
)

type Scanner struct {
	Position    Point
	beacons     []*Point
	orientation Point
}

func (s *Scanner) RotateBack(rule string) {
	for _, b := range s.beacons {
		b.RotateBack(rule)
	}
}

func parseScanner(lines []string) *Scanner {
	beacons := make([]*Point, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		beacons = append(beacons, &Point{
			X: aoc.ParseInt(parts[0]),
			Y: aoc.ParseInt(parts[1]),
			Z: aoc.ParseInt(parts[2]),
		})
	}
	return &Scanner{
		Position:    Point{X: 0, Y: 0, Z: 0},
		beacons:     beacons,
		orientation: Point{X: 0, Y: 0, Z: 0},
	}
}

func (s *Scanner) AbsoluteBeacons() []Point {
	result := make([]Point, len(s.beacons))
	for i, b := range s.beacons {
		result[i] = b.Plus(s.Position)
	}
	return result
}

func (s *Scanner) RelativeBeacons() []Point {
	result := make([]Point, len(s.beacons))
	for i, b := range s.beacons {
		result[i] = b.Clone()
	}
	return result
}

func (s *Scanner) ToString() string {
	result := ""
	for _, x := range s.AbsoluteBeacons() {
		result += x.ToString() + "\n"
	}
	return result[0 : len(result)-1]
}

func (s *Scanner) Clone() *Scanner {
	beacons := make([]*Point, len(s.beacons))
	for i, b := range s.beacons {
		clone := b.Clone()
		beacons[i] = &clone
	}
	return &Scanner{
		Position:    s.Position.Clone(),
		beacons:     beacons,
		orientation: s.orientation.Clone(),
	}
}

func (s *Scanner) GetRotated(rule string) *Scanner {
	beacons := make([]*Point, len(s.beacons))
	for i, b := range s.beacons {
		clone := b.Clone()
		for _, r := range rule {
			if r == 'x' {
				clone.RotateX()
			} else if r == 'y' {
				clone.RotateY()
			} else if r == 'z' {
				clone.RotateZ()
			}
		}
		beacons[i] = &clone
	}
	return &Scanner{
		Position: s.Position.Clone(),
		beacons:  beacons,
	}
}

func (s *Scanner) AddBeacons(beacons []Point, deltaPoint Point) {
	m := make(map[Point]bool)
	for _, b := range s.beacons {
		m[*b] = true
	}
	for _, b := range beacons {
		d := deltaPoint.Plus(b)
		if _, exist := m[d]; !exist {
			s.beacons = append(s.beacons, &d)
		}
	}
}
