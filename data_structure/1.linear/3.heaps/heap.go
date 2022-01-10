package main

import (
	"container/heap"
	"fmt"
)

//IntegerHeap a type
type IntegerHeap []int

//IntegerHeap a method - get the length of IntegerHeap
func (iHeap IntegerHeap) Len() int {
	return len(iHeap)
}

//IntegerHeap a method - checks if the element of i index is less than of j index
func (iHeap IntegerHeap) Less(i, j int) bool {
	return iHeap[i] < iHeap[j]
}

//IntegerHeap a method - swaps the element of i with j
func (iHeap IntegerHeap) Swap(i, j int) {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

//IntegerHeap method - pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

//IntegerHeap method - pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

// main method 
func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 3, 5}

	heap.Init(intHeap)
	heap.Push(intHeap, 4)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(intHeap))
	}
}
