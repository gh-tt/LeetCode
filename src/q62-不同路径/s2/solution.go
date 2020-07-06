package main

import "fmt"

func main()  {
	fmt.Println(uniquePaths(7,3))
}

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j >= 1 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
