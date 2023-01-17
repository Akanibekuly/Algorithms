package leetcode

func minFlipsMonoIncr(s string) int {
	total := 0
	for i := range s {
		if s[i] == '1' {
			total++
		}
	}

	if total == 0 || total == len(s) {
		return 0
	}

	count, min, ln := 0, total, len(s)
	if min > ln-total {
		min = ln - total
	}
	//fmt.Println(min)

	for i := range s {
		if s[i] == '1' {
			count++
		}

		tmp := ln + 2*count - (i + 1 + total)
		//fmt.Println(i, count, tmp)
		if tmp <= min {
			min = tmp
		}

		if min == 0 {
			return 0
		}
	}

	return min
}
