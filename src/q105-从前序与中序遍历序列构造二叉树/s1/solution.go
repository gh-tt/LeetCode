package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main()  {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	root :=buildTree(preorder,inorder)
	fmt.Println(root)
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

//前序遍历 preorder = [3,9,20]
//中序遍历 inorder = [9,3,20]
/*
	 3
	/  \
    9  20

*/
