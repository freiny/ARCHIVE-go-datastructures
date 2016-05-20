package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Tree Format: parent.leftChild.rightChild
	root := buildTree(`
		8.5.4
		5.9.7 4..11
		9.. 7.1.12 11.3.
		1.. 12.2. 3..
		2..
	`)

	root.traverseLevelOrder(print)
	fmt.Println()
	// OUTPUT:
	// 8.5.4
	// 5.9.7
	// 4..11
	// 9..
	// 7.1.12
	// 11.3.
	// 1..
	// 12.2.
	// 3..
	// 2..

}

func print(t *BinaryTree, level int, levelChange bool) {
	if levelChange {
		fmt.Println()
	}

	if t != nil {
		fmt.Print(t.value)
	}
	fmt.Print(".")
	if t.left != nil {
		fmt.Print(t.left.value)
	}
	fmt.Print(".")
	if t.right != nil {
		fmt.Print(t.right.value)
	}
	fmt.Print(" ")
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

func (t *BinaryTree) insertLeft(value byte) {
	if t.left == nil {
		t.left = &BinaryTree{value: value}
	} else {
		t.left = &BinaryTree{value: value, left: t.left}
	}
}

func (t *BinaryTree) insertRight(value byte) {
	if t.right == nil {
		t.right = &BinaryTree{value: value}
	} else {
		t.right = &BinaryTree{value: value, right: t.right}
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
func (t *BinaryTree) traverseLevelOrder(f func(*BinaryTree, int, bool)) {
	order := map[int][]*BinaryTree{}
	order[0] = []*BinaryTree{}

	var traverse func(bt *BinaryTree, level int, order map[int][]*BinaryTree)
	traverse = func(bt *BinaryTree, level int, order map[int][]*BinaryTree) {
		if bt == nil {
			return
		}
		order[level] = append(order[level], bt)
		traverse(bt.left, level+1, order)
		traverse(bt.right, level+1, order)
	}
	traverse(t, 0, order)

	for i := 0; order[i] != nil; i++ {
		levelChange := true
		for _, node := range order[i] {
			f(node, i, levelChange)
			levelChange = false
		}
	}
}

func buildTree(s string) *BinaryTree {

	linkMap := map[byte][2]byte{}
	treeMap := map[byte]*BinaryTree{}

	rows := strings.Split(strings.TrimSpace(s), "\n")
	iRootKey, _ := strconv.Atoi(strings.Split(strings.TrimSpace(rows[0]), ".")[0])
	rootKey := byte(iRootKey)
	for _, row := range rows {
		trees := strings.Split(row, " ")
		toByte := func(s string) byte {
			i, _ := strconv.Atoi(s)
			return byte(i)
		}

		for _, tree := range trees {
			t := strings.Split(strings.TrimSpace(tree), ".")
			root := toByte(t[0])
			left := toByte(t[1])
			right := toByte(t[2])
			linkMap[root] = [2]byte{left, right}

			if root != 0 {
				treeMap[root] = &BinaryTree{value: root}
			}
			if treeMap[left] == nil && left != 0 {
				treeMap[left] = &BinaryTree{value: left}
			}
			if treeMap[right] == nil && right != 0 {
				treeMap[right] = &BinaryTree{value: right}
			}
		}

	}

	for k, _ := range linkMap {
		leftIndex := linkMap[k][0]
		rightIndex := linkMap[k][1]
		treeMap[k].left = treeMap[leftIndex]
		treeMap[k].right = treeMap[rightIndex]
	}

	return treeMap[rootKey]
}
