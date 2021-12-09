package day08

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

type InputType struct {
	SignalPatterns []string
	OutputValues   []string
}

func parseLine(line string) InputType {
	split := strings.Split(line, "|")
	signalPatterns := strings.Split(strings.TrimSpace(split[0]), " ")
	outputValues := strings.Split(strings.TrimSpace(split[1]), " ")
	return InputType{
		SignalPatterns: signalPatterns,
		OutputValues:   outputValues,
	}
}

func parseLines(lines []string) []InputType {
	result := make([]InputType, 0)
	for _, line := range lines {
		result = append(result, parseLine(line))
	}
	return result
}

type DigitType struct {
	Value   int
	Signals map[int32]bool
}

type WireOptions struct {
	Value string
}

func (w *WireOptions) SetValue(v string) {
	w.Value = v
}

func findFirstCharByCount(groups []string, count int) int32 {
	counters := make(map[int32]int)

	for _, group := range groups {
		for _, c := range group {
			counters[c] = counters[c] + 1
			if counters[c] == count {
				return c
			}
		}
	}

	panic("not found")
}

func solve(lines []string, chars string, scores map[int32]int, digitSpecs []DigitType) (int, int) {
	countSimpleDigits := 0
	sumTotal := 0

	for _, input := range parseLines(lines) {
		wires := make(map[int32]*WireOptions, 0)
		for _, c := range chars {
			wires[c] = &WireOptions{
				Value: chars,
			}
		}

		completedSegments := make([]int32, 0)
		// Determine most common by scoring only (segment f, 9 occurrences)
		mostCommonWireInput := findFirstCharByCount(input.SignalPatterns, scores['f'])
		// ... this segment must be "f".
		wires[mostCommonWireInput].SetValue("f")
		completedSegments = append(completedSegments, mostCommonWireInput)
		// This means, the other segment of digit 1 must be the segment "c". Just find the short 2-length item.
		for _, pattern := range input.SignalPatterns {
			if len(pattern) == 2 {
				if int32(pattern[0]) == mostCommonWireInput {
					c := int32(pattern[1])
					wires[c].SetValue("c")
					completedSegments = append(completedSegments, c)
				} else {
					c := int32(pattern[0])
					wires[c].SetValue("c")
					completedSegments = append(completedSegments, c)
				}
				break
			}
		}
		// Digit 7 has the same as 1, but additional segment "a", find via the 3-length item.
		for _, pattern := range input.SignalPatterns {
			if len(pattern) == 3 {
				for _, c := range pattern {
					found := false
					// filter c not in completedSegments
					for _, cs := range completedSegments {
						if c == cs {
							found = true
							break
						}
					}
					if !found {
						wires[c].SetValue("a")
						completedSegments = append(completedSegments, c)
						break
					}
				}
				break
			}
		}
		// The next shortest digit is 4 (4 segments), of which c and f are already discovered.
		// This left segments b and d can be found when
		// Of the 4-length item only two options are possible and of all 6-length items (0, 6, 9)
		// is b always and d only 2 times set.
		{
			candidates := make(map[int32]int)
			// find the two candidates left
			for _, pattern := range input.SignalPatterns {
				if len(pattern) == 4 {
					for _, c := range pattern {
						found := false
						// filter c not in completedSegments
						for _, cs := range completedSegments {
							if cs == c {
								found = true
								break
							}
						}
						if !found {
							candidates[c] = 0
						}
					}

					break
				}
			}
			// Look which of them is in every 6-length group
			patternCnt := 0
			for _, pattern := range input.SignalPatterns {
				if len(pattern) == 6 {
					for _, c := range pattern {
						for cc := range candidates {
							if cc == c {
								candidates[cc] = candidates[cc] + 1
							}
						}
					}
					patternCnt++
				}
			}
			// The character found always is "b"
			// ... which means the other one must be "d"
			for cc, ct := range candidates {
				if ct == patternCnt {
					wires[cc].SetValue("b")
					completedSegments = append(completedSegments, cc)
				} else {
					wires[cc].SetValue("d")
					completedSegments = append(completedSegments, cc)
				}
			}
		}
		// Now only segment "e" and "f" are left.
		// "g" has a score of 7, "e" only 4. For the incompleted ones, count the occurrences again.
		{
			candidates := make(map[int32]int)
			// find the candidates left (more than 1 option)
			for inChar, outOptions := range wires {
				if len(outOptions.Value) > 1 {
					candidates[inChar] = 0
				}
			}
			// Look which of them is in every 6-length group
			for _, pattern := range input.SignalPatterns {
				for _, c := range pattern {
					for cc := range candidates {
						if cc == c {
							candidates[cc] = candidates[cc] + 1
						}
					}
				}
			}
			// The character found 7-times is for segment "g"
			// ... which means the other one must be "e"
			for cc, ct := range candidates {
				if ct == scores['g'] {
					wires[cc].SetValue("g")
					completedSegments = append(completedSegments, cc)
				} else {
					wires[cc].SetValue("e")
					completedSegments = append(completedSegments, cc)
				}
			}
		}
		optimizeOptions(wires)
		//renderWires(wires)

		// validate
		for _, opts := range wires {
			if len(opts.Value) != 1 {
				panic("invalid wiring")
			}
		}

		numOutput := ""
		for _, signalStr := range input.OutputValues {
			res := ""
			for _, signalChr := range signalStr {
				res += wires[signalChr].Value
			}
			d := findDigit(digitSpecs, res)

			if d == 1 || d == 4 || d == 7 || d == 8 {
				countSimpleDigits++
			}
			numOutput += fmt.Sprintf("%d", d)
		}
		sumTotal += aoc.ParseInt(numOutput)
	}
	return countSimpleDigits, sumTotal
}

