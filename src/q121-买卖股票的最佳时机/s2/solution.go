package main

import "fmt"

func main() {
	s := []int{30,2,40}
	fmt.Println(maxProfit(s))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	left := make([]int, n)
	right := make([]int, n)
	res := 0

	minLeft := prices[0]
	for i := 0; i < n; i++ {
		if prices[i] < minLeft {
			minLeft = prices[i]
			left[i] = minLeft
		} else {
			left[i] = minLeft
		}
	}

	maxRight := prices[n-1]
	for j := n - 1; j >= 0; j-- {
		if prices[j] > maxRight {
			maxRight = prices[j]
			right[j] = maxRight
		} else {
			right[j] = maxRight
		}
	}

	for k := 0; k < n; k++ {
		if left[k] < right[k] {
			tmp := right[k] - left[k]
			if tmp > res {
				res = tmp
			}
		}
	}

	fmt.Println(left,right)
	return res
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
