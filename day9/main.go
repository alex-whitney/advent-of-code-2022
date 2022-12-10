package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type position struct {
	x int
	y int
}

func (tail position) follow(head position) position {
	dx := head.x - tail.x
	dy := head.y - tail.y

	if dx <= 1 && dx >= -1 && dy <= 1 && dy >= -1 {
		return tail
	}

	if dx > 1 {
		tail.x += 1
		if math.Abs(float64(dy)) <= 1 {
			tail.y += dy
		}
	}
	if dx < -1 {
		tail.x -= 1
		if math.Abs(float64(dy)) <= 1 {
			tail.y += dy
		}
	}
	if dy > 1 {
		tail.y += 1
		if math.Abs(float64(dx)) <= 1 {
			tail.x += dx
		}
	}
	if dy < -1 {
		tail.y -= 1
		if math.Abs(float64(dx)) <= 1 {
			tail.x += dx
		}
	}

	return tail
}

func (p *position) hash() string {
	return fmt.Sprintf("%d_%d", p.x, p.y)
}

type instruction struct {
	direction string
	distance  int
}

func (i *instruction) step(p position) position {
	switch i.direction {
	case "U":
		p.y += 1
		break
	case "D":
		p.y -= 1
		break
	case "L":
		p.x -= 1
		break
	case "R":
		p.x += 1
		break
	default:
		panic("unknown direction: " + i.direction)
	}

	return p
}

func newInstruction(line string) (*instruction, error) {
	parts := strings.Split(line, " ")
	d, err := strconv.Atoi(parts[1])
	return &instruction{
		direction: parts[0],
		distance:  d,
	}, err
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
	head := position{x: 0, y: 0}
	tail := position{x: 0, y: 0}

	visited := make(map[string]bool)
	visited[tail.hash()] = true

	for _, instr := range d.instructions {
		for n := 0; n < instr.distance; n++ {
			head = instr.step(head)
			tail = tail.follow(head)
			visited[tail.hash()] = true
		}
	}

	return strconv.Itoa(len(visited)), nil
}

func (d *Today) Part2() (string, error) {
	knots := make([]position, 10)

	visited := make(map[string]bool)
	visited[knots[9].hash()] = true

	for _, instr := range d.instructions {
		for n := 0; n < instr.distance; n++ {
			knots[0] = instr.step(knots[0])
			for i := 1; i < len(knots); i++ {
				knots[i] = knots[i].follow(knots[i-1])
			}
			visited[knots[9].hash()] = true
		}
	}

	return strconv.Itoa(len(visited)), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
