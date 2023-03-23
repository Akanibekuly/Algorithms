package Number_of_Operations_to_Make_Network_Connected

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeConnected(t *testing.T) {
	cases := []struct {
		name        string
		n           int
		connections [][]int
		res         int
	}{
		{
			name:        "2 components",
			n:           4,
			connections: [][]int{{0, 1}, {0, 2}, {1, 2}},
			res:         1,
		},
		{
			name:        "3 components",
			n:           6,
			connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}},
			res:         2,
		},
		{
			name:        "not enough edges",
			n:           6,
			connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}},
			res:         -1,
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert.Equal(tt, c.res, makeConnected(c.n, c.connections), fmt.Sprintf("case %d", i+1))
		})
	}
}
