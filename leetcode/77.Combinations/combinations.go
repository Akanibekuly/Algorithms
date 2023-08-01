package combinations

func combine(n int, k int) [][]int {
	res := make([][]int, 0, cnk(n, k))

	// 111000
	arr := make([]int, n)
	for i := 0; i < k; i++ {
		arr[i] = 1
	}

	for {
		// 110100 -> 124
		curr := make([]int, 0, k)
		for i, v := range arr {
			if v == 1 {
				curr = append(curr, i+1)
			}
		}

		res = append(res, curr)

		/// 111000 -> 110100
		if !switchToTheNextVal(arr) {
			break
		}
	}

	return res
}

func switchToTheNextVal(arr []int) bool {
	count := 0
	for i := len(arr) - 2; i >= 0; i-- {

		if arr[i] == 1 && arr[i+1] == 0 {
			arr[i], arr[i+1] = 0, 1
			for j := i + 2; j < len(arr); j++ {
				if count > 0 {
					arr[j] = 1
					count--
				} else {
					arr[j] = 0
				}
			}

			return true
		} else if arr[i+1] == 1 {
			count++
		}
	}

	return false
}

func cnk(n, k int) int {
	res := 1
	if n == k {
		return 1
	}

	if k > n/2 {
		return cnk(n, n-k)
	}

	for i := n; i > n-k; i-- {
		res *= i
	}

	for i := 1; i <= k; i++ {
		res /= i
	}
	return res
}
