package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{1, &ListNode{2, &ListNode{5, nil}}}
	l2 := ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	re := mergeTwoLists(&l1, &l2)
	for re != nil {
		fmt.Println(re.Val)
		re = re.Next
	}
}
//迭代法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	preHead := new(ListNode)

	pre := preHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}

	return preHead.Next
}
