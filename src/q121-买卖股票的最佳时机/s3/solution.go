package main

import "fmt"

func main() {
	s := []int{6, 8, 30, 2, 30, 7, 8, 7, 7}
	fmt.Println(maxProfit(s))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([]int, n)
	minPrice := prices[0]

	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
