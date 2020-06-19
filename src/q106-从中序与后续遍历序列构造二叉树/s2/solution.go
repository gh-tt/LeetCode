package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	root := buildTree(inorder, postorder)
	fmt.Println(root)
}

func buildTree(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	m := make(map[int]int)
	for k, v := range inorder {
		m[v] = k
	}
	return myBuildTree(m, inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}
func myBuildTree(m map[int]int, inorder, postorder []int, inLeft, inRight, postLeft, postRight int) *TreeNode {
	if postLeft > postRight || inLeft>inRight {
		return nil
	}
	rootVal := postorder[postRight]
	rootIndex := m[rootVal]
	rightTreeSize := inRight - rootIndex

	root := &TreeNode{Val: rootVal}
	root.Right = myBuildTree(m, inorder, postorder, rootIndex+1, inRight, postLeft+rootIndex-inLeft, postRight-1)
	root.Left = myBuildTree(m, inorder, postorder, inLeft, rootIndex-1, postLeft,postRight-rightTreeSize-1)
	return root
}

//中序遍历 inorder = [9,3,20]
//后序遍历 postorder = [9,20,3]
/*
	 3
	/  \
    9  20

*/
