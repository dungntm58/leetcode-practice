package top_question

import "container/heap"

// Find Median from Data Stream
// The median is the middle value in an ordered integer list.
// If the size of the list is even, there is no middle value and the median is the mean of the two middle values.
// - For example, for arr = [2,3,4], the median is 3.
// - For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:
// - MedianFinder() initializes the MedianFinder object.
// - void addNum(int num) adds the integer num from the data stream to the data structure.
// - double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.

type MedianFinder struct {
	minHeap, maxHeap *IntHeap
}

func ConstructorMedianFinder() MedianFinder {
	return MedianFinder{
		minHeap: NewIntHeap(false),
		maxHeap: NewIntHeap(true),
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.minHeap, num)
	heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	if this.minHeap.Len() < this.maxHeap.Len() {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.Peek().(int))
	}
	min, max := this.minHeap.Peek().(int), this.maxHeap.Peek().(int)
	return (float64(min) + float64(max)) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
