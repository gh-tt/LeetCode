package main

func maximalSquare(matrix [][]byte) int {
	res := 0
	m := len(matrix)
	if m < 1 {
		return res
	}
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			} else if i == 0 || j == 0 {
				dp[i][j] = 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
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
