package queue

import (
	"github.com/dmarler/go-structures/list"
)

type Queue struct {
	list *list.Slinkedlist
}

func Pop(queue *Queue) string {
	value := queue.list.Head.Value

	queue.list.Head = queue.list.Head.Next

	return value
}

func Peek(queue *Queue) string {
	return queue.list.Head.Value
}

func Push(queue *Queue, value string) {
	list.Append(queue.list, value)
}

func NewQueue() *Queue {
	q := &Queue{&list.Slinkedlist{Head: nil, Tail: nil}}
	return q
}
