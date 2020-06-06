package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 7, 3, 4, 3, 3, 5}

	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	max := 0
	for k := range m {
		if !m[k-1] {
			curNum := k
			count := 1
			for m[curNum+1] {
				count++
				curNum++
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
