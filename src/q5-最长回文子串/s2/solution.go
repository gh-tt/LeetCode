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
	n := len(s)
	if n <= 1 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < n; i++ {
		left1, right1 := appendAroundCenter(s, n, i, i)
		left2, right2 := appendAroundCenter(s, n, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func appendAroundCenter(s string, strLen, left, right int) (int, int) {
	for left >= 0 && right < strLen && s[left] == s[right] {
		left, right = left-1, right+1
	}
	return left + 1, right - 1
}
