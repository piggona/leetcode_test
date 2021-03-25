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
		Val:  -101,
		Next: head,
	}
	var prev, cur, next *listnode.ListNode
	var flag bool
	cur = top
	next = head
	for cur != nil && cur.Next != nil {
		if cur.Val != next.Val {
			if flag {
				prev.Next = next
				flag = false
			} else {
				prev = cur
			}
		} else {
			flag = true
		}
		cur = next
		next = next.Next
	}
	if flag {
		prev.Next = nil
	}
	return top.Next
}

func main() {
	list := listnode.CreateListNodeByList([]int{3, 3, 4, 4})
	fmt.Println(listnode.IterListNode(deleteDuplicates(list)))
}
