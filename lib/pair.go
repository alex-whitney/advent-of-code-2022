package lib

type Pair[T any, S any] struct {
	Left  T
	Right S
}

func NewPair[T any, S any](left T, right S) Pair[T, S] {
	return Pair[T, S]{
		Left:  left,
		Right: right,
	}
}
