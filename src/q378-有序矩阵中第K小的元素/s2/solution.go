package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{[]int{1, 3}, []int{2, 3}}
	fmt.Println(kthSmallest(matrix, 1))
}
func kthSmallest(matrix [][]int, k int) int {
	res := merge(matrix, 0, len(matrix)-1)
	return res[k-1]
}
func merge(matrix [][]int, l, r int) []int {
	if l == r {
		return matrix[l]
	}
	mid := (l + r) >> 1
	return mergeTwo(merge(matrix, l, mid), merge(matrix, mid+1, r))
}
func mergeTwo(left, right []int) (res []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return
}
