package main

import (
	"container/heap"
	"fmt"
)

type IHeap [][3]int

func (I IHeap) Len() int {
	return len(I)
}

func (I IHeap) Less(i, j int) bool {
	return I[i][0] < I[j][0]
}

func (I IHeap) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func (I *IHeap) Push(x interface{}) {
	*I = append(*I, x.([3]int))
}

func (I *IHeap) Pop() interface{} {
	old := *I
	n := len(old)
	x := old[n-1]
	*I = old[:n-1]
	return x
}
func main() {
	matrix := [][]int{[]int{7, 9, 100}, []int{2, 40, 80}, []int{4, 50, 110}}
	fmt.Println(kthSmallest(matrix, 5))
}
func kthSmallest(matrix [][]int, k int) int {
	h := &IHeap{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, [3]int{matrix[i][0], i, 0})
	}
	for i := 0; i < k-1; i++ {
		curr := heap.Pop(h).([3]int)
		if curr[2] != len(matrix)-1 {
			heap.Push(h, [3]int{matrix[curr[1]][curr[2]+1], curr[1], curr[2] + 1})
		}
	}
	re := heap.Pop(h)
	return re.([3]int)[0]
}
