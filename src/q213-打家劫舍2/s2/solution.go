package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 1}

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

	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {
	//n := len(nums)

	if end-start == 0 {
		return nums[start]
	}

	first := nums[start]
	second := max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
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
