package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Today struct {
	flowRates map[string]int
	tunnels   map[string][]string
	distances map[string]map[string]int
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.flowRates = make(map[string]int)
	d.tunnels = make(map[string][]string)
	d.distances = make(map[string]map[string]int)
	for _, line := range raw {
		var valve string
		var flowRate int
		fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnel leads to valve", &valve, &flowRate)
		d.flowRates[valve] = flowRate

		parts := strings.Split(line, ", ")
		s := make([]string, len(parts))
		for i, part := range parts {
			s[i] = part[len(part)-2:]
		}
		d.tunnels[valve] = s

		d.distances[valve] = make(map[string]int)
		for _, other := range d.tunnels[valve] {
			d.distances[valve][other] = 1
		}
	}

	for start := range d.tunnels {
		neighborhood := slices.Clone(d.tunnels[start])
		for len(neighborhood) > 0 {
			next := neighborhood[0]
			neighborhood = neighborhood[1:]

			for _, nextnext := range d.tunnels[next] {
				if d.distances[start][nextnext] > 0 || nextnext == start {
					continue
				}

				d.distances[start][nextnext] = d.distances[start][next] + 1
				neighborhood = append(neighborhood, nextnext)
			}
		}
	}

	return nil
}

func (d *Today) next(n string, open map[string]int, closed []string, minute int, score int) (int, map[string]int) {
	maxScore := score
	result := open
	for _, next := range closed {
		if open[next] > 0 {
			continue
		}

		dist := d.distances[n][next]
		nextRound := dist + minute + 1
		if nextRound >= 30 {
			continue
		}

		a := maps.Clone(open)
		a[next] = minute + dist + 1

		additionalScore := (30 - nextRound) * d.flowRates[next]
		newScore, newOpen := d.next(next, a, closed, nextRound, score+additionalScore)

		if maxScore < newScore {
			maxScore = newScore
			result = newOpen
		}
	}

	return maxScore, result
}

func (d *Today) next2(open1 map[string]int, open2 map[string]int, closed []string, score int, mScore *int) (int, map[string]int, map[string]int) {
	maxScore := score
	r1 := open1
	r2 := open2
	for _, next := range closed {
		if open1[next] > 0 || open2[next] > 0 {
			continue
		}

		a := maps.Clone(open1)
		b := maps.Clone(open2)

		aMaxMinute := lib.Max(maps.Values(a))
		bMaxMinute := lib.Max(maps.Values(b))

		var n string
		minute := aMaxMinute
		addTo := a
		if bMaxMinute < aMaxMinute {
			addTo = b
			minute = bMaxMinute
		}
		for k, v := range addTo {
			if v == minute {
				n = k
			}
		}

		dist := d.distances[n][next]
		nextRound := dist + minute + 1
		if nextRound >= 26 {
			continue
		}

		addTo[next] = minute + dist + 1

		additionalScore := (26 - nextRound) * d.flowRates[next]
		newScore, result1, result2 := d.next2(a, b, closed, score+additionalScore, mScore)

		if maxScore < newScore {
			maxScore = newScore
			r1 = result1
			r2 = result2
		}
	}

	if maxScore > *mScore {
		*mScore = maxScore
		fmt.Printf("Highest score so far: %d\n", maxScore)
	}

	return maxScore, r1, r2
}

func (d *Today) Part1() (string, error) {
	position := "AA"
	closed := make([]string, 0)
	for node := range d.tunnels {
		if d.flowRates[node] > 0 {
			closed = append(closed, node)
		}
	}

	val, result := d.next(position, map[string]int{}, closed, 0, 0)
	fmt.Printf("%+v\n", result)

	return strconv.Itoa(val), nil
}

func (d *Today) Part2() (string, error) {
	closed := make([]string, 0)
	for node := range d.tunnels {
		if d.flowRates[node] > 0 {
			closed = append(closed, node)
		}
	}

	a := map[string]int{"AA": 0}
	b := map[string]int{"AA": 0}

	// so this runs for forever. more optimization needed, but I just let this run and
	// eventually it spat out the right answer.
	ms := 0
	val, r1, r2 := d.next2(a, b, closed, 0, &ms)

	fmt.Printf("%+v\n%+v\n", r1, r2)

	return strconv.Itoa(val), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
