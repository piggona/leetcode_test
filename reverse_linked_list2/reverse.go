package main

import (
	"fmt"

	"github.com/piggona/leetcode_test/listnode"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *listnode.ListNode, left int, right int) *listnode.ListNode {
	sentry := &listnode.ListNode{}
	top := sentry
	sentry.Next = head
	var cur *listnode.ListNode = top
	var next *listnode.ListNode
	var prev *listnode.ListNode
	var count int
	for i := 0; i < left-1; i++ {
		cur = cur.Next
	}
	top = cur
	cur = top.Next
	for cur != nil && count <= right-left {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		count++
	}
	if top.Next != nil {
		top.Next.Next = cur
		top.Next = prev
	}
	return sentry.Next
}

func main() {
	result := reverseBetween(listnode.CreateListNodeByList([]int{1, 2, 3, 4, 5}), 2, 4)
	fmt.Println(listnode.IterListNode(result))
}
