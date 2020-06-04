package main

import (
	"fmt"
	"strings"
)

func main() {
	re := convert("leetcode", 2)
	fmt.Println(re)
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	if len(s) < numRows {
		numRows = len(s)
	}

	var tmp = make([]string, numRows)

	curr := 0
	flag := -1

	for _, val := range s {
		tmp[curr] += string(val)

		if curr == 0 || curr == numRows-1 {
			flag = -flag
		}
		curr += flag
	}

	return strings.Join(tmp, "")
}
