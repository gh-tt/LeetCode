package main

import (
	"fmt"
	"sort"
)

func main()  {
	matrix :=[][]int{[]int{1,3},[]int{2,3}}
	fmt.Println(kthSmallest(matrix,3))
}
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, matrix[i]...)
	}
	sort.Ints(arr)
	return arr[k-1]
}
