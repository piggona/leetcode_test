package main

import (
	"fmt"

	"github.com/piggona/leetcode_test/listnode"
)

func deleteDuplicates(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return nil
	}
	top := &listnode.ListNode{
		Next: head,
	}
	var prev, cur, next *listnode.ListNode
	prev = top
	cur = head
	for cur != nil {
		next = cur.Next
		if next != nil && cur.Val == next.Val {
			prev.Next = next
		} else {
			prev = cur
		}
		cur = next
	}
	return top.Next
}

func main() {
	list := listnode.CreateListNodeByList([]int{1, 1})
	fmt.Println(listnode.IterListNode(deleteDuplicates(list)))
}
