package palindromepartitioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromePartioning(t *testing.T) {
	cases := []struct {
		name        string
		s           string
		expectedRes [][]string
	}{
		{
			name:        "first case",
			s:           "aab",
			expectedRes: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name:        "all diff",
			s:           "abcd",
			expectedRes: [][]string{{"a", "b", "c", "d"}},
		},
		{
			name:        "one letter",
			s:           "a",
			expectedRes: [][]string{{"a"}},
		},
	}
	for i, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert.Equal(tt,
				c.expectedRes, partition(c.s),
				"case %d", i,
			)
		})

	}
}

// func TestIsPolindrome(t *testing.T) {
// 	cases := []struct {
// 		s     string
// 		isPol bool
// 	}{
// 		{s: "abvba", isPol: true},
// 	}
// 	for i, c := range cases {
// 		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
// 			assert.Equal(tt, isPolindrome(c.s), c.isPol)
// 		})
// 	}
// }
