package interleavingstring

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1+l2 != len(s3) {
		return false
	}

	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s3[i+j-1] == s1[i-1]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s3[i+j-1] == s2[j-1]
			}
		}
	}

	return dp[l1][l2]
}

func isInterleaveDP(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true
	l1, l2 := len(s1), len(s2)
	curr := make(map[string][]int)
	curr["0,0"] = []int{0, 0}

	// bfs
	for len(curr) > 0 {
		// fmt.Println(curr)
		next := make(map[string][]int)
		for _, coor := range curr {
			// fmt.Println(coor, s3[coor[0]+coor[1]:], s1[coor[0]:], s2[coor[1]:])
			// check -> s1
			if coor[0] < l1 && s3[coor[0]+coor[1]] == s1[coor[0]] {
				key := fmt.Sprintf("%d,%d", coor[0]+1, coor[1])
				if _, ok := curr[key]; !ok {
					dp[coor[0]+1][coor[1]] = true
					next[key] = []int{coor[0] + 1, coor[1]}
				}
			}

			// check -> s2
			if coor[1] < l2 && s3[coor[0]+coor[1]] == s2[coor[1]] {
				key := fmt.Sprintf("%d,%d", coor[0], coor[1]+1)
				if _, ok := curr[key]; !ok {
					dp[coor[0]][coor[1]+1] = true
					next[key] = []int{coor[0], coor[1] + 1}
				}
			}
		}

		curr = next
		// print(dp)
	}

	return dp[l1][l2]
}

func print(arr [][]bool) {
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("--------------------------------")
}
