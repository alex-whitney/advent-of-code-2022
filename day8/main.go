package main

import (
	"strconv"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Today struct {
	forest [][]int
}

func (d *Today) Init(input string) error {
	lines, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.forest = make([][]int, len(lines))
	for i, line := range lines {
		d.forest[i] = make([]int, len(line))
		for j, c := range line {
			d.forest[i][j] = int(c) - 48

		}
	}

	return nil
}

func coord(r int, c int) string {
	return strconv.Itoa(r) + "_" + strconv.Itoa(c)
}

func (d *Today) Part1() (string, error) {
	visible := make(map[string]bool)
	for r, row := range d.forest {
		visible[coord(r, 0)] = true
		last := row[0]
		for c := 1; c < len(row); c++ {
			if last < row[c] {
				visible[coord(r, c)] = true
				last = row[c]
			}
		}

		visible[coord(r, len(row)-1)] = true
		last = row[len(row)-1]
		for c := len(row) - 1; c >= 0; c-- {
			if last < row[c] {
				visible[coord(r, c)] = true
				last = row[c]
			}
		}
	}

	// puzzle input is square
	for c := range d.forest {
		col := lib.Column(d.forest, c)

		visible[coord(0, c)] = true
		last := col[0]
		for r := 1; r < len(col); r++ {
			if last < col[r] {
				visible[coord(r, c)] = true
				last = col[r]
			}
		}

		visible[coord(len(col)-1, c)] = true
		last = col[len(col)-1]
		for r := len(col) - 1; r >= 0; r-- {
			if last < col[r] {
				visible[coord(r, c)] = true
				last = col[r]
			}
		}
	}

	return strconv.Itoa(len(visible)), nil
}

func (d *Today) score(r int, c int) int {
	base := d.forest[r][c]

	if r == 0 || c == 0 || r == len(d.forest) || c == len(d.forest) {
		return 0
	}

	col := lib.Column(d.forest, c)
	up := 0
	for i := r - 1; i >= 0; i-- {
		up++
		if col[i] >= base {
			break
		}
	}

	down := 0
	for i := r + 1; i < len(d.forest); i++ {
		down++
		if col[i] >= base {
			break
		}
	}

	row := d.forest[r]
	left := 0
	for i := c - 1; i >= 0; i-- {
		left++
		if row[i] >= base {
			break
		}
	}

	right := 0
	for i := c + 1; i < len(d.forest); i++ {
		right++
		if row[i] >= base {
			break
		}
	}

	return up * down * left * right
}

func (d *Today) Part2() (string, error) {
	max := 0
	for r := 1; r < len(d.forest)-1; r++ {
		for c := 1; c < len(d.forest)-1; c++ {
			score := d.score(r, c)
			if score > max {
				max = score
			}
		}
	}

	return strconv.Itoa(max), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
