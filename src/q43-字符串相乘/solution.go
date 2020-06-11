package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiply("123", "100"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	res := make([]int, n1+n2)

	for i := n1 - 1; i >= 0; i-- {
		num1 := num1[i] - '0'
		for j := n2 - 1; j >= 0; j-- {
			num2 := num2[j] - '0'
			sum := res[i+j+1] + int(num1)*int(num2)
			res[i+j] += sum / 10
			res[i+j+1] = sum % 10
		}
	}

	str := ""
	for k, v := range res {
		if k == 0 && v == 0 {
			continue
		}
		str += string(v + '0')
	}
	return str
}
