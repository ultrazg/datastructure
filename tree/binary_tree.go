package main

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

type binaryTree struct {
	root *treeNode
}

func newBinaryTree() *binaryTree {
	return &binaryTree{}
}

func (bt *binaryTree) insert(val int) {
	newNode := &treeNode{val: val}

	if bt.root == nil {
		bt.root = newNode

		return
	}

	current := bt.root
	for {
		if val < current.val {
			if current.left == nil {
				current.left = newNode

				return
			}

			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode

				return
			}

			current = current.right
		}
	}
}

func printTree(node *treeNode) {
	if node == nil {
		return
	}

	printTree(node.left)
	fmt.Println(node.val)
	printTree(node.right)
}

func main() {
	bt := newBinaryTree()

	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, val := range values {
		bt.insert(val)
	}

	fmt.Println("Binary Tree:")
	printTree(bt.root)
}
