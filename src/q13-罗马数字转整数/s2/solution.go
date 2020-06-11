package main

import "fmt"

func main() {
	fmt.Println(romanToInt("XL"))
}

func romanToInt(s string) int {
	m := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	res := 0
	n := len(s)
	for k, v := range s {
		if k < n-1 && m[string(v)] < m[string(s[k+1])] {
			res -= m[string(v)]
		} else {
			res += m[string(v)]
		}
	}
	return res
}
