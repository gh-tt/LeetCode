package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	dfs(root, &res, 0)

	if len(res) < 2 {
		return res
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func dfs(root *TreeNode, res *[][]int, depth int) {
	if root == nil {
		return
	}
	if len(*res) > depth {
		(*res)[depth] = append((*res)[depth], root.Val)
	} else {
		*res = append(*res, []int{root.Val})
	}
	dfs(root.Left, res, depth+1)
	dfs(root.Right, res, depth+1)
}
