package main

import "fmt"

func main() {
	fmt.Println(intToRoman(4))
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	n := len(values)
	str := ""
	for i := 0; i < n && num >= 0; i++ {
		for values[i] <= num {
			num -= values[i]
			str += symbols[i]
		}
	}
	return str
}
