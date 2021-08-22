/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 3 { return n }
	pp := 1
	p := 2
	
	for i := 3; i <= n; i++ {
		cur := pp + p
		pp = p
		p = cur
	}
	return p
}
// @lc code=end

