package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{0, 0, 0, 0}
	fmt.Println(threeSum(s))
}

func threeSum(nums []int) [][]int {
	n := len(nums)

	var tmp [][]int
	if n < 3 {
		return tmp
	}
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			return tmp
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := n - 1

		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				tmp = append(tmp, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return tmp
}
