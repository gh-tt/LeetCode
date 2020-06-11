package main

import (
	"container/list"
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
	stack := list.New()

	for i := 0; i < n; i++ {
		for stack.Len() > 0 && T[i] > T[stack.Back().Value.(int)] {
			res[stack.Back().Value.(int)] = i - stack.Back().Value.(int)
			stack.Remove(stack.Back())
		}
		stack.PushBack(i)
	}
	return res
}
