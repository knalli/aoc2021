package day03

import (
	"fmt"
	"github.com/knalli/aoc"
	"strconv"
)

func solve1(lines []string) error {
	gammaRate := ""
	epsilonRate := ""
	lineLength := len(lines[0])
	for i := 0; i < lineLength; i++ {
		count1 := 0
		for _, line := range lines {
			if line[i] == '1' {
				count1++
			}
		}
		if count1 > len(lines)/2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	gammaRateInt := binary2Int(gammaRate)
	epsilonRateInt := binary2Int(epsilonRate)
	aoc.PrintSolution(fmt.Sprintf("gamma = %s (%d), epsilon = %s (%d)", gammaRate, gammaRateInt, epsilonRate, epsilonRateInt))
	aoc.PrintSolution(fmt.Sprintf("power consumption = %d", gammaRateInt*epsilonRateInt))
	return nil
}

func binary2Int(s string) int64 {
	if v, err := strconv.ParseInt(s, 2, 64); err != nil {
		return 0
	} else {
		return v
	}
}

func solve2(lines []string) error {
	oxygen := lines
	lineLength := len(lines[0])
	for i := 0; i < lineLength; i++ {
		oxygen = filter(oxygen, i, true)
		if len(oxygen) == 1 {
			break
		}
	}
	co2scrubber := lines
	for i := 0; i < lineLength; i++ {
		co2scrubber = filter(co2scrubber, i, false)
		if len(co2scrubber) == 1 {
			break
		}
	}
	oxygenInt := binary2Int(oxygen[0])
	co2scrubberInt := binary2Int(co2scrubber[0])
	aoc.PrintSolution(fmt.Sprintf("oxygen = %s (%d), co2 scrubber = %s (%d)", oxygen[0], oxygenInt, co2scrubber[0], co2scrubberInt))
	aoc.PrintSolution(fmt.Sprintf("life support rating = %d", oxygenInt*co2scrubberInt))
	return nil
}

func filter(lines []string, pos int, mostCommon bool) []string {
	result := make([]string, 0)
	count1 := 0
	for _, line := range lines {
		if line[pos] == '1' {
			count1++
		}
	}
	if mostCommon {
		// most common
		if count1*2 < len(lines) {
			for _, line := range lines {
				if line[pos] == '0' {
					result = append(result, line)
				}
			}
		} else {
			for _, line := range lines {
				if line[pos] == '1' {
					result = append(result, line)
				}
			}
		}
	} else {
		// least common
		if count1*2 < len(lines) {
			for _, line := range lines {
				if line[pos] == '1' {
					result = append(result, line)
				}
			}
		} else {
			for _, line := range lines {
				if line[pos] == '0' {
					result = append(result, line)
				}
			}
		}
	}
	return result
}
