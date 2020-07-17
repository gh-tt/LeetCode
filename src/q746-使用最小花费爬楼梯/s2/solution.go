package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

func minCostClimbingStairs(cost []int) int {
	first := 0
	second := 0
	res := 0
	for i := 2; i <= len(cost); i++ {
		res = min(second+cost[i-1], first+cost[i-2])
		first = second
		second = res
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
