package repeatedsubstringpattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s); i++ {
		if len(s)%i == 0 {
			if strings.Repeat(s[:i], len(s)/i) == s {
				return true
			}
		}
	}
	return false
}
