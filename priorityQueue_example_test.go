package priorityQueue_test

import (
	"github.com/x1m3/priorityQueue"
	"fmt"
)

type Item int

func (i Item) HigherPriorityThan(other priorityQueue.Interface) bool {
	return i > other.(Item)
}

func ExampleSimple() {

	list := priorityQueue.New()
	list.Push(Item(1))
	list.Push(Item(10))
	list.Push(Item(5))
	list.Push(Item(7))
	list.Push(Item(2))
	list.Push(Item(3))
	list.Push(Item(4))
	list.Push(Item(6))
	list.Push(Item(8))
	list.Push(Item(9))

	for {
		r := list.Pop()
		if r==nil {
			break
		}
		fmt.Println(r)
	}
	// Output:
	// 10
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
}
