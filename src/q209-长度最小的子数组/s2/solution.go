package main

import "sort"

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	ans := n + 1
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		index := sort.SearchInts(sums, target)
		if index <= n {
			ans = min(ans, index-i+1)
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func search(sums []int, target int) int {
	l, r := 0, len(sums)
	for l < r {
		mid := (l + r) >> 1
		if sums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
