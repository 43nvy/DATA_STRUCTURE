package doubly

import "fmt"

type node struct {
	data int

	prev *node
	next *node
}

type LinkedList struct {
	head *node
	tail *node
}

func New() *LinkedList {
	return &LinkedList{}
}

func newNode(data int, prev, next *node) *node {
	return &node{
		data: data,

		prev: prev,
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

	var counter int
	current := ll.head

	for current != nil {
		current = current.next
		counter++
	}

	return counter
}

func (ll *LinkedList) AddFirst(data int) {
	newHead := newNode(data, nil, ll.head)
	ll.head.prev = newHead
	ll.head = newHead
}

func (ll *LinkedList) Append(data int) {
	newTail := newNode(data, ll.tail, nil)

	if ll.isEmpty() {
		ll.head = newTail
		ll.tail = newTail
	}

	ll.tail.next = newTail
	ll.tail = newTail
}

func (ll *LinkedList) AddBeforeN(n, data int) {
	current := ll.head
	index := 0

	for index != n {
		if current.next == nil {
			panic("out of range")
		}

		current = current.next
		index++
	}

	newNode := newNode(data, current, current.next)
	current.next.prev = newNode
	current.next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.head

	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.next
	}

	fmt.Println("nil")
}
