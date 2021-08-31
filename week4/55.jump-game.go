/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	curReach := 0
	maxReach := 0
	for i := 0; i < len(nums) - 1; i++ {
		maxReach = max(maxReach, i + nums[i])
		if curReach == i {
			curReach = maxReach
		}
	}
	return curReach >= len(nums) - 1
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
// @lc code=end

