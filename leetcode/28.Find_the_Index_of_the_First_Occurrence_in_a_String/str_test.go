package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	cases := []struct {
		haystack, needle string
		res              int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, strStr(c.haystack, c.needle))
		})
	}
}
