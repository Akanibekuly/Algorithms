package countsubarrayswithfixedbounds

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var count int64
	lastMin, lastMax, leftBound := -1, -1, -1

	for i := range nums {
		if nums[i] >= minK && nums[i] <= maxK {
			if nums[i] == minK {
				lastMin = i
			}
			if nums[i] == maxK {
				lastMax = i
			}
			count += int64(max(0, min(lastMin, lastMax)-leftBound))
		} else {
			leftBound, lastMin, lastMax = i, -1, -1
		}
	}
	return count
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
