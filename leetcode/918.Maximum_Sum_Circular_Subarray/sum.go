package leetcode

import (
	"math"
)

// elegant O(n) Kadane's algorithm
func maxSubarraySumCircular(nums []int) int {
	var total, currMax, currMin int
	sumMax, sumMin := nums[0], nums[0]
	for _, item := range nums {
		currMax = max(currMax+item, item)
		sumMax = max(sumMax, currMax)

		currMin = min(currMin+item, item)
		sumMin = min(sumMin, currMin)

		total += item
	}

	if sumMax > 0 {
		return max(sumMax, total-sumMin)
	}

	return sumMax
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

// brute force O(n^2)
func maxSubarraySumCircularBrute(nums []int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}

	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}

			if !(i == 0 && j == len(nums)-1) && total-sum > max {
				max = total - sum
			}
		}
	}

	return max
}
