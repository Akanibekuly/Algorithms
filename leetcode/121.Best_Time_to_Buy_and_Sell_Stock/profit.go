package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	var res int
	min := prices[0]
	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > res {
			res = prices[i] - min
		}
	}

	return res
}
