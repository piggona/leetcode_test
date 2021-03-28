package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateBinaryTreeByList(list []interface{}) *TreeNode {
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	queue := []*TreeNode{}
	var right bool
	var tree *TreeNode
	for _, val := range list {
		value, ok := val.(int)
		var node *TreeNode
		if ok {
			node = &TreeNode{Val: value}
		}
		if len(queue) == 0 {
			if !ok {
				return nil
			}
			tree = &TreeNode{Val: value}
			queue = append(queue, tree)
		} else {
			front := queue[0]
			if right {
				front.Right = node
				queue = queue[1:]
			} else {
				front.Left = node
			}
			right = !right
			if node != nil {
				queue = append(queue, node)
			}
		}
	}
	return tree
}

func DisplayBinaryTree(tree *TreeNode) []interface{} {
	result := []interface{}{tree.Val}
	queue := []*TreeNode{tree}
	for len(queue) != 0 {
		nodes := []*TreeNode{queue[0].Left, queue[0].Right}
		for _, node := range nodes {
			if node != nil {
				result = append(result, node.Val)
				queue = append(queue, node)
			} else {
				result = append(result, nil)
			}
		}
		queue = queue[1:]
	}
	return result
}
