package main

func bstToGst(root *TreeNode) *TreeNode {

}

func test1038(root *TreeNode) *TreeNode {
	res := root
	var n int
	if root == nil {
		return res
	}

	if n > root.Right.Val {
		res.Right.Val = n
	}

}
