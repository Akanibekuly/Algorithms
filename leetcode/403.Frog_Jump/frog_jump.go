package frogjump

func canCross(stones []int) bool {
	l := len(stones)
	if stones[1] != 1 {
		return false
	}

	if l == 2 {
		return true
	}

	dest := l - 1
	var getLow int64 = (1 << 32) - 1
	mem := make(map[int64]bool)

	var solve func(int64) bool
	solve = func(prv int64) bool {
		if v, ok := mem[prv]; ok {
			return v
		}
		idx := int(prv >> 32)
		if idx == dest {
			return true
		}

		jp := int(prv & int64(getLow))
		currPos := stones[idx]
		for ngo := idx + 1; ngo < l; ngo++ {
			gap := stones[ngo] - currPos
			if gap+1 < jp {
				continue
			}

			if gap > jp+1 {
				mem[prv] = false
				return false
			}

			if solve(int64(ngo)<<32 | int64(gap)) {
				mem[prv] = true
				return true
			}
		}

		return false
	}

	return solve((1 << 32) + 1)
}

func canCrossDP(stones []int) bool {
	hash := make(map[int]int, len(stones))
	for i := range stones {
		hash[stones[i]] = i
	}

	dp := [2001][2001]bool{}
	dp[0][0] = true

	for index := 0; index < len(stones); index++ {
		for prevJump := 0; prevJump <= len(stones); prevJump++ {
			if dp[index][prevJump] {
				if _, ok := hash[stones[index]+prevJump]; ok {
					dp[hash[stones[index]+prevJump]][prevJump] = true
				}
				if _, ok := hash[stones[index]+prevJump+1]; ok {
					dp[hash[stones[index]+prevJump+1]][prevJump+1] = true
				}
				if _, ok := hash[stones[index]+prevJump-1]; ok {
					dp[hash[stones[index]+prevJump-1]][prevJump-1] = true
				}
			}
		}
	}

	for i := 0; i <= len(stones); i++ {
		if dp[len(stones)-1][i] {
			return true
		}
	}

	return false
}
