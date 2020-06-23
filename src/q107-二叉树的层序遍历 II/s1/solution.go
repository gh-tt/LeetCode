package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res, list := make([][]int, 0), []*TreeNode{root}
	for len(list) > 0 {
		l := len(list)
		tmp := make([]int, 0)
		for i := 0; i < l; i++ {
			node := list[0]
			list = list[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		res = append([][]int{tmp}, res...)
	}
	return res
}
