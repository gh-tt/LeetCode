package main

import (
	"fmt"
)

func main() {
	fmt.Println(addDigits(678))
}
func addDigits(num int) int {
	for num >= 10 {
		next := 0
		for num != 0 {
			next = next + num%10
			num /= 10
		}
		num = next
	}
	return num
}
