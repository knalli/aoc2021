package day06

import (
	"fmt"
	"github.com/knalli/aoc"
)

type wrapper struct {
	Value int
}

func (w *wrapper) Increment() {
	w.Value++
}

func (w *wrapper) Decrement() {
	w.Value--
}

func (w *wrapper) ToString() string {
	return fmt.Sprintf("%d", w.Value)
}

func populate(lines []string, lastDay int) error {
	list := aoc.NewLinkedList(true)
	for _, n := range aoc.ParseInts(lines[0], ",") {
		list.Add(&wrapper{
			Value: n,
		})
	}
	fmt.Printf("Initial state: %s\n", renderList(list))
	trace := false
	debug := true
	statBefore := 0
	for day := 1; day <= lastDay; day++ {
		created := make([]*wrapper, 0)
		list.Each(func(e *aoc.LinkElement) {
			w := e.Value().(*wrapper)
			if w.Value == 0 {
				w.Value = 6
				created = append(created, &wrapper{
					Value: w.Value + 2,
				})
			} else {
				w.Decrement()
			}
		})
		for _, create := range created {
			list.AddAfter(list.Back(), create)
		}
		if trace {
			fmt.Printf("After %2d days: %s\n", day, renderList(list))
		} else if debug {
			fmt.Printf("After %2d days: %d fishs\n", day, list.Len())
		}

		if statBefore > 0 {
			statNow := list.Len()
			fmt.Printf("∂ +%d (%2.8f)\n", statNow-statBefore, float64(statNow)/float64(statBefore))
		}

		statBefore = list.Len()
	}
	aoc.PrintSolution(fmt.Sprintf("After %2d days: %d fishs\n", lastDay, list.Len()))
	return nil
}

func renderList(list *aoc.LinkedList) string {
	return list.ToString(func(e *aoc.LinkElement) string {
		w := e.Value().(*wrapper)
		s := w.ToString()
		if e.Next() != list.Front() {
			s += ","
		}
		return s
	})
}

func populate2(lines []string, lastDay int) error {
	days := make(map[int]int)
	total := int64(0)

	// initial fill
	for _, n := range aoc.ParseInts(lines[0], ",") {
		total++
		for d := n; d < lastDay; d += 7 {
			days[d]++
		}
	}
	// for each day
	// * take the daily spawns
	// * add them to total
	// * add the spawns in all consecutive day loops
	for d := 1; d <= lastDay; d++ {
		if spawns, found := days[d]; found {
			total += int64(spawns)
			for i := d + 9; i < lastDay; i += 7 {
				days[i] += spawns
			}
		}
	}

	aoc.PrintSolution(fmt.Sprintf("After %2d days: %d fishs\n", lastDay, total))
	return nil
}

func solve1(lines []string) error {
	return populate(lines, 80)
}

func solve2(lines []string) error {
	return populate2(lines, 256)
}
