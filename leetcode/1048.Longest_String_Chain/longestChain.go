package longeststringchain

import (
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	hash := make(map[string]int, len(words))
	ans := 0
	for _, w := range words {
		longest := 0
		for i := range w {
			sub := w[:i] + w[i+1:]
			longest = max(longest, hash[sub]+1)
		}

		hash[w] = longest
		ans = max(ans, longest)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
