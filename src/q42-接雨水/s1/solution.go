package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 3, 1, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(nums))
}

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	res := 0
	for i := 1; i < n-1; i++ {
		maxLeft := 0
		for l := i - 1; l >= 0; l-- {
			if height[l] > maxLeft {
				maxLeft = height[l]
			}
		}

		maxRight := 0
		for r := i + 1; r < n; r++ {
			if height[r] > maxRight {
				maxRight = height[r]
			}
		}

		minLR := min(maxRight, maxLeft)
		if minLR > height[i] {
			res = res + minLR - height[i]
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
