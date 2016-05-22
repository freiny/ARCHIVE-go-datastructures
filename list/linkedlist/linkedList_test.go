package linkedlist_test

import (
	"testing"

	"github.com/freiny/go-ds/list/linkedlist"
)

func TestAppend(t *testing.T) {
	l := linkedlist.New()
	l.Append(1)
	actual := l.Head().Value
	if actual != 1 {
		t.Errorf("Append(%v) = ")
	}
	l.Append(3)
	l.Append(5)
	l.Append(7)
}
