package minimumreplacementstosortthearray

func minimumReplacement(nums []int) int64 {
	var count int64
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			continue
		}

		numElements := (nums[i+1] + nums[i] - 1) / nums[i+1]
		count += int64(numElements) - 1
		nums[i] = nums[i] / numElements
	}
	return count
}
