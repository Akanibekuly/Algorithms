package twosum

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i := range nums {
		if idx, ok := dict[target-nums[i]]; ok {
			return []int{idx, i}
		}
		dict[nums[i]] = i
	}

	return nil
}
