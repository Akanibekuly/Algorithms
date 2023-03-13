package pascal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalsTriangle(t *testing.T) {
	cases := []struct {
		numRows int
		res     [][]int
	}{
		{1, [][]int{{1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.res, generate(c.numRows))
		})
	}
}
