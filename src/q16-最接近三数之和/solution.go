package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(s, 1))
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)

	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if tmp > target {
				r--
			} else if tmp < target {
				l++
			} else {
				return tmp
			}
			if distance(res, target) > distance(tmp, target) {
				res = tmp
			}
		}
	}

	return res
}

func distance(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
