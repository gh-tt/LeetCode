package main

import "fmt"

func main() {
	matrix := [][]byte{{'1', '1'}, {'1', '1'}}
	fmt.Println(maximalSquare(matrix))
}
func maximalSquare(matrix [][]byte) int {
	res := 0
	m := len(matrix)
	if m < 1 {
		return res
	}
	n := len(matrix[0])
	dp := make([]int, n)
	tmp := 0
	for i := 0; i < m; i++ {
		tmp = 0
		for j := 0; j < n; j++ {
			nextTmp := dp[j]
			if matrix[i][j] == '0' {
				dp[j] = 0
			} else if i == 0 || j == 0 {
				dp[j] = 1
				if res < 1 {
					res = 1
				}
			} else {
				dp[j] = min(tmp, min(dp[j], dp[j-1])) + 1
				if dp[j] > res {
					res = dp[j]
				}
			}
			tmp = nextTmp
		}
	}
	return res * res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
