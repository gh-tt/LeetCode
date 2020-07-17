package main

import "fmt"

func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
