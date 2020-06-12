package main

import "fmt"

func main() {
	re := isHappy(19)
	fmt.Println(re)
}

func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		fast = step(step(fast))
		slow = step(slow)
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
