package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := ""
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res += strconv.Itoa(node.Val) + ","
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			res += "null,"
		}
	}
	return res[:len(res)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	vals, i := strings.Split(data, ","), 1
	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if vals[i] != "null" {
			nodeVal, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: nodeVal}
			queue = append(queue, node.Left)
		}
		i++
		if vals[i] != "null" {
			nodeVal, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{Val: nodeVal}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
