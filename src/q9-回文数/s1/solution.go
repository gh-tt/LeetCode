package main

import (
	"fmt"
	"strconv"
)

func main() {
	re := isPalindrome(11)
	fmt.Println(re)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmp := strconv.Itoa(x)

	n := len(tmp)
	for i := 0; i < n/2; i++ {
		if tmp[i] != tmp[n-i-1] {
			return false
		}
	}
	return true
}
