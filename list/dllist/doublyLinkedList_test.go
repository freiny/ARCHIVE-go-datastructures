package dllist_test

import (
	"testing"

	"github.com/freiny/go-ds/list/dllist"
	"github.com/freiny/go-util/ftest"
)

func TestNew(t *testing.T) {
	l := dllist.New()

	tests := []ftest.Test{
		{0, l.Len(), 0},
		{1, l.First().Value, nil},
		{1, l.Last().Value, nil},
	}

	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

}

func TestAppend(t *testing.T) {
	l := dllist.New()
	l.Append(3)
	tests := []ftest.Test{
		{0, l.Len(), 1},
		{1, l.First().Value, 3},
		{2, l.Last().Value, 3},
	}

	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

	l.Init()
	l.Append(1)
	tests = []ftest.Test{
		{0, l.Len(), 1},
		{1, l.First().Value, 1},
		{2, l.Last().Value, 1},
	}

	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

	l.Init()
	l.Append(3)
	l.Append(4)
	l.Append(5)
	tests = []ftest.Test{
		{0, l.Len(), 3},
		{1, l.First().Value, 3},
		{2, l.First().Next().Value, 4},
		{3, l.Last().Value, 5},
	}

	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

}

func TestPrepend(t *testing.T) {
	l := dllist.New()
	l.Prepend(3)
	tests := []ftest.Test{
		{0, l.First().Value, 3},
	}

	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

	l.Init()
	l.Prepend(5)
	l.Prepend(3)
	l.Prepend(1)
	tests = []ftest.Test{
		{0, l.First().Value, 1},
		{1, l.First().Next().Value, 3},
		{2, l.Last().Value, 5},
	}

	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}

func TestInsertBefore(t *testing.T) {
	l := dllist.New()
	l.InsertBefore(3, l.First())
	tests := []ftest.Test{
		{0, l.First().Value, 3},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

	l.Init()
	l.InsertBefore(5, l.First())
	l.InsertBefore(1, l.First())
	l.InsertBefore(3, l.Last())
	tests = []ftest.Test{
		{0, l.First().Value, 1},
		{1, l.First().Next().Value, 3},
		{2, l.Last().Value, 5},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}

func TestInsertAfter(t *testing.T) {
	l := dllist.New()
	l.InsertAfter(3, l.First())
	tests := []ftest.Test{
		{0, l.First().Value, 3},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

	l.Init()
	l.InsertAfter(1, l.Last())
	l.InsertAfter(5, l.Last())
	l.InsertAfter(3, l.First())
	tests = []ftest.Test{
		{0, l.First().Value, 1},
		{1, l.First().Next().Value, 3},
		{2, l.Last().Value, 5},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}

func TestTraverse(t *testing.T) {
	arr := make([]interface{}, 0, 0)
	f := func(e *dllist.Element) *dllist.Element {
		arr = append(arr, e.Value)
		return e
	}

	l := dllist.New()
	l.Traverse(f)
	tests := []ftest.Test{
		{0, arr, make([]interface{}, 0, 0)},
	}
	ftest.AssertDeep(t, tests, func(s string) { t.Errorf(s) })

	l.Init()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Traverse(f)
	tests = []ftest.Test{
		{0, l.Len(), 5},
		{1, arr, []interface{}{1, 2, 3, 4, 5}},
	}

	ftest.AssertDeep(t, tests, func(s string) { t.Errorf(s) })
}

func TestDelete(t *testing.T) {
	arr := make([]interface{}, 0, 0)
	f := func(e *dllist.Element) *dllist.Element {
		arr = append(arr, e.Value)
		return e
	}

	l := dllist.New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Delete(l.First().Next())
	l.Delete(l.First().Next().Next())
	l.Traverse(f)
	tests := []ftest.Test{
		{0, l.Len(), 3},
		{1, arr, []interface{}{1, 3, 5}},
	}

	ftest.AssertDeep(t, tests, func(s string) { t.Errorf(s) })
}
