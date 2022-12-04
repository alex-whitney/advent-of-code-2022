package lib

import (
	"strconv"
	"strings"
)

func ParseIntegerSlice(row string, delimiter string) ([]int, error) {
	in := strings.Split(row, delimiter)

	out := make([]int, len(in))
	var err error
	for i, val := range in {
		out[i], err = strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}
