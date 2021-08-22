/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	majority := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
			count++
		} else if nums[i] == majority {
			count++
		} else {
			count--
		}
	}
	return majority
}
// @lc code=end

