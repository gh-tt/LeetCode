package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 3, 3, 5}

	fmt.Println(removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {
	n := len(nums)
	left := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++

		}

	}
	return left
}
