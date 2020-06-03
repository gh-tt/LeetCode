package main

import (
	"fmt"
)

func main() {
	s := "abcdc"
	length := longestPalindrome(s)
	fmt.Println(length)
}

func longestPalindrome(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	dp := make([][]bool, length)
	end := length
	maxLen := 1

	for l := length - 1; l >= 0; l-- {
		dp[l] = make([]bool, length)
		dp[l][l] = true

		for r := length - 1; r > l; r-- {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}

			if dp[l][r] {
				if maxLen <= r-l+1 {
					maxLen = r - l + 1
					end = r+1
				}
			}
		}

	}
	return s[end-maxLen : end]
}
