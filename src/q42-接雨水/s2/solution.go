package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 3, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(nums))
}

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	maxLeft, maxRight := make([]int, n), make([]int, n)
	res := 0
	maxLeft[0] = height[0]
	for l := 1; l < n; l++ {
		maxLeft[l] = max(height[l], maxLeft[l-1])
	}

	maxRight[n-1] = height[n-1]
	for r := n - 2; r >= 0; r-- {
		maxRight[r] = max(height[r], maxRight[r+1])
	}

	for i := 1; i < n-1; i++ {
		res = res + min(maxLeft[i], maxRight[i]) - height[i]
	}
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
