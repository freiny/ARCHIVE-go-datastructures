package linkedlist_test

import (
	"testing"

	"github.com/freiny/go-ds/list/linkedlist"
)

func TestNew(t *testing.T) {
	l := linkedlist.New()
	tests := []bool{
		l.Head() == nil,
		l.Tail() == nil,
	}
	msg := []string{
		"l.Head() == nil",
		"l.Tail() == nil",
	}
	for i, pass := range tests {
		if !pass {
			t.Errorf("Append(1) test failed: %q", msg[i])
		}
	}

}

func TestAppend(t *testing.T) {
	l := linkedlist.New()
	l.Append(1)
	tests := []bool{
		l.Head().Value == 1,
		l.Tail().Value == 1,
	}
	msg := []string{
		"l.Head().Value == 1",
		"l.Tail().Value == 1",
	}
	for i, pass := range tests {
		if !pass {
			t.Errorf("Append(1) test failed: %q", msg[i])
		}
	}

	l.Init()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	tests = []bool{
		l.Head().Value == 1,
		l.Head().Next().Value == 2,
		l.Tail().Value == 3,
	}
	msg = []string{
		"l.Head().Value == 1",
		"l.Head().Next().Value == 2",
		"l.Tail().Value == 3",
	}
	for i, pass := range tests {
		if !pass {
			t.Errorf("Append(1...3) test failed: %q", msg[i])
		}
	}

}
