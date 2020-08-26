package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inorderRecursive(root)
}

func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append(inorderRecursive(root.Left), root.Val)
	rest = append(rest, inorderRecursive(root.Right)...)
	return rest
}

func main() {
	//var t TreeNode
	//t.Val = 10
	//a := inorderRecursive(&t)
	//fmt.Println(a)
	//var b []int

}
