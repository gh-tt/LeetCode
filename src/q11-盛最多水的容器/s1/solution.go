package main

import "fmt"

func main() {
	s := []int{1,8,30,2,30,4,8,3,7}
	fmt.Println(maxArea(s))
}

func maxArea(height []int) int {
	n := len(height)

	tmp := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			h := min(height[i], height[j])
			area := (j - i) * h

			if area > tmp {
				tmp = area
			}
		}
	}
	return tmp
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
