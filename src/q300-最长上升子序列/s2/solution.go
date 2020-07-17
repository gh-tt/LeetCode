package main

import "fmt"

func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	tail := make([]int, n)
	tail[0] = nums[0]
	end := 0
	for i := 1; i < n; i++ {
		if nums[i] > tail[end] {
			end++
			tail[end] = nums[i]
		} else {
			left := 0
			right := end
			for left < right {
				mid := left + (right-left)/2
				if tail[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			tail[left] = nums[i]
		}
	}
	end++
	return end
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
