package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 2, 4}

	fmt.Println(rob(nums))
}

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		tmp := second
		second = max(first+nums[i], second)
		first = tmp
	}
	return second
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
