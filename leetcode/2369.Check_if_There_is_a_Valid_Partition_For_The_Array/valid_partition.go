package valid

func validPartition(nums []int) bool {
	dp := make([]bool, 3)
	dp[0] = true

	for i := range nums {
		ans := false

		if i > 0 && dp[(i-1)%3] {
			ans = ans || nums[i] == nums[i-1]
		}

		if i > 1 && dp[(i-2)%3] {
			ans = ans || (nums[i] == nums[i-1] && nums[i-1] == nums[i-2])
			if !ans {
				ans = ans || (nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1)
			}
		}

		dp[(i+1)%3] = ans
	}

	return dp[len(nums)%3]
}

func validPartitionBottomUpDP(nums []int) bool {
	dp := make([]bool, len(nums)+1)
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		if i > 0 && dp[i-1] {
			dp[i+1] = dp[i+1] || (nums[i] == nums[i-1])
		}

		if i > 1 && dp[i-2] {
			dp[i+1] = dp[i+1] || (nums[i] == nums[i-1] && nums[i-1] == nums[i-2])
			dp[i+1] = dp[i+1] || (nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1)
		}

	}

	return dp[len(nums)]
}

func validPartitionReccursive(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	if len(nums) == 1 {
		return false
	}
	found := false
	if len(nums) >= 2 && nums[0] == nums[1] {
		found = found || validPartitionReccursive(nums[2:])
	}
	if len(nums) >= 3 && nums[0] == nums[1] && nums[0] == nums[2] {
		found = found || validPartitionReccursive(nums[3:])
	}
	if len(nums) >= 3 && nums[0]+1 == nums[1] && nums[1]+1 == nums[2] {
		found = found || validPartitionReccursive(nums[3:])
	}

	return found
}
