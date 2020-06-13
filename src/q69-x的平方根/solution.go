package main

import "fmt"

func main() {
	fmt.Println(mySqrt(66))
}

func mySqrt(x int) int {
	l, r := 0, x
	res := 0

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
