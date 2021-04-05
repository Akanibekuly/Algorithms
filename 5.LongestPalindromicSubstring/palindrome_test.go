package main

import (
	"fmt"
	"testing"
)

var Answers = map[string]string{
	"a":     "a",
	"babad": "bab",
	"cbbd":  "bb",
	"ac":    "a",
}

func TestLongestPalindrome(t *testing.T) {
	for k, v := range Answers {
		if res := LongestPalindrome(k); res != v {
			t.Errorf("Error expected %s got %s\n", v, res)
		}
		fmt.Printf("DEBUG inout=%s res=%s\n", k, v)
	}

}
