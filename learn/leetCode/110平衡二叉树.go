package main

import "fmt"

type TreeNode110 struct {
	Val   int
	Left  *TreeNode110
	Right *TreeNode110
}

func main() {
	//平衡二叉树的定义是：二叉树的每个节点左子树和右子树高度差的绝对值不超过1，
	//当节点高度为1时，即是叶节点
	var a, b, c, d, e TreeNode110
	a.Val = 1
	b.Val = 2
	c.Val = 3
	d.Val = 4
	e.Val = 5

	a.Left = &b
	a.Right = &c
	c.Left = &d
	c.Right = &e
	var res []int
	preOrder(&a, &res)
	fmt.Println(res)
}

func preOrder(root *TreeNode110, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

func isBalanced(root *TreeNode110) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode110) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
