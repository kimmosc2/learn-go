package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	Sort([]int{2, 3, 7, 9, 1, 5, 7})
}

// Sort利用二叉树实现插入排序
func Sort(values []int) {
	root := new(tree)
	for _, j := range values {
		add(root, j)
	}
	values = appendValues(values[:0],root)
	fmt.Println(values)
}

func add(root *tree, value int) *tree {
	if root == nil {
		root := new(tree)
		root.value = value
		return root
	}
	if value < root.value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values,t.left)
		values = append(values,t.value)
		values = appendValues(values,t.right)
	}
	return values
}
