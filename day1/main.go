package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Today struct {
	elves [][]int
}

func (d *Today) Init(input string) error {
	fileContents, err := lib.ReadFile(input)
	elfs := strings.Split(fileContents, "\n\n")
	if err != nil {
		return err
	}

	d.elves = make([][]int, len(elfs))
	for i, elfInput := range elfs {
		result, err := lib.ParseIntegerSlice(elfInput, "\n")
		if err != nil {
			return err
		}

		d.elves[i] = result
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	max := 0

	for _, elf := range d.elves {
		total := lib.Sum(elf)
		if total > max {
			max = total
		}
	}

	return strconv.Itoa(max), nil
}

func (d *Today) Part2() (string, error) {
	n := len(d.elves)

	totals := make([]int, n)
	for i, elf := range d.elves {
		totals[i] = lib.Sum(elf)
	}

	sort.Ints(totals)

	return strconv.Itoa(totals[n-1] + totals[n-2] + totals[n-3]), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
