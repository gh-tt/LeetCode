package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for i := 0; i < len(lists); i++ {
		res = merge(res, lists[i])
	}
	return res
}

func merge(left, right *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy

	for left != nil && right != nil {
		if left.Val <= right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}
	return dummy.Next
}
