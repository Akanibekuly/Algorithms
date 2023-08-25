package interleavingstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsInterleave(t *testing.T) {
	cases := []struct {
		name, s1, s2, s3 string
		res              bool
	}{
		// {"empty strings", "", "", "", true},
		{"ok", "aabcc", "dbbca", "aadbbcbcac", true},
		// {"false", "aabcc", "dbbca", "aadbbbaccc", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, isInterleave(c.s1, c.s2, c.s3))
		})
	}
}
