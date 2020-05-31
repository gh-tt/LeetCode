package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root.Left, root.Right)
}

func check(treeLeft *TreeNode, treeRight *TreeNode) bool {
	if treeLeft == nil && treeRight == nil {
		return true
	}

	if treeLeft == nil || treeRight == nil {
		return false
	}

	return treeLeft.Val == treeRight.Val &&
		check(treeLeft.Left, treeRight.Right) &&
		check(treeLeft.Right, treeRight.Left)
}
