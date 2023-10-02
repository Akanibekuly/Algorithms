package decodedstringatindex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeStringAtIndex(t *testing.T) {
	cases := []struct {
		name   string
		s, res string
		k      int
	}{
		{"leetcode", "leet2code3", "o", 10},
		{"ha22", "ha22", "h", 5},
		{"a2345678999999999999999", "a2345678999999999999999", "a", 1},
		{"abc", "abc", "b", 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, decodeAtIndex(c.s, c.k))
		})
	}
}
