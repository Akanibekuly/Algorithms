package minimumpenaltyforashop

import "fmt"

func bestClosingTime(customers string) int {
	yt, nt := 0, 0
	for _, r := range customers {
		switch r {
		case 'Y':
			yt++
		case 'N':
			nt++
		}
	}

	if yt == len(customers) {
		return len(customers)
	}

	if nt == len(customers) {
		return 0
	}

	min, mini := len(customers)+1, -1
	n := 0
	for i := 0; i < len(customers); i++ { // check [0:n-1]
		pen := n + yt
		if pen < min {
			min = pen
			mini = i
		}
		fmt.Println(i, pen)
		switch customers[i] {
		case 'Y':
			yt--
		case 'N':
			n++
		}
	}

	// check [n]
	if n < min {
		mini = len(customers)
		fmt.Println(mini, n)
	}

	return mini
}
