package main

import (
	"fmt"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2022/lib"
	"golang.org/x/exp/slices"
)

func hashPoint(point lib.Pair[int, int]) string {
	return fmt.Sprintf("(%d,%d)", point.Left, point.Right)
}

type path struct {
	steps []lib.Pair[int, int]
}

func (p path) hash() string {
	hash := ""
	for _, step := range p.steps {
		hash += hashPoint(step)
	}
	return hash
}

func (p path) copy() path {
	return path{
		steps: slices.Clone(p.steps),
	}
}

type Today struct {
	grid  [][]string
	start lib.Pair[int, int]
	end   lib.Pair[int, int]
}

func (d *Today) Init(input string) error {
	var err error
	d.grid, err = lib.ReadDelimitedFile(input, "")
	if err != nil {
		return err
	}

	for r, row := range d.grid {
		for c, val := range row {
			if val == "S" {
				d.start = lib.NewPair(r, c)
			} else if val == "E" {
				d.end = lib.NewPair(r, c)
			}
		}
	}

	return nil
}

func isAtMostOneHigher(a string, b string) bool {
	if a == "S" {
		return true
	}
	if b == "E" {
		b = "z"
	}

	first := a[0]
	second := b[0]

	return (int(second) - int(first)) <= 1
}

func (d *Today) Part1() (string, error) {
	visited := map[string]int{}
	visited[hashPoint(d.start)] = 0

	paths := []path{
		{steps: []lib.Pair[int, int]{d.start}},
	}

	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]

		last := path.steps[len(path.steps)-1]

		for _, mod := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			next := lib.NewPair(last.Left+mod[0], last.Right+mod[1])
			hash := hashPoint(next)

			newPath := path.copy()
			newPath.steps = append(newPath.steps, next)

			if next.Left >= len(d.grid) || next.Right >= len(d.grid[0]) || next.Left < 0 || next.Right < 0 {
				continue
			}

			currentValue := d.grid[last.Left][last.Right]
			nextValue := d.grid[next.Left][next.Right]
			if !isAtMostOneHigher(currentValue, nextValue) {
				continue
			}

			if d.grid[next.Left][next.Right] == "E" {
				return strconv.Itoa(len(newPath.steps) - 1), nil
			}

			// just looking for the shortest path, so if the BFS already got to this point, we can
			// discard this one
			_, ok := visited[hash]
			if !ok {
				visited[hash] = len(newPath.steps)
				paths = append(paths, newPath)
			}
		}
	}

	return "Failed to find a solution", nil
}

func (d *Today) Part2() (string, error) {
	// same problem, but in reverse. Look for the first "a", starting from E

	visited := map[string]int{}
	visited[hashPoint(d.end)] = 0

	paths := []path{
		{steps: []lib.Pair[int, int]{d.end}},
	}

	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]

		last := path.steps[len(path.steps)-1]

		for _, mod := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			next := lib.NewPair(last.Left+mod[0], last.Right+mod[1])
			hash := hashPoint(next)

			newPath := path.copy()
			newPath.steps = append(newPath.steps, next)

			if next.Left >= len(d.grid) || next.Right >= len(d.grid[0]) || next.Left < 0 || next.Right < 0 {
				continue
			}

			currentValue := d.grid[last.Left][last.Right]
			nextValue := d.grid[next.Left][next.Right]
			// gotta flip these from part1
			if !isAtMostOneHigher(nextValue, currentValue) {
				continue
			}

			if d.grid[next.Left][next.Right] == "a" {
				return strconv.Itoa(len(newPath.steps) - 1), nil
			}

			// just looking for the shortest path, so if the BFS already got to this point, we can
			// discard this one
			_, ok := visited[hash]
			if !ok {
				visited[hash] = len(newPath.steps)
				paths = append(paths, newPath)
			}
		}
	}

	return "Failed to find a solution", nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
