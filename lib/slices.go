package lib

func Reverse[T any](slice []T) []T {
	ret := make([]T, len(slice))
	for i, val := range slice {
		ret[len(slice)-i-1] = val
	}
	return ret
}

func Column[T any](mat [][]T, col int) []T {
	row := make([]T, len(mat))
	for j := 0; j < len(mat); j++ {
		row[j] = mat[j][col]
	}
	return row
}

func Transpose[T any](mat [][]T) [][]T {
	ret := make([][]T, len(mat[0]))
	for i := range ret {
		ret[i] = make([]T, len(mat))
		for j := range ret[i] {
			ret[i][j] = mat[j][i]
		}
	}
	return ret
}
