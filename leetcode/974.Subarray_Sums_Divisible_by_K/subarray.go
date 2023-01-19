package subarraysumsdivisiblebyk

func subarraysDivByK(nums []int, k int) int {
	hash := make(map[int]int, k)
	hash[0] = 1
	var sum, count int
	for _, num := range nums {
		sum = (sum + num%k + k) % k
		count += hash[sum]
		hash[sum]++
	}

	return count
}
