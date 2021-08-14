/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */
// 1. count frequency then sort by frequency 
// 2. count frequency, use a min heap of size k
// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// count frequency
  count := map[int]int{}
	for _, n := range nums {
		count[n]++
	}
	
	mh := make(minHeap, 0)
	heap.Init(&mh)
	for n, c := range count {
		heap.Push(&mh, []int{n, c})
		if mh.Len() > k {
			// if size > k, pop the least frequent one
			heap.Pop(&mh)
		}
	}
	res := make([]int, k)
	for i, n := range mh {
		res[i] = n[0]
	}
	return res
}
// implement Len, Less, Swap, Push, Pop
type minHeap [][]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.([]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}



// @lc code=end

