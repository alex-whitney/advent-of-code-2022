package main

import (
	"fmt"
	"strconv"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type Shape []lib.Point[int]

type Rock struct {
	Shape    Shape
	Position lib.Point[int]
}

func (r *Rock) MoveLeft() Rock {
	return Rock{
		Shape:    r.Shape,
		Position: lib.NewPoint(r.Position.X-1, r.Position.Y),
	}
}

func (r *Rock) MoveRight() Rock {
	return Rock{
		Shape:    r.Shape,
		Position: lib.NewPoint(r.Position.X+1, r.Position.Y),
	}
}

func (r *Rock) MoveDown() Rock {
	return Rock{
		Shape:    r.Shape,
		Position: lib.NewPoint(r.Position.X, r.Position.Y-1),
	}
}

func (r *Rock) Top() int {
	return r.Position.Y
}

func (r *Rock) Bottom() int {
	dy := 0
	for _, p := range r.Shape {
		if p.Y < dy {
			dy = p.Y
		}
	}
	return r.Position.Y + dy
}

func (r *Rock) Left() int {
	return r.Position.X
}

func (r *Rock) Right() int {
	dx := 0
	for _, p := range r.Shape {
		if p.X > dx {
			dx = p.X
		}
	}
	return r.Position.X + dx
}

func (r *Rock) Overlaps(r2 Rock) bool {
	for _, p1 := range r.Shape {
		for _, p2 := range r2.Shape {
			if p1.X+r.Position.X == r2.Position.X+p2.X &&
				p1.Y+r.Position.Y == r2.Position.Y+p2.Y {
				return true
			}
		}
	}
	return false
}

func (r *Rock) OverlapsAny(rocks []Rock) bool {
	for _, other := range rocks {
		if r.Overlaps(other) {
			return true
		}
	}
	return false
}

type Today struct {
	shapes      []Shape
	shapeHeight []int
	instr       string
}

func (d *Today) Init(input string) error {
	d.shapes = make([]Shape, 5)
	d.shapes[0] = Shape{
		lib.NewPoint(0, 0),
		lib.NewPoint(1, 0),
		lib.NewPoint(2, 0),
		lib.NewPoint(3, 0),
	}
	d.shapes[1] = Shape{
		lib.NewPoint(1, 0),
		lib.NewPoint(0, -1),
		lib.NewPoint(1, -1),
		lib.NewPoint(2, -1),
		lib.NewPoint(1, -2),
	}
	d.shapes[2] = Shape{
		lib.NewPoint(2, 0),
		lib.NewPoint(2, -1),
		lib.NewPoint(0, -2),
		lib.NewPoint(1, -2),
		lib.NewPoint(2, -2),
	}
	d.shapes[3] = Shape{
		lib.NewPoint(0, 0),
		lib.NewPoint(0, -1),
		lib.NewPoint(0, -2),
		lib.NewPoint(0, -3),
	}
	d.shapes[4] = Shape{
		lib.NewPoint(0, 0),
		lib.NewPoint(1, 0),
		lib.NewPoint(0, -1),
		lib.NewPoint(1, -1),
	}
	d.shapeHeight = []int{
		1, 3, 3, 4, 2,
	}

	raw, err := lib.ReadFile(input)
	if err != nil {
		return err
	}
	d.instr = raw

	return nil
}

func (d *Today) Part1() (string, error) {
	width := 7
	highest := 0

	instrCounter := 0
	blockCounter := 1
	fallingBlock := Rock{
		Shape:    d.shapes[0],
		Position: lib.NewPoint(2, 3),
	}

	staticRocks := []Rock{}

	for blockCounter < 2023 {
		instr := d.instr[instrCounter%len(d.instr)]
		instrCounter += 1
		if instr == '>' {
			tryMove := fallingBlock.MoveRight()
			if tryMove.Right() < width && !tryMove.OverlapsAny(staticRocks) {
				fallingBlock = tryMove
			}
		} else {
			tryMove := fallingBlock.MoveLeft()
			if tryMove.Left() >= 0 && !tryMove.OverlapsAny(staticRocks) {
				fallingBlock = tryMove
			}
		}

		stopped := false
		tryMove := fallingBlock.MoveDown()
		if tryMove.Bottom() < 0 || tryMove.OverlapsAny(staticRocks) {
			stopped = true
		}

		if stopped {
			staticRocks = append(staticRocks, fallingBlock)
			if fallingBlock.Top() > highest {
				highest = fallingBlock.Top()
			}

			shapeIndex := blockCounter % 5
			fallingBlock = Rock{
				Shape:    d.shapes[shapeIndex],
				Position: lib.NewPoint(2, highest+3+d.shapeHeight[shapeIndex]),
			}

			blockCounter += 1

		} else {
			fallingBlock = tryMove
		}
	}

	return strconv.Itoa(highest + 1), nil
}

func (d *Today) Part2() (string, error) {
	width := 7
	highest := 0

	instrCounter := 0
	blockCounter := 1
	fallingBlock := Rock{
		Shape:    d.shapes[0],
		Position: lib.NewPoint(2, 3),
	}

	staticRocks := []Rock{}

	last := 0
	cycles := []int{}
	blocks := []int{}
	heights := []int{}
	for len(cycles) < 10 {
		if instrCounter%len(d.instr) == 0 {
			cycles = append(cycles, instrCounter/len(d.instr))
			blocks = append(blocks, blockCounter)
			heights = append(heights, highest-last)
			last = highest
		}

		instr := d.instr[instrCounter%len(d.instr)]
		instrCounter += 1
		if instr == '>' {
			tryMove := fallingBlock.MoveRight()
			if tryMove.Right() < width && !tryMove.OverlapsAny(staticRocks) {
				fallingBlock = tryMove
			}
		} else {
			tryMove := fallingBlock.MoveLeft()
			if tryMove.Left() >= 0 && !tryMove.OverlapsAny(staticRocks) {
				fallingBlock = tryMove
			}
		}

		stopped := false
		tryMove := fallingBlock.MoveDown()
		if tryMove.Bottom() < 0 || tryMove.OverlapsAny(staticRocks) {
			stopped = true
		}

		if stopped {
			staticRocks = append(staticRocks, fallingBlock)
			if fallingBlock.Top() > highest {
				highest = fallingBlock.Top()
			}

			shapeIndex := blockCounter % 5
			fallingBlock = Rock{
				Shape:    d.shapes[shapeIndex],
				Position: lib.NewPoint(2, highest+3+d.shapeHeight[shapeIndex]),
			}

			if blockCounter == 4501 {
				fmt.Printf("Special block: %d\n", highest)
			}

			blockCounter += 1
		} else {
			fallingBlock = tryMove
		}
	}

	fmt.Printf("%+v\n%+v\n%+v\n", cycles, blocks, heights)

	for i := 1; i < len(blocks); i++ {
		fmt.Printf("%d: %d - %d\n", blocks[i], blocks[i]-blocks[i-1], heights[i])
	}

	// Instead of figuring out how to programmatically do this, I'm just going to solve this by hand..
	//
	// Output from puzzle:
	/*
		1748: 1747 - 2772
		3493: 1745 - 2752
		5238: 1745 - 2752
		6983: 1745 - 2752
		8728: 1745 - 2752
		10473: 1745 - 2752
		12218: 1745 - 2752
		13963: 1745 - 2752
		15708: 1745 - 2752
	*/
	// 1000000000000 - 1747 = 999,999,998,253
	//   + 2772
	// 999,999,998,253 / 1745 = 573,065,901
	//	 + 2752*573065901
	// 999,999,998,253 % 1745 = 1,008
	//
	// Just going to run again and spit out the 1,008th block after 3493
	//   7115-2772-2752=1591

	//  === 1577077363915

	return "manual!", nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
