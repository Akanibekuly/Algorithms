package main

import "fmt"

func countGood(nums []int, k int) int64 {
	// fmt.Println(nums)
	var count int64
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		dict[nums[i]]++
		var last int
		for j := i + 1; j < len(nums); j++ {
			dict[nums[j]]++
			last += dict[nums[j]] - 1

			if last >= k {
				count += int64(len(nums) - j)
				fmt.Println(nums[i:j+1], last, j, len(nums)-j)
				break
			}
		}
	}

	return count
}

// func countGood(nums []int, k int) (count int64) {
// 	// fmt.Println(nums)
// 	var i, j, last int
// 	dict := make(map[int]int)
// 	for i < len(nums) {
// 		for ; j < len(nums); j++ {
// 			// if j == len(nums) {
// 			// 	return
// 			// }
// 			dict[nums[j]]++
// 			last += dict[nums[j]] - 1
// 			fmt.Println(dict, last)

// 			if last >= k {
// 				count += int64(len(nums) - j)
// 				fmt.Println(nums[i:j+1], last, j, len(nums)-j)
// 				break
// 			}
// 		}

// 		for ; i < len(nums); i++ {
// 			last -= dict[nums[i]] - 1
// 			dict[nums[i]]--
// 			if last < k {
// 				break
// 			}
// 			count++
// 		}

// 		if j == len(nums)-1 && last < k {
// 			break
// 		}
// 	}

// 	return
// }
