package thekweakestrowsinamatrix

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	res := []int{}
	for j := 0; j < len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			if mat[i][j] == 0 && ((j == 0) || (mat[i][j-1] != 0)) {
				res = append(res, i)
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		if mat[i][len(mat[0])-1] == 1 {
			res = append(res, i)
		}
	}
	return res[:k]
}

func kWeakestRowsSorting(mat [][]int, k int) []int {
	countOnes := make([][2]int, len(mat))
	for i := range mat {
		countOnes[i][0] = i
		ones := 0
		for j := range mat[i] {
			if mat[i][j] == 0 {
				break
			}
			ones++
		}
		countOnes[i][1] = ones
	}

	sort.Slice(countOnes, func(i, j int) bool {
		if countOnes[i][1] == countOnes[j][1] {
			return countOnes[i][0] < countOnes[j][0]
		}
		return countOnes[i][1] < countOnes[j][1]
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, countOnes[i][0])
	}

	return res
}
