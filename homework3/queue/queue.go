package queue

import "fmt"

var first *node
var last *node

type node struct {
	data string
	next *node
}

// Push enqueues item
func Push(item string) {
	n := &node{
		data: item,
	}
	if last == nil {
		first = n
		last = n
		return
	}
	last.next = n
	last = n
}

// Pop dequeues item
func Pop() string {
	if first == nil {
		return ""
	}
	item := first.data
	first = first.next
	if first == nil {
		last = nil
	}
	return item
}

// Run shows queue demo
func Run() {
	fmt.Println()
	fmt.Println("Queue")

	fmt.Println("Push: first")
	Push("first")

	fmt.Println("Push: second")
	Push("second")

	fmt.Println("Push: third")
	Push("third")

	item := Pop()
	fmt.Printf("Pop: %v\n", item)

	item = Pop()
	fmt.Printf("Pop: %v\n", item)

	item = Pop()
	fmt.Printf("Pop: %v\n", item)
}
