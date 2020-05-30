package main

import "fmt"

func main() {
	nums := []int{2,3,5,1}

	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return max(robRange(nums[:n-1]), robRange(nums[1:]))
}

func robRange(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		tmp := second
		second = max(second, first+nums[i])
		first = tmp
	}
	return second
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}
