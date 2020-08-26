package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	//如果t是一个空指针，就返回 &tree{value: value}
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	//如果value小于t.value
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		fmt.Println(t.right == nil)
		t.right = add(t.right, value)
	}
	return t
}

//将每次输入的value添加到values切片中
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	b := []int{1, 5, 9, 7, 3}
	Sort(b)
	fmt.Println(b)
}
