package main

import (
	"fmt"
)

func main() {
	num := 220

	fmt.Println(translateNum(num))
}

func translateNum(num int) int {
	dp := []int{1}

	tmp := -1
	curr := 0
	i := 0

	for num > 0 {
		dp = append(dp, 0)
		i++

		curr = num % 10
		num = num / 10
		dp[i] = dp[i-1]
		if curr != 0 && tmp >= 0 && curr*10+tmp <= 25 {
			dp[i] += dp[i-2]
		}
		tmp = curr
	}
	fmt.Println(dp)
	return dp[i]
}
