package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type monkey struct {
	number          int
	items           []int
	operator        string
	operatorValue   int
	testValue       int
	testTrueMonkey  int
	testFalseMonkey int
}

func (m *monkey) toString() string {
	ret := "Monkey " + strconv.Itoa(m.number) + ": "
	for _, item := range m.items {
		ret += " " + strconv.Itoa(item)
	}
	return ret
}

func newMonkey(block string) (*monkey, error) {
	ret := &monkey{}
	var err error

	lines := strings.Split(block, "\n")

	ret.number, err = strconv.Atoi(lines[0][7 : len(lines[0])-1])
	if err != nil {
		return nil, err
	}

	items := strings.Split(lines[1][18:], ", ")
	ret.items = make([]int, len(items))
	for i, item := range items {
		ret.items[i], err = strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
	}

	ret.operator = lines[2][23:24]

	if lines[2][25:] == "old" {
		// let's just cheat
		ret.operatorValue = -1
	} else {
		ret.operatorValue, err = strconv.Atoi(lines[2][25:])
		if err != nil {
			return nil, err
		}
	}

	ret.testValue, err = strconv.Atoi(lines[3][21:])
	if err != nil {
		return nil, err
	}

	ret.testTrueMonkey, err = strconv.Atoi(lines[4][29:])
	if err != nil {
		return nil, err
	}

	ret.testFalseMonkey, err = strconv.Atoi(lines[5][30:])
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type Today struct {
	monkeys []*monkey
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadFile(input)
	if err != nil {
		return err
	}

	rawMonkeys := strings.Split(raw, "\n\n")
	d.monkeys = make([]*monkey, len(rawMonkeys))
	for i, m := range rawMonkeys {
		d.monkeys[i], err = newMonkey(m)
		if err != nil {
			return err
		}
	}

	return err
}

func (d *Today) Part1() (string, error) {
	/*
		counter := make([]int, len(d.monkeys))

		for round := 0; round < 20; round++ {
			for i, monkey := range d.monkeys {
				for _, item := range monkey.items {
					counter[i] += 1

					if monkey.operator == "*" {
						if monkey.operatorValue == -1 {
							item = item * item
						} else {
							item = item * monkey.operatorValue
						}
					} else {
						item = item + monkey.operatorValue
					}

					item = item / 3

					destIndex := monkey.testFalseMonkey
					if item%monkey.testValue == 0 {
						destIndex = monkey.testTrueMonkey
					}
					d.monkeys[destIndex].items = append(d.monkeys[destIndex].items, item)
				}
				monkey.items = make([]int, 0)
			}
		}

		sort.Ints(counter)
		counter = lib.Reverse(counter)

		return strconv.Itoa(counter[0] * counter[1]), nil
	*/
	return "skipped", nil
}

func (d *Today) Part2() (string, error) {
	counter := make([]int, len(d.monkeys))

	factor := 1
	for _, m := range d.monkeys {
		factor = factor * m.testValue
	}

	for round := 0; round < 10000; round++ {
		for i, monkey := range d.monkeys {
			for _, item := range monkey.items {
				counter[i] += 1

				if monkey.operator == "*" {
					if monkey.operatorValue == -1 {
						item = item * item
					} else {
						item = item * monkey.operatorValue
					}
				} else {
					item = item + monkey.operatorValue
				}

				item = item % factor

				destIndex := monkey.testFalseMonkey
				if item%monkey.testValue == 0 {
					destIndex = monkey.testTrueMonkey
				}
				d.monkeys[destIndex].items = append(d.monkeys[destIndex].items, item)
			}
			monkey.items = make([]int, 0)
		}
	}

	sort.Ints(counter)
	counter = lib.Reverse(counter)

	return strconv.Itoa(counter[0] * counter[1]), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
