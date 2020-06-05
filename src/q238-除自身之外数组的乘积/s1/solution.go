package main

import (
	"fmt"
)

func main() {
	s := []int{2,2,3,4,5}
	fmt.Println(productExceptSelf(s))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)

	left, right, res := make([]int, n), make([]int, n), make([]int, n)

	left[0] = 1

	for i := 1; i < n; i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	right[n-1] = 1

	for j := n - 2; j >= 0; j-- {
		right[j] = nums[j+1] * right[j+1]
	}

	for k := 0; k < n; k++ {
		res[k] = left[k] * right[k]
	}

	return res
}
