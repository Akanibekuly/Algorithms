package maxdotproductoftwosubsequences

import "math"

func maxDotProduct(nums1 []int, nums2 []int) int {
	firstMax, secondMax := math.MinInt32, math.MinInt32
	firstMin, secondMin := math.MaxInt32, math.MaxInt32
	for _, num := range nums1 {
		firstMax = max(firstMax, num)
		firstMin = min(firstMin, num)
	}

	for _, num := range nums2 {
		secondMax = max(secondMax, num)
		secondMin = min(secondMin, num)
	}

	if firstMax < 0 && secondMin > 0 {
		return firstMax * secondMin
	}

	if firstMin > 0 && secondMax < 0 {
		return firstMin * secondMax
	}

	memo := make([][]int, len(nums1))
	for i := range memo {
		memo[i] = make([]int, len(nums2))
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == len(nums1) || j == len(nums2) {
			return 0
		}

		if memo[i][j] != 0 {
			return memo[i][j]
		}

		use := nums1[i]*nums2[j] + dp(i+1, j+1)
		memo[i][j] = max(use, max(dp(i+1, j), dp(i, j+1)))

		return memo[i][j]
	}

	return dp(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
