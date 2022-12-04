package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Range struct {
	start int
	end   int
}

func NewRange(stringRange string) Range {
	parts := strings.Split(stringRange, "-")

	ret := Range{}
	var err error

	ret.start, err = strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	ret.end, err = strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return ret
}

func (r *Range) contains(r2 Range) bool {
	return r.start <= r2.start &&
		r.end >= r2.end
}

func (r *Range) overlaps(r2 Range) bool {
	return !(r.start > r2.end ||
		r.end < r2.start)
}

type Pair struct {
	first  Range
	second Range
}

func NewPair(line string) Pair {
	parts := strings.Split(line, ",")

	return Pair{
		first:  NewRange(parts[0]),
		second: NewRange(parts[1]),
	}
}

type Today struct {
	pairs []Pair
}

func (d *Today) Init(input string) error {
	pairs, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.pairs = make([]Pair, len(pairs))
	for i, pair := range pairs {
		d.pairs[i] = NewPair(pair)
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	count := 0
	for _, pair := range d.pairs {
		if pair.first.contains(pair.second) || pair.second.contains(pair.first) {
			count = count + 1
		}
	}

	return strconv.Itoa(count), nil
}

func (d *Today) Part2() (string, error) {
	count := 0
	for _, pair := range d.pairs {
		if pair.first.overlaps(pair.second) {
			count = count + 1
		}
	}

	return strconv.Itoa(count), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
