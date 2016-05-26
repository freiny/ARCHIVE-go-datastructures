package bintree

// New creates and initializes a new binary tree
func New() *Tree {
	return new(Tree).Init()
}

// Tree is a binary tree
type Tree struct {
	Value       interface{}
	left, right *Tree
}

// Init initializes and clears tree
func (t *Tree) Init() *Tree {
	t.Value = nil
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

// InsertLeft inserts new tree with "value" between parent node and left child
// returns inserted tree node or returns nil on failed insert
func (t *Tree) InsertLeft(value interface{}) *Tree {
	new := Tree{Value: value}
	if t.left == nil {
		t.left = &new
	} else {
		new.left = t.left
		t.left = &new
	}
	return &new
}

// InsertRight inserts new tree with "value" between parent node and left child
// returns inserted tree node or returns nil on failed insert
func (t *Tree) InsertRight(value interface{}) *Tree {
	new := Tree{Value: value}
	if t.right == nil {
		t.right = &new
	} else {
		new.right = t.right
		t.right = &new
	}
	return &new
}
