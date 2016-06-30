package bintree_test

import (
	"testing"

	"github.com/freiny/go-datastructures/tree/bintree"
	"github.com/freiny/go-util/ftest"
)

func TestNew(t *testing.T) {
	bt := bintree.New()
	tests := []ftest.Test{
		{0, bt.Value, nil},
		{1, bt.Left(), (*bintree.Tree)(nil)},
		{2, bt.Right(), (*bintree.Tree)(nil)},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })

	bt = bintree.New("abc")
	tests = []ftest.Test{
		{3, bt.Value, "abc"},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}

func TestInsertLeft(t *testing.T) {
	bt := bintree.New(0)
	bt.InsertLeft(1)
	tests := []ftest.Test{
		{0, bt.Value, 0},
		{1, bt.Left().Value, 1},
		{2, bt.Right(), (*bintree.Tree)(nil)},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}

func TestInsertRight(t *testing.T) {
	bt := bintree.New(0)
	bt.InsertRight(1)
	tests := []ftest.Test{
		{0, bt.Value, 0},
		{1, bt.Right().Value, 1},
		{2, bt.Left(), (*bintree.Tree)(nil)},
	}
	ftest.Assert(t, tests, func(s string) { t.Errorf(s) })
}
