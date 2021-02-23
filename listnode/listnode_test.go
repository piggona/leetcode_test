package listnode

import "testing"

func TestSimple(t *testing.T) {
	list := []int{2, 4, 3}
	listNode := CreateListNodeByList(list)
	result := IterListNode(listNode)
	for i := 0; i < len(list); i++ {
		if list[i] != result[i] {
			t.Errorf("Bad result: %v\n", result)
			return
		}
	}
}
