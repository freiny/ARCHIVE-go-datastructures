package bintree

import (
	"strconv"
	"strings"
)

// New creates and initializes a new binary tree
func New(value ...interface{}) *Tree {
	return new(Tree).Init(value...)
}

// Tree is a binary tree
type Tree struct {
	Value       interface{}
	left, right *Tree
}

// Init initializes and clears tree
func (t *Tree) Init(value ...interface{}) *Tree {
	t.Value = nil
	if len(value) != 0 {
		t.Value = value[0]
	}
	// t.Value = nil
	t.left = nil
	t.right = nil
	return t
}

// Left returns left child of tree
func (t *Tree) Left() *Tree {
	return t.left
}

// Right returns right child of tree
func (t *Tree) Right() *Tree {
	return t.right
}

// insertLeft inserts new tree between parent node and left child
// returns inserted tree node or returns nil on failed insert
func (t *Tree) insertLeft(new *Tree) *Tree {
	if t.left == nil {
		t.left = new
	} else {
		new.left = t.left
		t.left = new
	}
	return new
}

// InsertLeft inserts new tree with "value" between parent node and left child
// returns inserted tree node or returns nil on failed insert
func (t *Tree) InsertLeft(value interface{}) *Tree {
	return t.insertLeft(&Tree{Value: value})
}

// insertRight inserts new tree between parent node and right child
// returns inserted tree node or returns nil on failed insert
func (t *Tree) insertRight(new *Tree) *Tree {
	if t.right == nil {
		t.right = new
	} else {
		new.right = t.right
		t.right = new
	}
	return new
}

// InsertRight inserts new tree with "value" between parent node and left child
// returns inserted tree node or returns nil on failed insert
func (t *Tree) InsertRight(value interface{}) *Tree {
	return t.insertRight(&Tree{Value: value})
}

func (t *Tree) delete() {
}

func (t *Tree) PreOrder(f func(*Tree)) {
	t.order(f, "pre")
}

func (t *Tree) InOrder(f func(*Tree)) {
	t.order(f, "in")
}

func (t *Tree) PostOrder(f func(*Tree)) {
	t.order(f, "post")
}

func (t *Tree) order(f func(*Tree), orderType string) {
	var traverse func(bt *Tree)
	traverse = func(bt *Tree) {
		if bt != nil {
			if orderType == "pre" {
				f(bt)
			}
			traverse(bt.left)
			if orderType == "in" {
				f(bt)
			}
			traverse(bt.right)
			if orderType == "post" {
				f(bt)
			}
		}
	}
	traverse(t)
}

func (t *Tree) LevelOrder(f func(*Tree), onLevelChange ...func()) {
	order := map[int][]*Tree{}
	order[0] = []*Tree{}

	var traverse func(bt *Tree, level int, order map[int][]*Tree)
	traverse = func(bt *Tree, level int, order map[int][]*Tree) {
		if bt == nil {
			return
		}
		order[level] = append(order[level], bt)
		traverse(bt.left, level+1, order)
		traverse(bt.right, level+1, order)
	}
	traverse(t, 0, order)

	for i := 0; order[i] != nil; i++ {
		if len(onLevelChange) > 0 {
			onLevelChange[0]()
		}
		for _, node := range order[i] {
			f(node)
		}
	}
}

// BuildTreeInt builds a binary tree structure with int values from a string representation
func BuildTreeInt(s string) *Tree {

	linkMap := map[int][2]int{}
	treeMap := map[int]*Tree{}

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
				treeMap[root] = &Tree{Value: root}
			}
			if treeMap[left] == nil && left != 0 {
				treeMap[left] = &Tree{Value: left}
			}
			if treeMap[right] == nil && right != 0 {
				treeMap[right] = &Tree{Value: right}
			}
		}

	}

	for k := range linkMap {
		leftIndex := linkMap[k][0]
		rightIndex := linkMap[k][1]
		treeMap[k].left = treeMap[leftIndex]
		treeMap[k].right = treeMap[rightIndex]
	}

	return treeMap[rootKey]
}
