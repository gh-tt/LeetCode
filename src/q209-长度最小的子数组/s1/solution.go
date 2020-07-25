package main

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	l, r, sum, m := 0, 0, 0, n+1
	for r < n {
		sum += nums[r]
		r++
		for sum >= s {
			m = min(m, r-l)
			sum -= nums[l]
			l++
		}
	}
	if m == n+1 {
		return 0
	}
	return m
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
