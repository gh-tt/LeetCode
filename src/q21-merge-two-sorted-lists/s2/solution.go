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
//递归法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res *ListNode

	if l1.Val >= l2.Val {
		res = l2
		l2.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		l1.Next = mergeTwoLists(l1.Next, l2)
	}

	return res
}
