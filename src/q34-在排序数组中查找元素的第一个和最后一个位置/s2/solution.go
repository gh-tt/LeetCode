package main

import (
	"fmt"
)

func main() {
	num := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(num, 7))
}
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) < 1 {
		return res
	}
	n := len(nums)
	start, end := 0, n-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		res[0] = start
	} else if nums[end] == target {
		res[0] = end
	} else {
		return res
	}

	start, end = 0, n-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[end] == target {
		res[1] = end
	} else if nums[start] == target {
		res[1] = start
	} else {
		return res
	}
	return res
}
