package main

import (
	"strconv"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Action int

const (
	Rock     Action = 1
	Paper           = 2
	Scissors        = 3
)

type Today struct {
	rounds   int
	opponent []Action
	guide    []Action
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadDelimitedFile(input, " ")
	if err != nil {
		return err
	}

	d.rounds = len(in)
	d.opponent = make([]Action, d.rounds)
	d.guide = make([]Action, d.rounds)

	for i, inputs := range in {
		d.opponent[i] = map[string]Action{
			"A": Rock,
			"B": Paper,
			"C": Scissors,
		}[inputs[0]]
		d.guide[i] = map[string]Action{
			"X": Rock,
			"Y": Paper,
			"Z": Scissors,
		}[inputs[1]]
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	score := 0
	for round := 0; round < d.rounds; round++ {
		played := d.guide[round]
		score += int(played)

		opponent := d.opponent[round]

		win := (played == Rock && opponent == Scissors) ||
			(played == Scissors && opponent == Paper) ||
			(played == Paper && opponent == Rock)
		if win {
			score += 6
		} else if played == opponent {
			score += 3
		}
	}

	return strconv.Itoa(score), nil
}

func (d *Today) Part2() (string, error) {
	score := 0
	for round := 0; round < d.rounds; round++ {
		opponent := d.opponent[round]

		played := d.guide[round]
		if played == Rock {
			played = map[Action]Action{
				Rock:     Scissors,
				Paper:    Rock,
				Scissors: Paper,
			}[opponent]
		} else if played == Paper {
			played = opponent
		} else {
			played = map[Action]Action{
				Scissors: Rock,
				Rock:     Paper,
				Paper:    Scissors,
			}[opponent]
		}

		score += int(played)

		win := (played == Rock && opponent == Scissors) ||
			(played == Scissors && opponent == Paper) ||
			(played == Paper && opponent == Rock)
		if win {
			score += 6
		} else if played == opponent {
			score += 3
		}
	}

	return strconv.Itoa(score), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
