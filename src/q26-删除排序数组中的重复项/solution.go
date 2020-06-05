package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 5}

	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	left, right := 1, 1

	for right < n {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
