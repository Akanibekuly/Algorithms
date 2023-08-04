package permutations

func permute(nums []int) [][]int {
	result := make([][]int, 0, factor(len(nums)))

	permuteElem(&result, []int{}, nums)
	return result
}

func permuteElem(result *[][]int, current []int, left []int) {
	if len(left) == 0 {
		*result = append(*result, current)
		return
	}

	for i := 0; i < len(left); i++ {
		newLeft := append(append([]int{}, left[:i]...), left[i+1:]...)
		newCurrent := append(current, left[i])

		permuteElem(result, newCurrent, newLeft)
	}
}

func factor(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}
