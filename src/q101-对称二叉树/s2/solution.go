package main

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queen := make([]*TreeNode, 2)

	queen[0] = root.Left
	queen[1] = root.Right

	for len(queen) > 0 {
		q := queen[0]
		p := queen[1]

		queen = queen[2:]

		if q == nil && p == nil {
			continue
		}

		if q == nil || p == nil {
			return false
		}

		if q.Val == p.Val {
			queen = append(queen, q.Left)
			queen = append(queen, p.Right)
			queen = append(queen, q.Right)
			queen = append(queen, p.Left)
		} else {
			return false
		}
	}

	return true
}
