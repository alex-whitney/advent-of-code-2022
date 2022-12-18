package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Today struct {
	sensors   []lib.Point[int]
	beacons   []lib.Point[int]
	distances []int

	maxX int
	minX int

	part1Target int
	part2Size   int
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.sensors = make([]lib.Point[int], len(raw))
	d.beacons = make([]lib.Point[int], len(raw))
	d.distances = make([]int, len(raw))
	for i, line := range raw {
		var sx, sy, bx, by int
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil {
			return err
		}

		d.sensors[i] = lib.NewPoint(sx, sy)
		d.beacons[i] = lib.NewPoint(bx, by)
		d.distances[i] = d.sensors[i].ManhattanDistance(d.beacons[i])

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

	pt := lib.NewPoint(0, d.part1Target)
	for x := d.minX * 2; x <= d.maxX*2; x++ {
		inBounds := false
		pt.X = x
		for i, sensor := range d.sensors {
			if pt.X == d.beacons[i].X && pt.Y == d.beacons[i].Y {
				inBounds = false
				break
			}

			dist := pt.ManhattanDistance(sensor)
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
	pt := lib.NewPoint(0, 0)
	for x := 0; x <= d.part2Size; x++ {
		for y := 0; y <= d.part2Size; y++ {
			pt.X = x
			pt.Y = y

			valid := true
			for i, sensor := range d.sensors {
				dist := pt.ManhattanDistance(sensor)
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
