package list

type link struct {
	Value string
	Next  *link
}

type Slinkedlist struct {
	Head *link
	Tail *link
}

func Append(ll *Slinkedlist, value string) {
	l := &link{Value: value, Next: nil}

	if ll.Tail == nil {
		ll.Tail = l
		ll.Head = l
	} else {
		ll.Tail.Next = l
		ll.Tail = ll.Tail.Next
	}
}

func Prepend(ll *Slinkedlist, value string) {
	l := &link{Value: value, Next: ll.Head}

	if ll.Head == nil {
		ll.Tail = l
		ll.Head = l
	} else {
		ll.Head = l
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
