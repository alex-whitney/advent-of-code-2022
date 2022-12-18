package main

import (
	"fmt"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

func hash(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func parseHash(s string) ([]int, error) {
	var x, y, z int
	_, err := fmt.Sscanf(s, "%d,%d,%d", &x, &y, &z)
	if err != nil {
		return nil, err
	}
	return []int{x, y, z}, nil
}

type Today struct {
	x     []int
	y     []int
	z     []int
	cubes map[string]bool
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.x = make([]int, len(raw))
	d.y = make([]int, len(raw))
	d.z = make([]int, len(raw))
	d.cubes = make(map[string]bool)
	for i, row := range raw {
		vals, err := parseHash(row)
		if err != nil {
			return err
		}
		d.x[i] = vals[0]
		d.y[i] = vals[1]
		d.z[i] = vals[2]
		d.cubes[row] = true
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	sides := 0

	for cube := 0; cube < len(d.x); cube++ {
		for _, side := range [][]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}} {
			sx := d.x[cube] + side[0]
			sy := d.y[cube] + side[1]
			sz := d.z[cube] + side[2]

			if !d.cubes[hash(sx, sy, sz)] {
				sides += 1
			}
		}
	}

	return strconv.Itoa(sides), nil
}

func (d *Today) Part2() (string, error) {
	outside := map[string]bool{"0,0,0": true}
	counter := 0

	next := []string{"0,0,0"}
	for len(next) > 0 {
		n := next[0]
		next = next[1:]

		nc, err := parseHash(n)
		if err != nil {
			return "", err
		}
		for _, side := range [][]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}} {
			sx := nc[0] + side[0]
			sy := nc[1] + side[1]
			sz := nc[2] + side[2]

			h := hash(sx, sy, sz)

			if sx < -1 || sy < -1 || sz < -1 || sx > 50 || sy > 50 || sz > 50 {
				continue
			}

			if outside[h] {
				continue
			}

			if !d.cubes[h] {
				next = append(next, h)
				outside[h] = true
			} else {
				counter += 1
			}
		}
	}

	return strconv.Itoa(counter), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
