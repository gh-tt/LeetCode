package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "a"))
}

func strStr(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)
	if nLen == 0 {
		return 0
	}
	if nLen > hLen {
		return -1
	}
	p := 0
	for p < hLen-nLen+1 {

		for p < hLen-nLen+1 && haystack[p] != needle[0] {
			p++
		}

		curLen, q := 0, 0
		for q < nLen && p < hLen && haystack[p] == needle[q] {
			curLen++
			q++
			p++
		}

		if curLen == nLen {
			return p - nLen
		}

		p = p - curLen + 1
	}
	return -1
}
