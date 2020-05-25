package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{9, nil}
	l2 := ListNode{5, nil}
	re := addTwoNumbers(&l1, &l2)
	fmt.Println(re, re.Next, re.Next.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	node1 := l1
	node2 := l2
	l3 := new(ListNode)
	node3 := l3
	carry := 0

	for node1 != nil || node2 != nil || carry > 0 {
		node3.Next = new(ListNode)
		node3 = node3.Next

		if node1 != nil {
			carry += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			carry += node2.Val
			node2 = node2.Next
		}

		node3.Val = carry % 10
		carry = carry / 10
	}

	return l3.Next
}
