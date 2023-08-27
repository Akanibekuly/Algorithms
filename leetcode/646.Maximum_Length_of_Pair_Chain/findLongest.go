package maximumlengthofpairchain

import "sort"

func findLongestChainGreedy(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	ans, tail := 0, -1001
	for _, p := range pairs {
		if p[0] > tail {
			ans++
			tail = p[0]
		}
	}

	return ans
}

func findLongestChainDP(pairs [][]int) int {
	n, dp := len(pairs), make([]int, len(pairs))

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})

	for i := range dp {
		dp[i] = 1
	}

	var ans int
	for i := n - 1; i > 0; i-- {
		for j := i + 1; j < n; j++ {
			if pairs[j][1] < pairs[j][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			ans = max(ans, dp[i])
		}
	}
	return ans
}

func findLongestChainRecursive(pairs [][]int) int {
	n, memo := len(pairs), make([]int, len(pairs))

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})

	var recursive func(i int) int
	recursive = func(i int) int {
		if memo[i] != 0 {
			return memo[i]
		}

		memo[i] = 1
		for j := i + 1; j < n; j++ {
			if pairs[i][1] < pairs[j][0] {
				memo[i] = max(memo[i], 1+recursive(j))
			}
		}

		return memo[i]
	}

	var ans int
	for i := 0; i < n; i++ {
		ind := i
		ans = max(ans, recursive(ind))
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
