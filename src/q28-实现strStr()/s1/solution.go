package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3, 4, 3, 3, 5}

	fmt.Println(strStr("nums", "12345"))
}

func strStr(haystack string, needle string) int {
	hLen,nLen := len(haystack),len(needle)
	if nLen == 0 {
		return 0
	}
	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
