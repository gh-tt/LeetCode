package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	fm := make(map[int]*TreeNode)
	visited := make(map[int]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			fm[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			fm[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		visited[p.Val] = true
		p = fm[p.Val]
	}
	for q != nil {
		if _, exist := visited[q.Val]; exist {
			return q
		}
		q = fm[q.Val]
	}
	return nil
}
