/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */
// 1. sort
// 2. quick select
// @lc code=start
func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, len(nums)-k)
}

func quickSelect(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		p := partition(nums, l, r)
		if p == k {
			return nums[p]
		} else if p < k {
			l = p + 1
		} else {
			r = p - 1
		}
	}
	return nums[r]
}

func partition(nums []int, l, r int) int {
	pIdx := l + rand.Intn(r-l)
	// put pivot to the end
	nums[r], nums[pIdx] = nums[pIdx], nums[r]

	p := nums[r]
	// idx where [l, i) < pivot, [i, r] >= pivot 
	i := l

	for j := l; j < r; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// i is the postion of pivot
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
// @lc code=end

