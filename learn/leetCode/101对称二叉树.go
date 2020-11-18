package main

func main() {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSymmetric(root.Right) == false && isSymmetric(root.Left) != true {
		return false
	} else if isSymmetric(root.Right) == true && isSymmetric(root.Left) == false {
		return true
	} else {
		return true
	}
}

func is(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if is(root.Right) == false && is(root.Left) != true {
		return false
	} else if is(root.Right) == true && is(root.Left) == false {
		return true
	} else {
		return true
	}
}
