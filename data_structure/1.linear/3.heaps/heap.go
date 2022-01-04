package main

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

//IntegerHeap method -pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

func main()  {

}
