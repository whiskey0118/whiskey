package main

func (root *TreeNode) min() int {
	if root == nil {
		return 0
	}

	for {
		if root.Left == nil {
			return root.Value
		}
		root = root.Left
	}
}
