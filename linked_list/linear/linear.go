package linear

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

func New() *LinkedList {
	return &LinkedList{}
}

func newNode(data int, next *node) *node {
	return &node{
		data: data,
		next: next,
	}
}

func (ll *LinkedList) isEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) Amount() int {
	if ll.isEmpty() {
		return 0
	}

	counter := 1
	current := ll.head

	for current.next != nil {
		current = current.next
		counter++
	}

	return counter
}

func (ll *LinkedList) AddFirst(data int) {
	newHead := newNode(data, ll.head)
	ll.head = newHead
}

func (ll *LinkedList) Append(data int) {
	newNode := newNode(data, nil)
	if ll.isEmpty() {
		ll.head = newNode
	}

	last := ll.head

	for last != nil {
		last = last.next
	}

	last = newNode
}

func (ll *LinkedList) AddBeforeN(n, data int) {
	current := ll.head
	index := 0

	for current != nil {
		if index == n {
			break
		}

		current = current.next
		index++
	}

	if index < n {
		panic("out of range")
	}

	newNode := newNode(data, current.next)
	current.next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println("nil")
}
