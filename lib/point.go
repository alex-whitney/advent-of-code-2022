package lib

import (
	"fmt"
)

type Point[T Number] struct {
	X T
	Y T
}

func NewPoint[T Number](x T, y T) Point[T] {
	return Point[T]{
		X: x,
		Y: y,
	}
}

func (p *Point[T]) String() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

func (p *Point[T]) ManhattanDistance(p2 Point[T]) T {
	var distance T
	if p2.X > p.X {
		distance += p2.X - p.X
	} else {
		distance += p.X - p2.X
	}

	if p2.Y > p.Y {
		distance += p2.Y - p.Y
	} else {
		distance += p.Y - p2.Y
	}

	return distance
}
