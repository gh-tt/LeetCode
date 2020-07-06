package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if obstacleGrid[i][j] == 1 || (j > 0 && dp[j-1] == 0) {
					dp[j] = 0
				} else {
					dp[j] = 1
				}
			} else if j == 0 {
				if obstacleGrid[i][j] == 1 || dp[j] == 0 {
					dp[j] = 0
				} else {
					dp[j] = 1
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					dp[j] = 0
				} else {
					dp[j] = dp[j-1] + dp[j]
				}
			}
		}
	}
	return dp[n-1]
}
