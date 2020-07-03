package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main()  {
	num :=[]int{-10,-6,-3,0,1,5,8,9}
	re :=sortedArrayToBST(num)
	fmt.Println(re)
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := (n - 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	if mid+1 < n {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
