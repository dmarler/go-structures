package main

import (
	"fmt"

	"github.com/dmarler/go-structures/list"
)

func main() {
	ll1 := list.Slinkedlist{Head: nil, Tail: nil}

	list.InsertAtHead(&ll1, "Head of linked list 1.")
	list.InsertAtEnd(&ll1, "Link 2.")

	for i := 0; i < 100; i++ {
		list.InsertAtEnd(&ll1, "Link "+fmt.Sprintf("%d", i))
	}

	fmt.Println(list.GetLinkAtIndex(&ll1, 17).Value)
	fmt.Println(list.GetLinkAtIndex(&ll1, 3).Value)
	fmt.Println(list.GetLinkAtIndex(&ll1, 70).Value)
	fmt.Println(list.GetLinkAtIndex(&ll1, 53).Value)
}
