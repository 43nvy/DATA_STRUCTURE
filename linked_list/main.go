package main

import (
	"fmt"

	"github.com/43nvy/DATA_STRUCTURE/linked_list/doubly"
)

func main() {
	linkedList := doubly.New()
	fmt.Println(linkedList.Amount())
	for i := 1; i < 6; i++ {
		linkedList.Append(i)
	}

	linkedList.Display()
	fmt.Println(linkedList.Amount())
	linkedList.AddFirst(0)
	linkedList.Display()
	fmt.Println(linkedList.Amount())
	linkedList.AddBeforeN(3, 11)
	linkedList.Display()
	fmt.Println(linkedList.Amount())

}
