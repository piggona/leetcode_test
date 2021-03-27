package main

import (
	"fmt"

	"github.com/piggona/leetcode_test/listnode"
)

func rotateRight(head *listnode.ListNode, k int) *listnode.ListNode {
	if head == nil {
		return nil
	}
	top := &listnode.ListNode{
		Next: head,
	}
	var length int
	var cur, prev, bottom, right *listnode.ListNode
	cur = head
	for cur.Next != nil {
		length++
		prev = cur
		cur = cur.Next
	}
	bottom = prev
	pos := k % length
	if pos == 0 {
		return head
	}
	right = head
	cur = head
	for pos > 0 {
		right = right.Next
		pos--
	}
	for right != bottom {
		right = right.Next
		cur = cur.Next
	}
	top.Next = cur.Next
	cur.Next = nil
	bottom.Next = head
	return top.Next
}

func rotateRightStandard(head *listnode.ListNode, k int) *listnode.ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		n++
		iter = iter.Next
	}
	add := n - k%n
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	result := iter.Next
	iter.Next = nil
	return result
}

func main() {
	list := listnode.CreateListNodeByList([]int{1, 2, 3, 4, 5})
	fmt.Println(listnode.IterListNode(rotateRight(list, 2)))
}
