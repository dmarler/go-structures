package tests

import (
	"testing"

	"github.com/dmarler/go-structures/list"
)

func TestInsertAtHead(t *testing.T) {
	value := "Testing Head"

	llTest := list.Slinkedlist{Head: nil, Tail: nil}

	list.InsertAtEnd(&llTest, value)

	expected := value

	if expected != llTest.Head.Value {
		t.Errorf("Expected Value: %s, Actual Value: %s", expected, llTest.Head.Value)
	}
}
