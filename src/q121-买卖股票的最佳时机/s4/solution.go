package main

import "fmt"

func main() {
	s := []int{6,5,9}
	fmt.Println(maxProfit(s))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	minPrice := prices[0]
	maxprofit := 0
	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxprofit {
			maxprofit = prices[i] - minPrice
		}
	}
	return maxprofit
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
