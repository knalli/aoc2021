package day14

import (
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"strings"
)

func countElements(str string) map[string]int {
	result := make(map[string]int)
	for _, c := range str {
		s := string(c)
		result[s] = result[s] + 1
	}
	return result
}

func solve1(lines []string) error {
	tpl := lines[0]
	rules := make(map[string]string)
	for _, line := range lines {
		if strings.Contains(line, "->") {
			parts := strings.Split(line, " -> ")
			rules[parts[0]] = parts[1]
		}
	}

	polymer := aoc.NewLinkedList(true)
	for _, c := range tpl {
		polymer.Add(string(c))
	}

	fmt.Printf("Template:       %s\n", tpl)
	for step := 1; step <= 10; step++ {
		last := ""
		polymer.Each(func(e *aoc.LinkElement) {
			val := e.Value().(string)
			if last != "" {
				s := last + val
				if rule, exist := rules[s]; exist {
					polymer.AddAfter(e.Prev(), rule)
				}
			}
			last = val
		})
		fmt.Printf("After step %03d: len=%d\n", step, len(polymer.ToString(func(e *aoc.LinkElement) string {
			return e.Value().(string)
		})))
	}

	finalPolymer := polymer.ToString(func(e *aoc.LinkElement) string {
		return e.Value().(string)
	})
	counts := countElements(finalPolymer)
	max := math.MinInt64
	min := math.MaxInt64
	for _, c := range counts {
		max = aoc.MaxInt(max, c)
		min = aoc.MinInt(min, c)
	}
	aoc.PrintSolution(fmt.Sprintf("max(c) - min(c) = %d", max-min))
	return nil
}

// another approach as solve1 won't work in finite time...
func solve2(lines []string) error {
	tpl := lines[0]
	fmt.Printf("Template:       %s\n", tpl)

	rules := make(map[string]string)
	for _, line := range lines {
		if strings.Contains(line, "->") {
			parts := strings.Split(line, " -> ")
			rules[parts[0]] = parts[1]
		}
	}

	elements := make(map[string]int)
	for _, c := range tpl {
		elements[string(c)]++
	}

	// the actual order of pair does not matter (because reactions will be inserted only, so left and right are persistent)
	pairs := make(map[string]int)
	for i := 0; i < len(tpl)-1; i++ {
		pairs[tpl[i:i+2]]++
	}

	for step := 0; step < 40; step++ {
		newPairs := make(map[string]int) // can be replaced
		for pair, v := range pairs {     // scaled via v, avoid iter/recur
			if reacted, exist := rules[pair]; exist {
				elements[reacted] += v
				newPairs[string(pair[0])+reacted] += v
				newPairs[reacted+string(pair[1])] += v
			} else {
				fmt.Println("unexpected")
			}
		}
		pairs = newPairs
	}

	max := math.MinInt64
	min := math.MaxInt64
	for _, c := range elements {
		max = aoc.MaxInt(max, c)
		min = aoc.MinInt(min, c)
	}
	aoc.PrintSolution(fmt.Sprintf("max(c) - min(c) = %d", max-min))
	return nil
}
