package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	curr := root
	pre := curr
	flag := 0
	for curr != nil && curr.Val != key {
		pre = curr
		if curr.Val < key {
			curr = curr.Right
			flag = 1
		} else {
			curr = curr.Left
			flag = -1
		}
	}
	if curr ==nil {
		return root
	}

	if root.Val == key {
		if root.Right != nil {
			tmp := root.Right
			rightPre := tmp
			for tmp != nil {
				rightPre = tmp
				tmp = tmp.Left
			}
			rightPre.Left = root.Left
			return root.Right
		} else {
			return root.Left
		}
	} else {
		if curr.Right != nil {
			tmp := curr.Right
			rightPre := tmp
			for tmp != nil {
				rightPre = tmp
				tmp = tmp.Left
			}
			rightPre.Left = curr.Left
			if flag == -1 {
				pre.Left = curr.Right
			} else {
				pre.Right = curr.Right
			}
		} else {
			if flag == -1 {
				pre.Left = curr.Left
			} else {
				pre.Right = curr.Left
			}
		}

	}
	return root
}
