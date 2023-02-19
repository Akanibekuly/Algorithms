package findthearrayconcatenationvalue

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	var ans int64
	for len(nums) > 0 {
		if len(nums) == 1 {
			ans += int64(nums[0])
			nums = nums[1:]
		} else {
			a, b := nums[0], nums[len(nums)-1]
			nums = nums[1 : len(nums)-1]
			tmp, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
			ans += int64(tmp)
		}
	}

	return ans
}
