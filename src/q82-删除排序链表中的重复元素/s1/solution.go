package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{6, nil}}}}
	re := deleteDuplicates(&head)
	printList(re)
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	m := make(map[int]bool)

	cur := head

	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			m[cur.Val] = true
		} else {
			m[cur.Val] = false
		}
		cur = cur.Next
	}
	var s []int
	for k, v := range m {
		if v == false {
			s = append(s, k)
		}
	}
	sort.Ints(s)

	dummy := new(ListNode)
	newL := dummy
	for _, v := range s {
		newL.Next = &ListNode{v, nil}
		newL = newL.Next
	}
	return dummy.Next
}
