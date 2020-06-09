package main

import (
	"fmt"
)

func main() {
	num := 220

	fmt.Println(translateNum(num))
}

func translateNum(num int) int {
	curr := 0
	tmp := -1
	res := 1
	last := res

	for num > 0 {
		curr = num % 10
		num = num / 10

		if curr != 0 && tmp >= 0 && curr*10+tmp <= 25 {
			res, last = res+last, res
		} else {
			last = res
		}
		tmp = curr
	}
	return res
}
