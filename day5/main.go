package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Instruction struct {
	count int
	from  int
	to    int
}

func NewInstruction(line string) (*Instruction, error) {
	parts := strings.Split(line, " ")

	count, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	from, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}

	to, err := strconv.Atoi(parts[5])
	if err != nil {
		return nil, err
	}

	return &Instruction{
		count: count,
		from:  from,
		to:    to,
	}, nil
}

type Today struct {
	instructions []*Instruction
	stacks       [][]string
}

func (d *Today) Init(input string) error {
	fullInput, err := lib.ReadFile(input)
	if err != nil {
		return err
	}

	sections := strings.Split(fullInput, "\n\n")

	stacks := strings.Split(sections[0], "\n")
	nStacks := (len(stacks[len(stacks)-1]) + 1) / 4
	d.stacks = make([][]string, nStacks)
	for i := len(stacks) - 2; i >= 0; i-- {
		row := []rune(stacks[i])
		for stack := 0; stack < nStacks; stack++ {
			val := row[stack*4+1]
			if val != ' ' {
				d.stacks[stack] = append(d.stacks[stack], string(val))
			}
		}
	}

	lines := strings.Split(sections[1], "\n")
	d.instructions = make([]*Instruction, len(lines))
	for i, line := range lines {
		d.instructions[i], err = NewInstruction(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	stacks := make([][]string, len(d.stacks))
	for i, s := range d.stacks {
		stacks[i] = make([]string, len(s))
		copy(stacks[i], s)
	}

	for _, instruction := range d.instructions {
		for i := 0; i < instruction.count; i++ {
			from := instruction.from - 1
			lenFrom := len(stacks[from])

			to := instruction.to - 1

			stacks[to] = append(stacks[to], stacks[from][lenFrom-1])
			stacks[from] = stacks[from][:lenFrom-1]
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result += stacks[i][len(stacks[i])-1]
	}

	return result, nil
}

func (d *Today) Part2() (string, error) {
	for _, instruction := range d.instructions {
		from := instruction.from - 1
		lenFrom := len(d.stacks[from])

		to := instruction.to - 1

		d.stacks[to] = append(d.stacks[to], d.stacks[from][lenFrom-instruction.count:]...)
		d.stacks[from] = d.stacks[from][:lenFrom-instruction.count]
	}

	result := ""
	for i := 0; i < len(d.stacks); i++ {
		result += d.stacks[i][len(d.stacks[i])-1]
	}

	return result, nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
