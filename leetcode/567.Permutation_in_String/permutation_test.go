package permutationinstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	cases := []struct {
		s1, s2 string
		res    bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"xyz", "asdcayxzsdvc", true},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			assert.Equal(tt, c.res, checkInclusion(c.s1, c.s2))
		})
	}
}
