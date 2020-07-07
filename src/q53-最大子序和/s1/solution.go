package main

import "fmt"

func main()  {
	nums :=[]int{1,2}
	fmt.Println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}
	dp := nums[0]
	res := dp
	for i := 1; i < n; i++ {
		if nums[i] < nums[i]+dp {
			dp = nums[i] + dp
		} else {
			dp = nums[i]
		}
		if dp > res {
			res = dp
		}
	}
	return res
}
