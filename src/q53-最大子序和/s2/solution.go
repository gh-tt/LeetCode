package main

import "fmt"

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	n := len(nums)
	res := nums[0]
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum > res {
				res = sum
			}
		}
	}
	return res
}
