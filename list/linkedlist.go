package list

type link struct {
	Value string
	Next  *link
}

type Slinkedlist struct {
	Head *link
	Tail *link
}

func InsertAtEnd(ll *Slinkedlist, value string) {
	l := &link{Value: value, Next: nil}

	if ll.Head == nil {
		ll.Head = l
	}
	if ll.Tail == nil {
		ll.Tail = l
	} else {
		ll.Tail.Next = &link{Value: value, Next: nil}
		ll.Tail = ll.Tail.Next
	}
}

func InsertAtHead(ll *Slinkedlist, value string) {
	l := &link{Value: value, Next: nil}

	if ll.Tail == nil {
		ll.Tail = l
	}
	if ll.Head == nil {
		ll.Head = l
	} else {
		ll.Head = &link{Value: value, Next: ll.Head}
	}
}

func GetLinkAtIndex(ll *Slinkedlist, index int) *link {
	counter := 0
	current := ll.Head

	if current == nil {
		return nil
	}

	for counter < index && current != nil {
		counter += 1
		current = current.Next
	}

	return current
}
