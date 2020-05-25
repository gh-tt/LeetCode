package main

import "fmt"

func main() {
	re := twoSum([]int{3, 2, 4}, 6)

	fmt.Println(re)
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
