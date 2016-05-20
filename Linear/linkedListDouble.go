package main

import "fmt"

func main() {
	l := DLinkedList{}

	l.insertAfter(l.tail, node{value: 1})
	l.insertAfter(l.tail, node{value: 3})
	l.insertAfter(l.tail, node{value: 5})
	l.traverseNext(print)
	// OUTPUT:
	// 1 false
	// 3 false
	// 5 true

	l.insertBefore(l.head.next, node{value: 2})
	l.insertBefore(l.tail, node{value: 4})
	l.traverseNext(print)
	// OUTPUT:
	// 1 false
	// 2 false
	// 3 false
	// 4 false
	// 5 true

	l.remove(l.head.next.next)
	l.remove(l.head)
	l.remove(l.tail)
	l.traverseNext(print)
	// OUTPUT:
	// 2 true false
	// 4 false true

	l.traversePrev(print)
	// OUTPUT:
	// 4 false true
	// 2 true false
}

func print(n *node) {
	fmt.Println(n.value, n.prev == nil, n.next == nil)
}

type node struct {
	value int
	prev  *node
	next  *node
}

type DLinkedList struct {
	head *node
	tail *node
}

func (l *DLinkedList) traverseNext(f func(*node)) {
	for n := l.head; n != nil; n = n.next {
		f(n)
	}
}

func (l *DLinkedList) traversePrev(f func(*node)) {
	for n := l.tail; n != nil; n = n.prev {
		f(n)
	}
}

func (l *DLinkedList) insertBefore(n *node, new node) {
	switch {
	case l.head == nil:
		l.head = &new
		l.tail = &new
	case n == l.head:
		new.next = l.head
		l.head.prev = &new
		l.head = &new
	default:
		n.prev.next = &new
		new.prev = n.prev
		new.next = n
		n.prev = &new
	}
}

func (l *DLinkedList) insertAfter(n *node, new node) {
	switch {
	case l.head == nil:
		l.head = &new
		l.tail = &new
	case n == l.tail:
		n.next = &new
		new.prev = n
		l.tail = &new
	default:
		new.next = n.next
		n.next = &new
	}
}

func (l *DLinkedList) remove(n *node) {
	switch {
	case l.head == nil:
	case n == l.head:
		l.head = n.next
		l.head.prev = nil
	case n == l.tail:
		l.tail = n.prev
		l.tail.next = nil
	default:
		n.prev.next = n.next
		n.next.prev = n.prev
	}
}
