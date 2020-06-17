package main

import "fmt"

func main() {
	A := []int{2, 2, 2}
	fmt.Println(maxScoreSightseeingPair(A))
}

func maxScoreSightseeingPair(A []int) int {
	n := len(A)
	res := A[0] + A[1] - 1
	maxI := A[0]
	for i := 1; i < n; i++ {
		maxI = max(maxI, A[i-1]+i-1)
		res = max(res, maxI+A[i]-i)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
