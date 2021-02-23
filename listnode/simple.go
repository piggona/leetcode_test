package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNodeByList(list []int) *ListNode {
	var result *ListNode
	var cur *ListNode
	result = &ListNode{
		Val: list[0],
	}
	cur = result
	for i := 1; i < len(list); i++ {
		temp := &ListNode{
			Val: list[i],
		}
		cur.Next = temp
		cur = temp
	}
	return result
}

func IterListNode(listNode *ListNode) []int {
	result := []int{}
	for listNode != nil {
		result = append(result, listNode.Val)
		listNode = listNode.Next
	}
	return result
}
