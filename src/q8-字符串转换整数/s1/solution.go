package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("  -" + strconv.Itoa(1<<33)))
}

func myAtoi(str string) int {
	res := 0
	n := len(str)

	for i := 0; i < n; i++ {
		if str[i] == ' ' {
			continue
		} else if str[i] == '+' || str[i] == '-' {
			tmp := str
			str = str[i+1:]
			if len(str) < 1 {
				return 0
			}
			if tmp[i] == '-' {
				res = getRes(str, res, true)
				res = -res
			} else {
				res = getRes(str, res, false)
			}
			return res
		} else if t := str[i] - '0'; t <= 9 {
			str = str[i:]
			return getRes(str, res, false)
		} else {
			return 0
		}
	}

	return res
}

func getRes(str string, res int, flag bool) int {
	for _, ch := range str {
		ch -= '0'
		if ch > 9 || ch < 0 {
			break
		}
		if flag == false && (res > math.MaxInt32/10 || (res == math.MaxInt32/10 && int(ch) > math.MaxInt32%10)) {
			return math.MaxInt32
		}
		if flag == true && (res > -math.MinInt32/10 || (res == -math.MinInt32/10 && int(ch) > -math.MinInt32%10)) {
			return -math.MinInt32
		}
		res = res*10 + int(ch)
	}
	return res
}
