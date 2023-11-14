package uniquelength3palindromicsubsequences

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountPalindromicSubsequence(t *testing.T) {
	cases := []struct {
		s   string
		res int
	}{
		{"aabca", 3},
		{"adc", 0},
		{"bbcbaba", 4},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			require.Equal(t, c.res, countPalindromicSubsequence(c.s))
		})
	}
}
