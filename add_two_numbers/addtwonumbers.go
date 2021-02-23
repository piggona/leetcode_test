package main

import (
	"fmt"

	"github.com/piggona/leetcode_test/listnode"
)

func addTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	var cf int
	var result = l1
	var former *listnode.ListNode
	var sum int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &listnode.ListNode{
				Val: 0,
			}
			former.Next = l1
		}
		if l2 == nil {
			l2 = &listnode.ListNode{
				Val: 0,
			}
		}
		sum = l2.Val + l1.Val + cf
		l1.Val = sum % 10
		cf = sum / 10
		former = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if cf != 0 {
		former.Next = &listnode.ListNode{
			Val: 1,
		}
	}
	return result
}

func main() {
	list2 := listnode.CreateListNodeByList([]int{9, 9, 9, 9, 9, 9, 9})
	list1 := listnode.CreateListNodeByList([]int{9, 9, 9, 9})
	result := addTwoNumbers(list1, list2)
	fmt.Println(listnode.IterListNode(result))
}
