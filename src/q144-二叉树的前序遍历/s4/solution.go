package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := []int{1}
	a = append(a, []int{1, 2}...)
	fmt.Println(a)
	a = append(a, 3, 1)
	fmt.Println(a)
}

/*

用于测试的树，与上例中相同
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1

*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, stack := make([]int, 0), []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, tmp.Val)
		}
	}
	return res
}

func preorderNew(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, stack := make([]int, 0), []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
