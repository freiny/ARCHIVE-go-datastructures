package dllist

// New returns a new Doubly Linked List
func New() *List {
	return new(List).Init()
}

// Element item in list
type Element struct {
	Value      interface{}
	prev, next *Element
}

// Prev returns previous element
func (e *Element) Prev() *Element {
	return e.prev
}

// Next returns next element
func (e *Element) Next() *Element {
	return e.next
}

// List is a list of elements
type List struct {
	root Element
	len  int
}

// Init initializes or clears the list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// Len returns list length
func (l *List) Len() int {
	return l.len
}

// First returns the first element of the list
func (l *List) First() *Element {
	return l.root.next
}

// Last returns the last element of the list
func (l *List) Last() *Element {
	return l.root.prev
}

// insert inserts new Element after e, then returns the inserted Element
func (l *List) insert(new *Element, e *Element) *Element {
	e.next.prev = new
	new.next = e.next
	new.prev = e
	e.next = new
	l.len++
	return new
}

// insert inserts new Element after e, then returns the inserted Element
func (l *List) insertValue(value interface{}, e *Element) *Element {
	return l.insert(&Element{Value: value}, e)
}

// Append appends Element with "value" to the end of the list
func (l *List) Append(value interface{}) *Element {
	return l.insertValue(value, l.root.prev)
}

// Prepend prepends Element with "value" to the front of the list
func (l *List) Prepend(value interface{}) *Element {
	return l.insertValue(value, &l.root)
}
