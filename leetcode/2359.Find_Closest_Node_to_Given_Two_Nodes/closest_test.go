package closest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosest(t *testing.T) {
	cases := []struct {
		edges             []int
		node1, node2, res int
	}{
		{[]int{2, 2, 3, -1}, 0, 1, 2},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, closestMeetingNode(c.edges, c.node1, c.node2))
		})
	}
}
