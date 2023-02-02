package verifyinganaliendictionary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienOrder(t *testing.T) {
	cases := []struct {
		words []string
		order string
		res   bool
	}{
		{
			words: []string{"hello", "leetcode"},
			order: "hlabcdefgijkmnopqrstuvwxyz",
			res:   true,
		},
		{
			words: []string{"word", "world", "row"},
			order: "worldabcefghijkmnpqstuvxyz",
			res:   false,
		},
		{
			words: []string{"apple", "app"},
			order: "abcdefghijklmnopqrstuvwxyz",
			res:   false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case: %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, isAlienSorted(c.words, c.order))
		})
	}
}
