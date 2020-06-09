package main

import (
	"fmt"
)

func main() {
	s := []int{7,6,4,3,1}
	fmt.Println(maxProfit(s))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxFit := 0
	for i := 1; i < n; i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			maxFit += tmp
		}
	}
	return maxFit
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
