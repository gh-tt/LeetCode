package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	res := 0
	flag := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			if flag == 1 && (res > math.MaxInt32/10 || (res == math.MaxInt32/10 && int(v-'0') > math.MaxInt32%10)) {
				return math.MaxInt32
			}
			if flag == -1 && (res > -math.MinInt32/10 || (res == -math.MinInt32/10 && int(v-'0') > -math.MinInt32%10)) {
				return math.MinInt32
			}
			res = res*10 + int(v-'0')
		} else if v == '+' && i == 0 {
			continue
		} else if v == '-' && i == 0 {
			flag = -flag
		} else {
			break
		}
	}
	res = res * flag
	return res
}
