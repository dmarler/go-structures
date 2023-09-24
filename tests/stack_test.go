package tests

import (
	"testing"

	"github.com/dmarler/go-structures/stack"
)

func TestPeek(t *testing.T) {
	s := stack.NewStack()

	stack.Push(s, "Insert 1")
	stack.Push(s, "Insert 2")
	stack.Push(s, "Insert 3")
	stack.Push(s, "Insert 4")
	stack.Push(s, "Insert 5")

	expected := "Insert 5"
	actual := stack.Peek(s)

	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = stack.Peek(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = stack.Peek(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = stack.Peek(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = stack.Peek(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}
}

func TestPop(t *testing.T) {
	s := stack.NewStack()

	stack.Push(s, "Insert 1")
	stack.Push(s, "Insert 2")
	stack.Push(s, "Insert 3")
	stack.Push(s, "Insert 4")
	stack.Push(s, "Insert 5")

	expected := "Insert 5"
	actual := stack.Pop(s)

	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 4"
	actual = stack.Pop(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 3"
	actual = stack.Pop(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 2"
	actual = stack.Pop(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 1"
	actual = stack.Pop(s)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}
}
