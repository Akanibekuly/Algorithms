package reshapethematrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixRashape(t *testing.T) {
	cases := []struct {
		mat, res [][]int
		r, c     int
	}{
		{
			mat: [][]int{{1, 2}, {3, 4}},
			res: [][]int{{1, 2, 3, 4}},
			r:   1,
			c:   4,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.res, matrixReshape(c.mat, c.r, c.c))
		})
	}
}
