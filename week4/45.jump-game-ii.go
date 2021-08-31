/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	curReach := 0
	maxReach := 0
	step := 0
	for i := 0; i < len(nums) - 1; i++ {
		maxReach = max(maxReach, i + nums[i])
		if i == curReach {
			step++
			curReach = maxReach
		}
	}
	return step
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
// @lc code=end

