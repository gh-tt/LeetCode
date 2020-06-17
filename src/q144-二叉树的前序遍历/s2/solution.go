package main

import "fmt"

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
//前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return res
}

//中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		res = append(res, stack[len(stack)-1].Val)
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

//后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return res
}
