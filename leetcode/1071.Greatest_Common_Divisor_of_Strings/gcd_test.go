package gcd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	cases := []struct {
		s1, s2, res string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, gcdOfStrings(c.s1, c.s2))
		})
	}
}
