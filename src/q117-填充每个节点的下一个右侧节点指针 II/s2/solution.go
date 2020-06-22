package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	curr := root
	for curr != nil {
		var pre *Node
		var down *Node
		for curr != nil {
			if curr.Left != nil {
				if pre != nil {
					pre.Next = curr.Left
				} else {
					down = curr.Left
				}
				pre = curr.Left
			}
			if curr.Right != nil {
				if pre != nil {
					pre.Next = curr.Right
				} else {
					down = curr.Right
				}
				pre = curr.Right
			}
			curr = curr.Next
		}
		curr = down
	}
	return root
}
