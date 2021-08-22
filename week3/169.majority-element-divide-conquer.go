/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	mid := l + (r-l)/2
	left := helper(nums, l, mid)
	right := helper(nums, mid+1, r)
	if left == right {
		return left
	}
	if count(nums, l, mid, left) > count(nums, mid+1, r, right) {
		return left
	}
	return right
}

func count(nums []int, l, r int, n int) int {
	c := 0
	for i := l; i <= r; i++ {
		if nums[i] == n {
			c++
		}
	}
	return c
}
// @lc code=end

