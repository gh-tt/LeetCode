package main

import (
	"fmt"
)

func main() {
	nums := []int{0,1,0,2,1,0,1,3,2,1,2,1}

	fmt.Println(trap(nums))
}

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	res := 0
	for left <= right {
		if leftMax < rightMax {
			res += max(0, leftMax-height[left])
			leftMax = max(leftMax, height[left])
			left++
		} else {
			res += max(0, rightMax-height[right])
			rightMax = max(rightMax, height[right])
			right--
		}

	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

