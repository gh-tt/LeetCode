package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	dfs(root, nil)
	return root
}
func dfs(node, next *Node) {
	if node != nil {
		node.Next = next
		dfs(node.Left, node.Right)
		if node.Next != nil {
			dfs(node.Right, node.Next.Left)
		} else {
			dfs(node.Right, nil)
		}
	}
}
