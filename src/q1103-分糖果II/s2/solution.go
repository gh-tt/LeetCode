package main

import "fmt"

func main() {
	re := distributeCandies(120, 3)
	fmt.Println(re)
}
func distributeCandies(candies int, num_people int) []int {
	re := make([]int, num_people)
	index := 0
	for candies > 0 {
		re[index%num_people] += index + 1
		candies -= index + 1
		index++
	}
	re[(index-1)%num_people] += candies
	return re
}
