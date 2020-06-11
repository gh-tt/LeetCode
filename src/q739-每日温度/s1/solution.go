package main

import "fmt"

func main() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	nT := dailyTemperatures(t)
	fmt.Println(nT)
}
func dailyTemperatures(T []int) []int {
	n := len(T)
	nT := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if T[j] > T[i] {
				nT[i] = j - i
				break
			}
		}
	}
	return nT
}
