package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	m := make(map[int]int)
	for k, v := range inorder {
		m[v] = k
	}
	return myBuildTree(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, m)
}

func myBuildTree(preorder, inorder []int, preLeft, preRight, inLeft, inRight int, m map[int]int) *TreeNode {
	if preLeft > preRight {
		return nil
	}
	rootVal := preorder[preLeft]
	//中序遍历中找到根节点index
	rootIndex := m[rootVal]
	//左子树的节点数
	leftTreeSize := rootIndex - inLeft

	root := &TreeNode{Val: rootVal}
	root.Left = myBuildTree(preorder, inorder, preLeft+1, preLeft+leftTreeSize, inLeft, rootIndex-1, m)
	root.Right = myBuildTree(preorder, inorder, preLeft+leftTreeSize+1, preRight, rootIndex+1, inRight, m)
	return root
}

//前序遍历 preorder = [3,9,20]
//中序遍历 inorder = [9,3,20]
/*
	 3
	/  \
    9  20

*/
