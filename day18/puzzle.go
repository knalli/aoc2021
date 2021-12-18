package day18

import (
	"fmt"
	"github.com/knalli/aoc"
)

func solve1(lines []string) error {
	var t *SfPair
	for _, line := range lines {
		pair := parseSfPair(line)
		pair.Reduce()
		if t == nil {
			t = pair
		} else {
			//fmt.Printf("  %s\n", t.ToString())
			//fmt.Printf("+ %s\n", pair.ToString())
			t = t.Add(pair)
			t.Reduce()
			//fmt.Printf("= %s\n", t.ToString())
		}
	}
	aoc.PrintSolution(fmt.Sprintf("final sum magnitude = %d", t.Magnitude()))
	return nil
}

func solve2(lines []string) error {
	max := 0
	for i1, l1 := range lines {
		for i2, l2 := range lines {
			if i1 == i2 {
				continue
			}
			p1 := parseSfPair(l1)
			p2 := parseSfPair(l2)
			p := p1.Add(p2)
			p.Reduce()
			m := p.Magnitude()
			if m > max {
				max = m
			}
		}
	}

	aoc.PrintSolution(fmt.Sprintf("max magnitude = %d", max))

	return nil
}
