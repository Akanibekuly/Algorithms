package uniquepaths

import "fmt"

func uniquePaths(m int, n int) int {
	if m > n {
		return uniquePaths(n, m)
	}
	dp := make([]int, m)

	dp[0] = 1
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			dp[i] += dp[i-1]
		}
		fmt.Println(dp)
	}

	return dp[m-1]
}
