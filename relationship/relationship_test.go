package main

import (
	"fmt"
	"testing"
)

var Answers = []int{1, 1, 3, 13, 75, 541, 4683, 47293, 545835, 7087261, 102247563}

// , 75, 541, 4683, 47293, 545835, 7087261, 102247563

func TestRelationship(t *testing.T) {
	for i, v := range Answers {
		res := Relationship(i)
		if res != v {
			t.Errorf("ERROR expected %d got %d", v, res)
		}
		fmt.Printf("DEBUG res: %d got %d\n", v, res)
	}
}
