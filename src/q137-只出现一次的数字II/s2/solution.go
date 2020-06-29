package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}
func singleNumber(nums []int) int {
	x, y := 0, 0
	for _, v := range nums {
		y = ^x & (y ^ v)
		x = ^y & (x ^ v)
	}
	return y
}
