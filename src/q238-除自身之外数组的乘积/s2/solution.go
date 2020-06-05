package main

import (
	"fmt"
)

func main() {
	s := []int{3, 2, 6, 4, 5}
	fmt.Println(productExceptSelf(s))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)

	res := make([]int, n)

	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	r := 1
	for j := n - 1; j >= 0; j-- {
		res[j] = res[j] * r
		r = r * nums[j]
	}
	return res
}
