package main

import (
	"container/heap"
	"fmt"
)

// Modified from: https://golang.org/pkg/container/heap/

type ID int

// An Item is something we manage in a priority queue.
type Item struct {
	ID           // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// Priority returns the Item's priority value.
func (i *Item) Priority() int {
	return i.priority
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// AddPriority adds the input value to each item's priority.
func AddPriority(pq PriorityQueue, add int) {
	for _, item := range pq {
		item.priority += add
		heap.Fix(&pq, item.index)
	}
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	l := len(pq)
	fmt.Println("pq len:", l, "i:", i, "j:", j)
	fmt.Println("pq", pq)
	if l > 1 {
		return pq[i].priority > pq[j].priority
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, id ID, priority int) {
	item.ID = id
	item.priority = priority
	heap.Fix(pq, item.index)
}
