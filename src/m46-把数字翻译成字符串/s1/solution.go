package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1221

	fmt.Println(translateNum(num))
}

func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {

		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '5') {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	fmt.Println(dp)
	return dp[n]
}
