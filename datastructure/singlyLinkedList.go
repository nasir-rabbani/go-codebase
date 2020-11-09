package main

import "fmt"

type node struct {
	next *node
	val  int
}

type list struct {
	head *node
}

func (ll *list) insert(key int) {
	newNode := &node{
		val: key,
	}
	if ll.head != nil {
		h := ll.head
		for h.next != nil {
			h = h.next
		}
		h.next = newNode
	} else {
		ll.head = newNode
	}

}

func (ll *list) disp() {
	head := ll.head

	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}

	fmt.Println("Done")
}

func main() {
	fmt.Println("Singly Linked List Impl")

	list := list{}
	list.insert(1)
	list.insert(2)
	list.insert(3)

	list.disp()
}
