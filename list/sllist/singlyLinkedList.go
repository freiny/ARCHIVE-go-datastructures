package sllist

import "fmt"

// New returns a new list
func New() *List {
	return new(List).Init()
}

// Print prints and element
func Print(e *Element) {
	fmt.Print(e.Value, " ")
}

// Element is an item in a List
type Element struct {
	Value interface{}
	next  *Element
	list  *List
}

// Next returns the next element in the list
func (e *Element) Next() *Element {
	return e.next
}

// List is linked list
type List struct {
	root Element
	tail *Element
	len  int
}

// Hey temp func
func (l *List) Hey() {
	fmt.Println("HEY")
}

// Init initializes list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.tail = &l.root
	l.len = 0
	return l
}

// Head returns first element in list
func (l *List) Head() *Element {
	if l.root.next == &l.root {
		return nil
	}
	return l.root.next
}

// Tail returns last element in list
func (l *List) Tail() *Element {
	if l.tail == &l.root {
		return nil
	}
	return l.tail
}

// insert inserts new after e, increment list length, returns inserted Element
func (l *List) insert(new, e *Element) *Element {
	switch {
	case l.Head() == nil:
		l.root.next = new
		new.next = &l.root
		l.tail = new
	default:
		new.next = e.next
		e.next = new
		if e == l.tail {
			l.tail = new
		}
	}

	l.len++
	return new
}

// insertValue inserts a new Element with value after e Element
func (l *List) insertValue(value interface{}, e *Element) *Element {
	return l.insert(&Element{Value: value}, e)
}

// InsertAfter inserts new Element with value after e Element
func (l *List) InsertAfter(value interface{}, e *Element) *Element {
	return l.insertValue(value, e)
}

// Append appends new Element with value at end of list
func (l *List) Append(value interface{}) *Element {
	return l.insertValue(value, l.Tail())
}

// Traverse iterates over list elements executing all functions fs for each Element
func (l *List) Traverse(fs ...func(*Element)) {
	for i := l.Head(); i != &l.root; i = i.Next() {
		for _, f := range fs {
			f(i)
		}
	}
}
