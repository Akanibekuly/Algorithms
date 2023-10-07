package buildarraywhereyoucanfindthemaximumexactlykcomparisons

var mod int = 1e9 + 7

func numOfArrays(n int, m int, k int) int {
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for l := range memo[i][j] {
				memo[i][j][l] = -1
			}
		}
	}

	var dp func(i, maxSoFar, remain int) int
	dp = func(i, maxSoFar, remain int) int {
		if i == n {
			if remain == 0 {
				return 1
			}

			return 0
		}

		if remain < 0 {
			return 0
		}

		if memo[i][maxSoFar][remain] != -1 {
			return memo[i][maxSoFar][remain]
		}

		var ans int
		for num := 1; num <= maxSoFar; num++ {
			ans = (ans + dp(i+1, maxSoFar, remain)) % mod
		}

		for num := maxSoFar + 1; num <= m; num++ {
			ans = (ans + dp(i+1, num, remain-1)) % mod
		}

		memo[i][maxSoFar][remain] = ans
		return ans
	}

	return dp(0, 0, k)
}
