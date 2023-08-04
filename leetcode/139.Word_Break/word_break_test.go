package wordbreak

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordBreak(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		wordDict []string
		result   bool
	}{
		{"leetcode", "leetcode", []string{"leet", "code"}, true},
		{"applepenapple", "applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.result, wordBreak(c.s, c.wordDict))
		})
	}
}
