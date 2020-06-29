package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		} else {
			node := root.Right
			for node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left
			return root.Right
		}
	}
	return root
}
