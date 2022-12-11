package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	tmat := Transpose(mat)
	assert.Equal(t, 3, len(tmat))
	assert.Equal(t, 2, len(tmat[0]))

	assert.Equal(t, 1, tmat[0][0])
	assert.Equal(t, 3, tmat[2][0])
	assert.Equal(t, 6, tmat[2][1])
}
