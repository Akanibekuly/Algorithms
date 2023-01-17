package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlip(t *testing.T) {
	cases := []struct {
		name        string
		s           string
		expectedRes int
	}{
		// {
		// 	name:        "1",
		// 	s:           "00110",
		// 	expectedRes: 1,
		// },
		// {
		// 	name:        "2",
		// 	s:           "010110",
		// 	expectedRes: 2,
		// },
		// {
		// 	name:        "3",
		// 	s:           "00011000",
		// 	expectedRes: 2,
		// },
		{
			name:        "4",
			s:           "11011",
			expectedRes: 1,
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert.Equal(tt, c.expectedRes, minFlipsMonoIncr(c.s), "case %d", i)
		})
	}
}
