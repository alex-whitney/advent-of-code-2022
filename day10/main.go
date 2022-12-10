package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type screen [][]bool

func (s screen) draw(cycle int, register int) {
	c := (cycle - 1) % 40
	r := cycle / 40

	if c >= register-1 && c <= register+1 {
		s[r][c] = true
	}
}

func (s screen) toString() string {
	out := "\n"
	for _, row := range s {
		for _, val := range row {
			if val {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return out
}

type instruction struct {
	command string
	value   int
}

func newInstruction(line string) (*instruction, error) {
	parts := strings.Split(line, " ")
	instr := &instruction{
		command: parts[0],
	}

	if len(parts) > 1 {
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		instr.value = val
	}

	return instr, nil
}

type Today struct {
	instructions []*instruction
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.instructions = make([]*instruction, len(raw))
	for i, line := range raw {
		d.instructions[i], err = newInstruction(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	register := 1
	signalStrength := 0
	cycle := 1

	for _, instr := range d.instructions {
		if (cycle-20)%40 == 0 {
			signalStrength += register * cycle
		}

		if instr.command == "addx" {
			cycle += 1
			if (cycle-20)%40 == 0 {
				signalStrength += register * cycle
			}

			register += instr.value
		}

		cycle += 1
	}

	return strconv.Itoa(signalStrength), nil
}

func (d *Today) Part2() (string, error) {
	register := 1
	cycle := 1

	screen := make(screen, 6)
	for i := range screen {
		screen[i] = make([]bool, 40)
	}

	for _, instr := range d.instructions {
		screen.draw(cycle, register)

		if instr.command == "addx" {
			cycle += 1
			screen.draw(cycle, register)

			register += instr.value
		}

		cycle += 1
	}

	return screen.toString(), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
