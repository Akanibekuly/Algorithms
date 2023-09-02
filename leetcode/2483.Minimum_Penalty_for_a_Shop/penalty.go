package minimumpenaltyforashop

func bestClosingTime(customers string) int {
	yt := 0
	for _, r := range customers {
		switch r {
		case 'Y':
			yt++
		}
	}

	min, mini := yt, -1
	for i := 0; i < len(customers); i++ { // check [0:n-1]

		switch customers[i] {
		case 'Y':
			yt--
		case 'N':
			yt++
		}

		if yt < min {
			min = yt
			mini = i
		}
	}

	return mini + 1
}
