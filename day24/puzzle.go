package day24

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

type Instruction struct {
	OpCode string
	Arg1   string
	Arg2   string
}

type AluRegister struct {
	W int
	X int
	Y int
	Z int
}

// Apply unused
func (r *AluRegister) Apply(ins *Instruction, input <-chan int) {
	setValue := func(name string, value int) error {
		if name == "w" {
			r.W = value
		} else if name == "x" {
			r.X = value
		} else if name == "y" {
			r.Y = value
		} else if name == "z" {
			r.Z = value
		} else {
			return errors.New("invalid instruction")
		}
		return nil
	}
	getValue := func(name string) (int, error) {
		if name == "w" {
			return r.W, nil
		} else if name == "x" {
			return r.X, nil
		} else if name == "y" {
			return r.Y, nil
		} else if name == "z" {
			return r.Z, nil
		} else {
			return 0, errors.New("invalid instruction")
		}
	}
	if ins.OpCode == "inp" {
		value := <-input
		if err := setValue(ins.Arg1, value); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "add" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1+value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "mul" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1*value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "div" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1/value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "mod" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1%value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "eql" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		result := 1
		if value1 != value2 {
			result = 0
		}
		if err := setValue(ins.Arg1, result); err != nil {
			panic(err)
		}
	}
}

func (r *AluRegister) ApplyVal(ins *Instruction, input int) {
	setValue := func(name string, value int) error {
		if name == "w" {
			r.W = value
		} else if name == "x" {
			r.X = value
		} else if name == "y" {
			r.Y = value
		} else if name == "z" {
			r.Z = value
		} else {
			return errors.New("invalid instruction")
		}
		return nil
	}
	getValue := func(name string) (int, error) {
		if name == "w" {
			return r.W, nil
		} else if name == "x" {
			return r.X, nil
		} else if name == "y" {
			return r.Y, nil
		} else if name == "z" {
			return r.Z, nil
		} else {
			return 0, errors.New("invalid instruction")
		}
	}
	if ins.OpCode == "inp" {
		value := input
		if err := setValue(ins.Arg1, value); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "add" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1+value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "mul" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1*value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "div" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1/value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "mod" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		if err := setValue(ins.Arg1, value1%value2); err != nil {
			panic(err)
		}
	} else if ins.OpCode == "eql" {
		value1, err := getValue(ins.Arg1)
		if err != nil {
			panic(err)
		}
		value2, err := getValue(ins.Arg2)
		if err != nil {
			value2 = aoc.ParseInt(ins.Arg2)
		}
		result := 1
		if value1 != value2 {
			result = 0
		}
		if err := setValue(ins.Arg1, result); err != nil {
			panic(err)
		}
	}
}

func (r *AluRegister) ToString() string {
	return fmt.Sprintf("W = %13d X = %13d Y = %13d Z = %13d", r.W, r.X, r.Y, r.Z)
}

func (r *AluRegister) Clone() *AluRegister {
	return &AluRegister{
		W: r.W,
		X: r.X,
		Y: r.Y,
		Z: r.Z,
	}
}

func solve1(lines []string) error {

	instructions := make([]Instruction, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")
		code := split[0]
		arg1 := split[1]
		arg2 := ""
		if len(split) > 2 {
			arg2 = split[2]
		}
		instructions = append(instructions, Instruction{
			OpCode: code,
			Arg1:   arg1,
			Arg2:   arg2,
		})
	}

	max, min := findMaxMin(instructions)

	aoc.PrintSolution(fmt.Sprintf("Max number = %d", max))
	aoc.PrintSolution(fmt.Sprintf("Min number = %d", min))

	return nil
}

func solve2(lines []string) error {
	aoc.PrintSolution("See part1")
	return nil
}

func findMaxMin(instructions []Instruction) (int, int) {

	type AluContainer struct {
		Register *AluRegister
		InputMax int
		InputMin int
	}

	alus := make([]*AluContainer, 0)
	alus = append(alus, &AluContainer{
		Register: &AluRegister{},
	})
	for _, ins := range instructions {
		if ins.OpCode == "inp" {
			// branch for each possible input
			newAlus := make([]*AluContainer, 0)
			indexes := make(map[AluRegister]int)
			for _, alu := range alus {
				for n := 1; n < 10; n++ {
					newAlu := AluContainer{
						Register: alu.Register.Clone(),
						InputMax: alu.InputMax*10 + n, // extend to the right
						InputMin: alu.InputMin*10 + n, // extend to the right
					}
					newAlu.Register.ApplyVal(&ins, n)
					// TODO If last ins (for z found), we could use a filter z==0
					if idx, found := indexes[*newAlu.Register]; found {
						// if the same register state was found already, use only the extreme one
						// if 2 was already found, but 4 has the same output: use 4 (searching for the highest number)
						newAlus[idx].InputMax = aoc.MaxInt(newAlus[idx].InputMax, newAlu.InputMax)
						newAlus[idx].InputMin = aoc.MinInt(newAlus[idx].InputMin, newAlu.InputMin)
					} else {
						indexes[*newAlu.Register] = len(newAlus)
						newAlus = append(newAlus, &newAlu)
					}
				}
			}
			alus = newAlus
			fmt.Printf("len alus = %d\n", len(alus))
		} else {
			for _, alu := range alus {
				alu.Register.Apply(&ins, nil)
			}
		}
	}

	maxValues := make([]int, 0)
	minValues := make([]int, 0)
	for _, alu := range alus {
		if alu.Register.Z == 0 {
			maxValues = append(maxValues, alu.InputMax)
			minValues = append(minValues, alu.InputMin)
		}
	}

	return aoc.MaxIntArrayValue(maxValues), aoc.MinIntArrayValue(minValues)
}

// old unused
func run(modelNumber int, instructions []Instruction) bool {
	input := make(chan int)
	s := fmt.Sprintf("%d", modelNumber)
	if strings.Contains(s, "0") {
		return false
	}
	go func() {
		for _, c := range s {
			input <- int(c) - 48
		}
		close(input)
	}()

	register := AluRegister{}
	//fmt.Printf("%s\n", register.ToString())

	for _, ins := range instructions {
		register.Apply(&ins, input)
		fmt.Printf("%s\n", register.ToString())
	}
	fmt.Printf("z=%d\n", register.Z)

	return register.Z == 0
}
