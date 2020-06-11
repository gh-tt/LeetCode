package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	n := len(values)
	res := 0
	for i := 0; i < n; i++ {

		for len(s) >= 2 {
			if s[:2] == symbols[i] {
				s = s[2:]
				res += values[i]
			} else if s[:1] == symbols[i] {

				s = s[1:]
				res += values[i]
			} else {
				break
			}
		}

		if len(s) == 1 && s == symbols[i] {
			s = s[1:]
			res += values[i]
		}
	}
	return res
}
