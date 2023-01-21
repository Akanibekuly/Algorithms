package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string

	var solve func(str string, tmp []string)
	solve = func(str string, tmp []string) {
		if len(str) == 0 {
			if len(tmp) == 4 {
				res = append(res, strings.Join(tmp, "."))
			}
			return
		}
		if len(tmp) >= 4 {
			return
		}

		for i := 1; i <= len(str) && i <= 3; i++ {
			snew, n := str[0:i], toInt(str[0:i])
			if snew != strconv.Itoa(n) {
				return
			}
			if n >= 0 && n < 256 {
				tmp = append(tmp, snew)
				solve(str[i:], tmp)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	solve(s, []string{})

	return res
}

func toInt(s string) int {
	var n int
	for _, d := range s {
		n *= 10
		n += int(d - '0')
	}
	return n
}
