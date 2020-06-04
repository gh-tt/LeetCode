package main

import "fmt"

func main() {
	s := []int{10,100,100,20}
	fmt.Println(maxArea(s))
}

func maxArea(height []int) int {
	n := len(height)

	i := 0
	j := n - 1

	tmp := 0
	for i != j {
		h := min(height[i], height[j])
		area := (j - i) * h

		if area > tmp {
			tmp = area
		}

		if height[i] >= height[j] {
			j--
		} else {
			i++
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
