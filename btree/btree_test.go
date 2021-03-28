package btree

import (
	"fmt"
	"testing"
)

func TestBtree(t *testing.T) {
	btree := CreateBinaryTreeByList([]interface{}{7, 3, 15, nil, nil, 8, 20})
	list := DisplayBinaryTree(btree)
	fmt.Println(list)
}
