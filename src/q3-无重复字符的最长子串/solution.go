package main

import "fmt"

func main() {
	s := "abccbaadbbef"
	length := lengthOfLongestSubString(s)
	fmt.Println(length)
}

func lengthOfLongestSubString(s string) int {
	m := make([]int, 128)

	n := len(s)

	j, re := 0, 0

	for i := 0; i < n; i++ {
		if i != 0 {
			m[s[i-1]]--
		}

		for j < n && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}

		re = max(re, j-i)
	}

	return re
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
