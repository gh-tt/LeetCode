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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}
