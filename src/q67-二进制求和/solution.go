package main

import "fmt"

func main() {
	a, b := "100", "10"
	fmt.Println(addBinary(a, b))
}
func addBinary(a string, b string) string {
	p, q := a, b
	if len(a) < len(b) {
		p, q = b, a
	}

	pLen, qLen := len(p), len(q)
	res := ""
	j := 0
	carry := 0
	for i := pLen - 1; i >= 0; i-- {
		tmp := 0
		if qLen >= pLen-i {
			tmp = int(p[i]-'0'+q[qLen-pLen+i]-'0') + carry
		} else {
			tmp = int(p[i]-'0') + carry
		}
		res = string(tmp%2+'0') + res
		carry = tmp / 2
		j++
	}
	if carry == 1 {
		res = "1" + res
	}

	return res
}
