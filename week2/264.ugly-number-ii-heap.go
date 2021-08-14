/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// 1. priority queue: min heap, find next ugly number
// 2. dp

// @lc code=start
func nthUglyNumber(n int) int {
    mh := make(minHeap, 0)
		heap.Init(&mh)
		k := 0
		heap.Push(&mh, 1)
		for n > 0 {
			top := heap.Pop(&mh).(int)
			// remove duplicates
			if top == k { continue }
			k = top
			heap.Push(&mh, k * 2)
			heap.Push(&mh, k * 3)
			heap.Push(&mh, k * 5)
			n--
		}
		return k
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h  minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

// @lc code=end