func optimizeOptions(wires map[int32]*WireOptions) {
	stickedChars := make([]int32, 0)
	for _, wireOptions := range wires {
		if len(wireOptions.Value) == 1 {
			stickedChars = append(stickedChars, int32(wireOptions.Value[0]))
		}
	}
	for _, wireOptions := range wires {
		for _, stickedChar := range stickedChars {
			if len(wireOptions.Value) != 1 && strings.Contains(wireOptions.Value, string(stickedChar)) {
				res := ""
				for _, c := range wireOptions.Value {
					if c != stickedChar {
						res += string(c)
					}
				}
				wireOptions.Value = res
			}
		}
	}
}

func renderWires(wires map[int32]*WireOptions) {
	fmt.Printf("====\n")
	for wChar, wOptions := range wires {
		fmt.Printf("%s => %s\n", string(wChar), wOptions.Value)
	}
	fmt.Printf("====\n")
}

func findDigit(specs []DigitType, search string) int {
	for _, spec := range specs {
		if len(spec.Signals) != len(search) {
			continue
		}
		found := true
		for c, b := range spec.Signals {
			if b && !strings.Contains(search, string(c)) {
				found = false
				break
			}
		}
		if !found {
			continue
		}
		return spec.Value
	}
	return -1
}

func solve1(lines []string) error {
	digitSpecs := buildDigitSpecs()
	chars := "abcdefg"

	scores := make(map[int32]int, 0)
	for _, spec := range digitSpecs {
		for c := range spec.Signals {
			scores[c] = scores[c] + 1
		}
	}

	countSimpleDigits, _ := solve(lines, chars, scores, digitSpecs)

	aoc.PrintSolution(fmt.Sprintf("Simple digits apperas %d times", countSimpleDigits))

	return nil
}

func solve2(lines []string) error {
	digitSpecs := buildDigitSpecs()
	chars := "abcdefg"

	scores := make(map[int32]int, 0)
	for _, spec := range digitSpecs {
		for c := range spec.Signals {
			scores[c] = scores[c] + 1
		}
	}

	_, sumTotal := solve(lines, chars, scores, digitSpecs)

	aoc.PrintSolution(fmt.Sprintf("Total sum %d", sumTotal))

	return nil
}
