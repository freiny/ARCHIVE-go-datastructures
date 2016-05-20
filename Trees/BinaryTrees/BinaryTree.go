package main

import "fmt"

func main() {
	root := &BinaryTree{value: 1}
	root.insert(2)
	root.insert(3)
	root.left.insert(4)
	root.left.insert(5)
	root.right.insert(6)
	root.right.insert(7)

	testPrint(root)
	testPrint(root.left, root.right)
	// print(root.left.left, root.left.right)
	// print(root.right.left, root.right.right)
}

func testPrint(t ...*BinaryTree) {
	for _, each := range t {
		fmt.Println(each.value, each.left.value, each.right.value)
		fmt.Println("-----------------------------------")
	}
}

func print(t *BinaryTree) {
	fmt.Print(t.value)
}

// BinaryTree ...
type BinaryTree struct {
	value       byte
	left, right *BinaryTree
}

func (t *BinaryTree) insert(value byte) {
	switch {
	case t.left == nil:
		t.left = &BinaryTree{value: value}
	case t.right == nil:
		t.right = &BinaryTree{value: value}
	default:
		t.left = &BinaryTree{value: value, left: t.left}
	}
}

func (t *BinaryTree) delete() {
}

// Tree Format: parent.leftChild.rightChild
// 8.5.4
// 5.9.7 4..11
// 9.. 7.1.12 11.3.
// 1.. 12.2. 3..
// 2..
//
// PreOrder - 8, 5, 9, 7, 1, 12, 2, 4, 11, 3
// InOrder - 9, 5, 1, 7, 2, 12, 8, 4, 3, 11
// PostOrder - 9, 1, 2, 12, 7, 5, 3, 11, 4, 8
// LevelOrder - 8, 5, 4, 9, 7, 11, 1, 12, 3, 2

// Depth-First
func (t *BinaryTree) traversePreOrder() {
}
func (t *BinaryTree) traverseInOrder() {
}
func (t *BinaryTree) traversePostOrder() {
}

// Breadth-First
func (t *BinaryTree) traverseLevelOrder(f func(*BinaryTree)) {
	order := map[int][]*BinaryTree{}
	fmt.Println(order)
}
