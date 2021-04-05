package main

import "fmt"

func main() {
	fmt.Println(LongestPalindrome("a"))
}

func LongestPalindrome(s string) string {
	max := ""
	for i := range s {
		// fmt.Println(i)
		for j := i; j < len(s); j++ {
			t := s[i : j+1]
			if isPolindrome(t) && len(max) < len(t) {
				max = t
				// fmt.Println(t, i, j)
			}
		}
	}
	return max
}

func isPolindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
