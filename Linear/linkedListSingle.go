package main

import "fmt"

func main() {

	l := LinkedList{}

	l.insertAfter(l.tail, node{value: 1})
	l.insertAfter(l.tail, node{value: 5})
	l.insertAfter(l.head, node{value: 3})
	l.traverse(print)
	// OUTPUT:
	// 1 false
	// 3 false
	// 5 true

	l.insertAfter(l.head.next, node{value: 4})
	l.insertAfter(l.head, node{value: 2})
	l.traverse(print)
	// OUTPUT:
	// 1 false
	// 2 false
	// 3 false
	// 4 false
	// 5 true

	fmt.Println(l.prev(l.head))
	fmt.Println(l.prev(l.head.next).value)
	fmt.Println(l.prev(l.tail).value)
	// OUTPUT:
	// <nil>
	// 1
	// 4

	l.remove(l.head.next.next)
	l.remove(l.head)
	l.remove(l.tail)
	l.traverse(print)
	// OUTPUT:
	// 2 false
	// 4 true

	l.insertBefore(l.tail, node{value: 3})
	l.traverse(print)
	// OUTPUT:
	// 2 false
	// 3 false
	// 4 true
}

func print(n *node) {
	fmt.Println(n.value, n.next == nil)
}

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (l *LinkedList) insertAfter(n *node, new node) {
	if l.head == nil {
		l.head = &new
		l.tail = &new
	} else {
		if n == l.tail {
			l.tail = &new
		}
		new.next = n.next
		n.next = &new
	}
}

func (l *LinkedList) insertBefore(n *node, new node) {

	switch {
	case n == l.head:
		if l.head == nil {
			l.head = &new
			l.tail = &new
		} else {
			l.head = &new
			l.head.next = n
		}
	default:
		prev := l.prev(n)
		prev.next = &new
		new.next = n
	}
}

func (l *LinkedList) remove(n *node) {
	switch {
	case n == nil || l.head == nil:
	case n == l.head:
		l.head = n.next
		n.next = nil
	case n == l.tail:
		prev := l.prev(n)
		l.tail = prev
		l.tail.next = nil
	default:
		prev := l.prev(n)
		prev.next = n.next
	}
}

func (l *LinkedList) traverse(f func(*node)) {
	for n := l.head; n != nil; n = n.next {
		f(n)
	}
}

func (l *LinkedList) prev(n *node) *node {
	if l.head == nil || n == l.head {
		return nil
	}

	var prev *node
	for i := l.head; i != nil; i = i.next {
		if i == n {
			return prev
		}
		prev = i
	}

	return nil
}
