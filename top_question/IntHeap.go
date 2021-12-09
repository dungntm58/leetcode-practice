package top_question

type IntHeap struct {
	nums  []int
	isMax bool
}

func NewIntHeap(isMax bool) *IntHeap {
	return &IntHeap{isMax: isMax}
}

func (h IntHeap) Len() int {
	return len(h.nums)
}

func (h IntHeap) Less(i, j int) bool {
	if h.isMax {
		return h.nums[i] > h.nums[j]
	}
	return h.nums[i] < h.nums[j]
}

func (h IntHeap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *IntHeap) Push(obj interface{}) {
	if v, ok := obj.(int); ok {
		h.nums = append(h.nums, v)
	}
}

func (h *IntHeap) Pop() interface{} {
	n := len(h.nums)
	lastVal := h.nums[n-1]
	h.nums = h.nums[:n-1]
	return lastVal
}

func (h IntHeap) Peek() interface{} {
	return h.nums[0]
}
