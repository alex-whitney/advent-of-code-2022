package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Rucksack struct {
	allItems          string
	firstCompartment  string
	secondCompartment string
	commonItem        rune
}

type Today struct {
	rucksacks []Rucksack
}

func (d *Today) Init(input string) error {
	in, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.rucksacks = make([]Rucksack, len(in))
	for i, line := range in {
		d.rucksacks[i] = Rucksack{}
		rucksack := &d.rucksacks[i]

		n := len(line)
		rucksack.allItems = line
		rucksack.firstCompartment = line[:(n / 2)]
		rucksack.secondCompartment = line[n/2:]

		common := make(map[rune]bool)
		for _, char := range rucksack.firstCompartment {
			if strings.ContainsRune(rucksack.secondCompartment, char) {
				common[char] = true
			}
		}

		if len(common) != 1 {
			return errors.New("something went wrong: " + line)
		}
		for k := range common {
			rucksack.commonItem = k
		}
	}

	return nil
}

func (d *Today) priority(char rune) int {
	// upper case
	if char < 91 {
		return int(char) - 64 + 26
	} else {
		return int(char) - 96
	}
}

func (d *Today) Part1() (string, error) {
	sum := 0
	for _, rucksack := range d.rucksacks {
		sum += d.priority(rucksack.commonItem)
	}

	return strconv.Itoa(sum), nil
}

func (d *Today) Part2() (string, error) {
	counter := 0
	sum := 0
	for counter < len(d.rucksacks) {
		first := d.rucksacks[counter]
		for _, char := range first.allItems {
			if strings.ContainsRune(d.rucksacks[counter+1].allItems, char) &&
				strings.ContainsRune(d.rucksacks[counter+2].allItems, char) {
				sum = sum + d.priority(char)
				break
			}
		}
		counter += 3
	}

	return strconv.Itoa(sum), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
