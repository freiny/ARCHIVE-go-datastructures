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

	root.levelOrder(printLevelOrder)
	fmt.Print("\n\n")
	// OUTPUT:
	// 8.5.4
	// 5.9.7 4..11
	// 9.. 7.1.12 11.3.
	// 1.. 12.2. 3..
	// 2..

	root.preOrder(print)
	fmt.Print("\n\n")
	// OUTPUT:
	// 8 5 9 7 1 12 2 4 11 3

}

func print(t *BinaryTree) {
	if t != nil {
		fmt.Print(t.value, " ")
	}
}

func printLevelOrder(t *BinaryTree, level int, levelChange bool) {
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
	value       int
	left, right *BinaryTree
}

func (t *BinaryTree) insert(value int) {
	switch {
	case t.left == nil:
		t.left = &BinaryTree{value: value}
	case t.right == nil:
		t.right = &BinaryTree{value: value}
	default:
		t.left = &BinaryTree{value: value, left: t.left}
	}
}

func (t *BinaryTree) insertLeft(value int) {
	if t.left == nil {
		t.left = &BinaryTree{value: value}
	} else {
		t.left = &BinaryTree{value: value, left: t.left}
	}
}

func (t *BinaryTree) insertRight(value int) {
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
// PreOrder - 8 5 9 7 1 12 2 4 11 3
// InOrder - 9 5 1 7 2 12 8 4 3 11
// PostOrder - 9 1 2 12 7 5 3 11 4 8
// LevelOrder - 8 5 4 9 7 11 1 12 3 2

func (t *BinaryTree) preOrder(f func(*BinaryTree)) {
	var traverse func(bt *BinaryTree)
	traverse = func(bt *BinaryTree) {
		if bt != nil {
			f(bt)
			traverse(bt.left)
			traverse(bt.right)
		}
	}

	traverse(t)
}

func (t *BinaryTree) inOrder() {
}
func (t *BinaryTree) postOrder() {
}

func (t *BinaryTree) levelOrder(f func(*BinaryTree, int, bool)) {
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

	linkMap := map[int][2]int{}
	treeMap := map[int]*BinaryTree{}

	rows := strings.Split(strings.TrimSpace(s), "\n")
	rootKey, _ := strconv.Atoi(strings.Split(strings.TrimSpace(rows[0]), ".")[0])
	for _, row := range rows {
		trees := strings.Split(row, " ")

		for _, tree := range trees {
			t := strings.Split(strings.TrimSpace(tree), ".")
			root, _ := strconv.Atoi(t[0])
			left, _ := strconv.Atoi(t[1])
			right, _ := strconv.Atoi(t[2])
			linkMap[root] = [2]int{left, right}

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
