package maximumsubarray

func maxSubArray(nums []int) int {
	var s, max int
	for i := range nums {
		s += nums[i]

		if s < 0 {
			s = 0
		} else if s > max {
			max = s
		}
	}

	return max
}
