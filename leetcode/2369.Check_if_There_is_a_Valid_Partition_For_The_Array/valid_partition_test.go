package valid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidPartition(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{"valid", []int{4, 4, 4, 5, 6}, true},
		{"not valid", []int{1, 1, 1, 2}, false},
		{"should be true", []int{803201, 803201, 803201, 803201, 803202, 803203}, true},
		{"two nums", []int{5, 5}, true},
		{"one number", []int{2}, false},
	}

	for _, c := range cases {
		require.Equal(t, c.res, validPartition(c.nums))
	}
}
