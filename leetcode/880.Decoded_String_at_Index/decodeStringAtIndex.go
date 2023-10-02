package decodedstringatindex

import (
	"unicode"
)

func decodeAtIndex(s string, k int) string {

	var (
		num, lastIdx int
		tape         string
	)

	for i, r := range s {
		if unicode.IsDigit(r) {
			if tape == "" { // digits starts
				tape = s[lastIdx:i]
			}

			num = num*10 + int(r-'0')
			if len(tape)*num > k {
				return string(tape[(k-1)%len(tape)])
			}

		} else if tape != "" { // start sitrings
			if k < len(tape)*num {
				return string(tape[(k-1)%len(tape)])
			}

			k -= len(tape) * num
			lastIdx, tape, num = i, "", 0

		}
	}

	if num == 0 {
		tape = s[lastIdx:]
		return string(tape[(k-1)%len(tape)])
	}

	return ""
}
