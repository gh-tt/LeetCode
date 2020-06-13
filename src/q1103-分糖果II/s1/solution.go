package main

import "fmt"

func main() {
	re := distributeCandies(120, 3)
	fmt.Println(re)
}
func distributeCandies(candies int, num_people int) []int {
	re := make([]int, num_people)
	index := 1
	for candies >= 1 {
		for i := 0; i < num_people && candies > 0; i++ {
			if index > candies {
				re[i] += candies
				candies = 0
			} else {
				re[i] += index
				candies -= index
				index++
			}
		}
	}
	return re
}
