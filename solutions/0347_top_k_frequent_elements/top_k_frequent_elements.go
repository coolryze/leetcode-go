package topkfrequentelements

import "container/heap"

type entry struct {
	num   int
	count int
}

type minHeap []entry

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].count < h[j].count }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(entry)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	h := &minHeap{}
	heap.Init(h)

	for num, count := range countMap {
		heap.Push(h, entry{num, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(entry).num
	}

	return res
}
