package repeatedsubstringpattern

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedSubstr(t *testing.T) {
	cases := []struct {
		s   string
		res bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, repeatedSubstringPattern(c.s))
		})
	}
}
