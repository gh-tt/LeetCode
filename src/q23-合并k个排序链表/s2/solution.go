package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoList(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwoList(left, right *ListNode) *ListNode {
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
