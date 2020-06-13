package main

import "fmt"

func main() {
	re := climbStairs(20)
	fmt.Println(re)
}

func climbStairs(n int) int {
	first := 0
	second := 1
	re := 1
	for i := 1; i <= n; i++ {
		re = first + second
		first = second
		second = re
	}
	return re
}
