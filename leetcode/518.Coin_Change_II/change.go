package coinchangeii

func change(amount int, coins []int) int {
	if len(coins) == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, v := range coins {
		for i := v; i <= amount; i++ {
			dp[i] += dp[i-v]
		}
	}

	return dp[amount]
}

func changeBottomtoTop(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= amount; j++ {
			if coins[i] > j {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = dp[i][j-coins[i]] + dp[i+1][j]
			}
		}
	}
	return dp[0][amount]
}

var Coins []int
var memo [][]int

func changeTopBottom(amount int, coins []int) int {
	Coins = coins
	memo = make([][]int, len(coins))
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return numberOfWays(0, amount)
}

func numberOfWays(i, amount int) int {
	if amount == 0 {
		return 1
	}

	if i == len(Coins) {
		return 0
	}

	if memo[i][amount] != -1 {
		return memo[i][amount]
	}

	if Coins[i] > amount {
		memo[i][amount] = numberOfWays(i+1, amount)
		return memo[i][amount]
	}

	memo[i][amount] = numberOfWays(i, amount-Coins[i]) + numberOfWays(i+1, amount)

	return memo[i][amount]
}
