package main

import "fmt"

func main()  {
	nums :=[]int{-1,0,4,6,7,9,11}
	fmt.Println(search(nums,11))
}
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
