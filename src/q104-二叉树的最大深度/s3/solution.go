package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	root  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	stack := []pair{pair{root, 1}}
	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		node := tmp.root
		depth := tmp.depth
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, pair{node.Right, depth + 1})
		}
		if node.Left != nil {
			stack = append(stack, pair{node.Left, depth + 1})
		}
		if node.Right == nil && node.Left == nil {
			res = max(res, depth)
		}
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
