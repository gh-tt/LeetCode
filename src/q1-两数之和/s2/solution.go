package main

import (
	"fmt"
)

func main() {
	re := twoSum([]int{3, 3, 4}, 6)

	fmt.Println(re)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for k, v := range nums {
		if _, ok := m[v]; ok {
			return []int{m[v], k}
		}
		m[target-v] = k
	}

	return nil
}
