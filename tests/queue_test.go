package tests

import (
	"testing"

	"github.com/dmarler/go-structures/queue"
)

func TestQueuePeek(t *testing.T) {

	q := queue.NewQueue()

	queue.Push(q, "Insert 1")
	queue.Push(q, "Insert 2")
	queue.Push(q, "Insert 3")
	queue.Push(q, "Insert 4")
	queue.Push(q, "Insert 5")

	expected := "Insert 1"
	actual := queue.Peek(q)

	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = queue.Peek(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = queue.Peek(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = queue.Peek(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	actual = queue.Peek(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}
}

func TestQueuePop(t *testing.T) {

	q := queue.NewQueue()

	queue.Push(q, "Insert 1")
	queue.Push(q, "Insert 2")
	queue.Push(q, "Insert 3")
	queue.Push(q, "Insert 4")
	queue.Push(q, "Insert 5")

	expected := "Insert 1"
	actual := queue.Pop(q)

	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 2"
	actual = queue.Pop(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 3"
	actual = queue.Pop(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 4"
	actual = queue.Pop(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}

	expected = "Insert 5"
	actual = queue.Pop(q)
	if expected != actual {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, actual)
	}
}
