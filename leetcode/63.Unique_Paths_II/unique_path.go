package uniquepathsii

func uniquePathsWithObstacles(dp [][]int) int {
	if dp[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := range dp {
		for j := range dp[i] {
			if i == 0 && j == 0 {
				continue
			}

			if dp[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
