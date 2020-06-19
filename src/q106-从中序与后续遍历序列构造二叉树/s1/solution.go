package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Println(root)
}

func buildTree(inorder, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	rootVal := postorder[n-1]
	i := 0
	for ; i < n; i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	root.Right = buildTree(inorder[i+1:], postorder[i:n-1])
	root.Left = buildTree(inorder[:i], postorder[:i])
	return root
}

//中序遍历 inorder = [9,3,20]
//后序遍历 postorder = [9,20,3]
/*
	 3
	/  \
    9  20

*/
