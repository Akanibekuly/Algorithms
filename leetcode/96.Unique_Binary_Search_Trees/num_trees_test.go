package uniquebinarysearchtrees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumTrees(t *testing.T) {
	cases := []struct {
		name   string
		n, res int
	}{
		{"3", 3, 5},
		{"1", 1, 1},
	}

	for _, c := range cases {
		require.Equal(t, c.res, numTrees(c.n))
	}
}
