package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 3, 3, 6}

	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	max := 1
	count := 1
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}

		if nums[i]+1 == nums[i+1] {
			count++
		} else {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}
