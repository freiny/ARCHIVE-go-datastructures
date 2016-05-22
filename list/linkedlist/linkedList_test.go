package linkedlist_test

import (
	"testing"

	"github.com/freiny/go-ds/list/linkedlist"
	"github.com/freiny/go-util/ftest"
)

func TestNew(t *testing.T) {
	l := linkedlist.New()

	tests := []ftest.Test{
		{0, l.Head(), (*linkedlist.Element)(nil)},
		{1, l.Tail(), (*linkedlist.Element)(nil)},
	}
	ftest.Assert(t, "New()", tests)

}

func TestAppend(t *testing.T) {
	l := linkedlist.New()
	l.Append(1)
	tests := []ftest.Test{
		{0, l.Head().Value, 1},
		{1, l.Tail().Value, 1},
	}
	ftest.Assert(t, "Append(1)", tests)

	l.Init()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	tests = []ftest.Test{
		{0, l.Head().Value, 1},
		{1, l.Head().Next().Value, 2},
		{2, l.Tail().Value, 3},
	}
	ftest.Assert(t, "Append(1,2,3)", tests)
}

func TestInsertAfter(t *testing.T) {
	l := linkedlist.New()

	l.InsertAfter(1, nil)
	l.InsertAfter(5, l.Tail())
	e := l.InsertAfter(3, l.Head())
	l.InsertAfter(4, e)

	tests := []ftest.Test{

		{0, l.Head().Value, 1},
		{1, l.Head().Next().Value, 3},
		{2, l.Head().Next().Next().Value, 4},
		{3, l.Tail().Value, 5},
	}
	ftest.Assert(t, "InsertAfter(...1,3,4,5)", tests)

}

func TestTraverse(t *testing.T) {
	l := linkedlist.New()
	l.Append(3)
	l.Append(5)
	l.Append(7)

	f := func(e *linkedlist.Element) {
		var in interface{}
		in = e.Value.(int) + 1
		e.Value = in
	}
	l.Traverse(f)

	tests := []ftest.Test{
		{0, l.Head().Value, 4},
		{1, l.Head().Next().Value, 6},
		{2, l.Tail().Value, 8},
	}
	ftest.Assert(t, "Append(3,5,7)", tests)
}
