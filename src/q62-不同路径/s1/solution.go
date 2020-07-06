package main

func uniquePaths(m int, n int) int {
	if n < 1 {
		return 1
	}
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[j] = 1
			} else if j == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}
	return dp[n-1]
}
