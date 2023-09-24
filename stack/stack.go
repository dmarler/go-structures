package stack

import (
	"github.com/dmarler/go-structures/list"
)

type Stack struct {
	list *list.Slinkedlist
}

func Pop(stack *Stack) string {
	value := stack.list.Head.Value

	stack.list.Head = stack.list.Head.Next

	return value
}

func Peek(stack *Stack) string {
	return stack.list.Head.Value
}

func Push(stack *Stack, value string) {
	list.Prepend(stack.list, value)
}

func NewStack() *Stack {
	s := &Stack{&list.Slinkedlist{Head: nil, Tail: nil}}
	return s
}
