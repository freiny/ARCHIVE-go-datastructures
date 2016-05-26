package sllist_test

import (
	"testing"

	"github.com/freiny/go-ds/list/sllist"
	"github.com/freiny/go-util/ftest"
)

func TestNew(t *testing.T) {
	l := sllist.New()

	tests := []ftest.Test{
		{0, l.Head(), (*sllist.Element)(nil)},
		{1, l.Tail(), (*sllist.Element)(nil)},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

}

func TestAppend(t *testing.T) {
	l := sllist.New()
	l.Append(1)
	tests := []ftest.Test{
		{0, l.Head().Value, 1},
		{1, l.Tail().Value, 1},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

	l.Init()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	tests = []ftest.Test{
		{0, l.Head().Value, 1},
		{1, l.Head().Next().Value, 2},
		{2, l.Tail().Value, 3},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}

func TestInsertAfter(t *testing.T) {
	l := sllist.New()

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
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

}

func TestTraverse(t *testing.T) {
	l := sllist.New()
	l.Append(3)
	l.Append(5)
	l.Append(7)

	f := func(e *sllist.Element) {
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
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}
