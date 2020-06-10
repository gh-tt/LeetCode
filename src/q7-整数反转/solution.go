package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-1 << 29))

}

func reverse(x int) int {
	res := 0

	for x != 0 {
		mod := x % 10
		x /= 10

		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && mod > math.MaxInt32%10) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && mod < math.MinInt32%10) {
			return 0
		}
		res = res*10 + mod
	}

	return res
}
