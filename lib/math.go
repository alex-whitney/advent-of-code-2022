package lib

type Number interface {
	int64 | float64 | int
}

func Sum[T Number](slice []T) T {
	var total T
	for _, val := range slice {
		total += val
	}
	return total
}
