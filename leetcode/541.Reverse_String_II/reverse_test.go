package reversestringii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStringII(t *testing.T) {
	cases := []struct {
		k      int
		s, res string
	}{
		{2, "abcdefg", "bacdfeg"},
		{2, "abcd", "bacd"},
		{8, "abcdefg", "gfedcba"},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, reverseStr(c.s, c.k))
			assert.Equal(tt, c.res, reverseStr2(c.s, c.k))
		})
	}
}
