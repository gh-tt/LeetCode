package main

import (
	"fmt"
)

func main() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	nT := dailyTemperatures(t)
	fmt.Println(t)
	fmt.Println(nT)
}
func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)
	var stack []int

	for i := 0; i < n; i++ {
		t := T[i]
		for len(stack) > 0 && t > T[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return res
}
