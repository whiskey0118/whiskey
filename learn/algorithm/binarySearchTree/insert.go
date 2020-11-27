package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) insert(value int) {
	var res *TreeNode
	res.Value = value
	if root == nil {
		root = res
	} else {
		insertNode(root, res)
	}

}

func insertNode(root, node *TreeNode) {
	if root.Value == node.Value {
		return
	}

	if root.Value > node.Value {
		if root.Left == nil {
			root.Left = node
		} else {
			insertNode(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			insertNode(root.Right, node)
		}
	}
}
