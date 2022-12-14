package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Today struct {
	paths [][]lib.Pair[int, int]

	// right,down -> #/""/o
	grid map[string]string

	maxDepth int
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadDelimitedFile(input, " -> ")
	if err != nil {
		return err
	}

	d.paths = make([][]lib.Pair[int, int], len(raw))
	for i, points := range raw {
		d.paths[i] = make([]lib.Pair[int, int], len(points))
		for j, ptStr := range points {
			parts := strings.Split(ptStr, ",")
			a, err := strconv.Atoi(parts[0])
			if err != nil {
				return err
			}
			b, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
			d.paths[i][j] = lib.NewPair(a, b)

			if d.maxDepth < b {
				d.maxDepth = b
			}
		}
	}

	d.initializeGrid()

	return nil
}

func hashPoint(p lib.Pair[int, int]) string {
	return fmt.Sprintf("%d,%d", p.Left, p.Right)
}

func (d *Today) initializeGrid() {
	d.grid = map[string]string{}

	for _, path := range d.paths {
		for i := 1; i < len(path); i++ {
			lastPoint := path[i-1]
			point := path[i]

			var nextPoint lib.Pair[int, int]
			if point.Left == lastPoint.Left {
				if point.Right > lastPoint.Right {
					point, lastPoint = lastPoint, point
				}
				for r := point.Right; r <= lastPoint.Right; r++ {
					nextPoint = lib.NewPair(point.Left, r)
					d.grid[hashPoint(nextPoint)] = "#"
				}
			} else {
				if point.Left > lastPoint.Left {
					point, lastPoint = lastPoint, point
				}
				for left := point.Left; left <= lastPoint.Left; left++ {
					nextPoint = lib.NewPair(left, point.Right)
					d.grid[hashPoint(nextPoint)] = "#"
				}
			}
		}
	}
}

func (d *Today) Part1() (string, error) {
	var sand lib.Pair[int, int]
	sandCounter := 0
	for sand.Right < d.maxDepth {
		sand = lib.NewPair(500, 0)
		moved := true
		for moved && sand.Right < d.maxDepth {
			moved = false
			for _, perm := range [][]int{{0, 1}, {-1, 1}, {1, 1}} {
				nextSand := lib.NewPair(sand.Left+perm[0], sand.Right+perm[1])
				hash := hashPoint(nextSand)
				if d.grid[hash] == "" {
					sand = nextSand
					moved = true
					break
				}
			}
		}
		d.grid[hashPoint(sand)] = "o"
		sandCounter += 1
	}

	return strconv.Itoa(sandCounter - 1), nil
}

func (d *Today) Part2() (string, error) {
	d.initializeGrid()

	var sand lib.Pair[int, int]
	sandCounter := 0
	for {
		sand = lib.NewPair(500, 0)
		if d.grid[hashPoint(sand)] == "o" {
			return strconv.Itoa(sandCounter), nil
		}

		moved := true
		for moved {
			moved = false
			for _, perm := range [][]int{{0, 1}, {-1, 1}, {1, 1}} {
				nextSand := lib.NewPair(sand.Left+perm[0], sand.Right+perm[1])
				if nextSand.Right >= d.maxDepth+2 {
					// the floor
					continue
				}

				hash := hashPoint(nextSand)
				if d.grid[hash] == "" {
					sand = nextSand
					moved = true
					break
				}
			}
		}
		d.grid[hashPoint(sand)] = "o"
		sandCounter += 1
	}

	return "Something didn't work", nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
