package main

import "fmt"

func main() {
	nums := []int{2, 1, 0, 0, 2, 1, 2, 2}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	p, q := 0, n-1
	curr := 0

	for curr <= q {
		if nums[curr] == 0 {
			tmp := nums[p]
			nums[p] = nums[curr]
			nums[curr] = tmp
			p++
			curr++
		} else if nums[curr] == 2 {
			tmp := nums[q]
			nums[q] = nums[curr]
			nums[curr] = tmp
			q--
		} else {
			curr++
		}

	}
}
