package lib

func Reverse[T interface{}](slice []T) []T {
	ret := make([]T, len(slice))
	for i, val := range slice {
		ret[len(slice)-i-1] = val
	}
	return ret
}

func Column[T interface{}](mat [][]T, col int) []T {
	row := make([]T, len(mat))
	for j := 0; j < len(mat); j++ {
		row[j] = mat[j][col]
	}
	return row
}
