package main

import (
	"fmt"
)

func main() {
	num := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(num, 11))
}
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) < 1 {
		return res
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == target {
			res[0] = i
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] == target {
			res[1] = i
			break
		}
	}
	return res
}
