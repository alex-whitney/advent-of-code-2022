package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Today struct {
	sensors   []lib.Pair[int, int]
	beacons   []lib.Pair[int, int]
	distances []int

	maxX int
	minX int

	part1Target int
	part2Size   int
}

func hashPoint(p lib.Pair[int, int]) string {
	return fmt.Sprintf("%d,%d", p.Left, p.Right)
}

func distance(p1 lib.Pair[int, int], p2 lib.Pair[int, int]) int {
	return int(math.Abs(float64(p2.Left-p1.Left)) + math.Abs(float64(p2.Right-p1.Right)))
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.sensors = make([]lib.Pair[int, int], len(raw))
	d.beacons = make([]lib.Pair[int, int], len(raw))
	d.distances = make([]int, len(raw))
	for i, line := range raw {
		var sx, sy, bx, by int
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil {
			return err
		}

		d.sensors[i] = lib.NewPair(sx, sy)
		d.beacons[i] = lib.NewPair(bx, by)
		d.distances[i] = distance(d.sensors[i], d.beacons[i])

		x := int(math.Max(float64(sx), float64(bx)))
		if d.maxX < x {
			d.maxX = x
		}
		x = int(math.Min(float64(sx), float64(bx)))
		if d.minX > x {
			d.minX = x
		}
	}

	if input == "/sample.txt" {
		d.part1Target = 10
		d.part2Size = 20
	} else {
		d.part1Target = 2000000
		d.part2Size = 4000000
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	counter := 0

	pt := lib.NewPair(0, d.part1Target)
	for x := d.minX * 2; x <= d.maxX*2; x++ {
		inBounds := false
		pt.Left = x
		for i, sensor := range d.sensors {
			if pt.Left == d.beacons[i].Left && pt.Right == d.beacons[i].Right {
				inBounds = false
				break
			}

			dist := distance(pt, sensor)
			if dist <= d.distances[i] {
				inBounds = true
			}
		}
		if inBounds {
			counter += 1
		}
	}

	return strconv.Itoa(counter), nil
}

func (d *Today) Part2() (string, error) {
	pt := lib.NewPair(0, 0)
	for x := 0; x <= d.part2Size; x++ {
		for y := 0; y <= d.part2Size; y++ {
			pt.Left = x
			pt.Right = y

			valid := true
			for i, sensor := range d.sensors {
				dist := distance(pt, sensor)
				if dist <= d.distances[i] {
					valid = false

					// skip to the end of the sensor's range
					y += d.distances[i] - dist
					break
				}
			}

			if valid {
				return strconv.Itoa(x*4000000 + y), nil
			}
		}
	}

	return "missed", nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
