/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
  i := 0
	for j, n := range nums {
		if n != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
// @lc code=end

