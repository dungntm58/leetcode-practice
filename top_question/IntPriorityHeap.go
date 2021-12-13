package top_question

type Item struct {
	Val     int
	Prority int
}

type IntPriorityHeap struct {
	values []Item
	isMax  bool
}

func NewIntPriorityHeap(isMax bool) *IntPriorityHeap {
	return &IntPriorityHeap{isMax: isMax}
}

func (h IntPriorityHeap) Len() int {
	return len(h.values)
}

func (h IntPriorityHeap) Less(i, j int) bool {
	if h.isMax {
		return h.values[i].Prority > h.values[j].Prority
	}
	return h.values[i].Prority < h.values[j].Prority
}

func (h IntPriorityHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *IntPriorityHeap) Push(obj interface{}) {
	if v, ok := obj.(Item); ok {
		h.values = append(h.values, v)
	}
}

func (h *IntPriorityHeap) Pop() interface{} {
	n := len(h.values)
	lastVal := h.values[n-1]
	h.values = h.values[:n-1]
	return lastVal
}

func (h IntPriorityHeap) Peek() interface{} {
	return h.values[0]
}
