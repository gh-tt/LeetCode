package main

import "fmt"

func main() {
	s := []int{6, 8, 30, 2, 30, 7, 8, 7, 7}
	fmt.Println(maxProfit(s))
}

func maxProfit(prices []int) int {
	n := len(prices)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j] > prices[i] {
				tmp := prices[j] - prices[i]
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
