package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Today struct {
	sequence string
}

func (d *Today) Init(input string) error {
	s, err := lib.ReadFile(input)
	if err != nil {
		return err
	}

	d.sequence = s
	return nil
}

func (d *Today) isUnique(s string) bool {
	for i := 0; i < len(s); i++ {
		char := s[i : i+1]
		if strings.Count(s, char) > 1 {
			return false
		}
	}
	return true
}

func (d *Today) Part1() (string, error) {
	i := 4
	for ; i <= len(d.sequence); i++ {
		substr := d.sequence[i-4 : i]
		if d.isUnique(substr) {
			break
		}
	}

	return strconv.Itoa(i), nil
}

func (d *Today) Part2() (string, error) {
	i := 14
	for ; i <= len(d.sequence); i++ {
		substr := d.sequence[i-14 : i]
		if d.isUnique(substr) {
			break
		}
	}

	return strconv.Itoa(i), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
