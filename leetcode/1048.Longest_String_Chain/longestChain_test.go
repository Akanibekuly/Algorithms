package longeststringchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestChain(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   int
	}{
		{"1 case", []string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
		{"2 case", []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.res, longestStrChain(c.words))
		})
	}
}
