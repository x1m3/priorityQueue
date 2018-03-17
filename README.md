# Just Another implementation of a Golang Priority Queue

A priority queue is like a regular queue ordered by a priority value. An item with more priority will be served before, regardless the order of insertion.

https://en.wikipedia.org/wiki/Priority_queue

## Why another priority queue?

There are many implementations of priority queues. PQ are usually implemented with heaps and there is a good example in the golang documentation. https://golang.org/pkg/container/heap/#example__priorityQueue

The advantage of this version over many others is memory contention and a clean interface with 2 methods: Push and Pop.

Other implementations never shrink the slice used in the heap, so memory used is always equal to the max value reached at any moment. This can be very dangerous for a server that is running 24x7.

This PQ versions hrinks the memory when it detects that the capacity of the slice is very big for the amount of items stored at some point.

## How to use it?

1. import github.com/x1m3/priorityQueue
2. Implement the priorityQueue.Interface in the items you want to store on it.
3. Push() and Pop() items.

Here you can find a simple example.

https://github.com/x1m3/priorityQueue/blob/master/priorityQueue_example_test.go







