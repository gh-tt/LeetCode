package main

import "fmt"

func main() {
	A := []int{2,2,2}
	fmt.Println(maxScoreSightseeingPair(A))
}

func maxScoreSightseeingPair(A []int) int {
	n := len(A)
	if n < 3 {
		return A[0] + A[1] - 1
	}
	maxLeft, maxRight := make([]int, n), make([]int, n)
	maxLeft[0] = A[0] + 0
	for i := 1; i < n-1; i++ {
		maxLeft[i] = max(maxLeft[i-1], A[i]+i)
	}
	maxRight[n-1] = A[n-1] - n + 1
	for j := n - 2; j >= 0; j-- {
		maxRight[j] = max(maxRight[j+1], A[j]-j)
	}

	maxLeft = maxLeft[:n-1]
	maxRight = maxRight[1:]
	res := maxLeft[0] + maxRight[0]
	for k := 0; k < n-1; k++ {
		if res < maxLeft[k]+maxRight[k] {
			res = maxLeft[k] + maxRight[k]
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
