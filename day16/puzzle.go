package day16

import (
	"fmt"
	"github.com/knalli/aoc"
	"math"
)

func sumVersions(message BitsMessage) int {
	res := message.Version
	if message.Type != 4 {
		for _, sm := range message.SubMessages {
			res += sumVersions(sm)
		}
	}
	return res
}

func calc(message BitsMessage) int {
	if message.Type == 0 {
		res := 0
		for _, s := range message.SubMessages {
			res += calc(s)
		}
		return res
	} else if message.Type == 1 {
		res := 1
		for _, s := range message.SubMessages {
			res *= calc(s)
		}
		return res
	} else if message.Type == 2 {
		res := math.MaxInt64
		for _, s := range message.SubMessages {
			res = aoc.MinInt(res, calc(s))
		}
		return res
	} else if message.Type == 3 {
		res := math.MinInt64
		for _, s := range message.SubMessages {
			res = aoc.MaxInt(res, calc(s))
		}
		return res
	} else if message.Type == 4 {
		return message.Value
	} else if message.Type == 5 {
		if calc(message.SubMessages[0]) > calc(message.SubMessages[1]) {
			return 1
		}
		return 0
	} else if message.Type == 6 {
		if calc(message.SubMessages[0]) < calc(message.SubMessages[1]) {
			return 1
		}
		return 0
	} else if message.Type == 7 {
		if calc(message.SubMessages[0]) == calc(message.SubMessages[1]) {
			return 1
		}
		return 0
	} else {
		panic("impossible")
	}
}

func solve1(lines []string) error {
	encoded := lines[0]
	decoded := decodeHex(encoded)
	message, err := parseBitsMessage(decoded)
	if err != nil {
		return err
	}
	//fmt.Println(message.ToString())
	aoc.PrintSolution(fmt.Sprintf("version sum = %d", sumVersions(*message)))
	return nil
}

func solve2(lines []string) error {
	encoded := lines[0]
	decoded := decodeHex(encoded)
	message, err := parseBitsMessage(decoded)
	if err != nil {
		return err
	}
	//fmt.Println(message.ToString())
	aoc.PrintSolution(fmt.Sprintf("calc result = %d", calc(*message)))
	return nil
}
