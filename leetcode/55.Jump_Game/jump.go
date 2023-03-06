package jump

func canJump(nums []int) bool {
	fastest := 0
	for i := 0; i <= fastest; i++ {
		if i == len(nums)-1 {
			return true
		}
		if i+nums[i] > fastest {
			fastest = i + nums[i]
		}
		if fastest >= len(nums)-1 {
			return true
		}
	}
	return false
}
