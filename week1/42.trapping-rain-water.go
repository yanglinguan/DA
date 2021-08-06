/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
  l := 0
	r := len(height) - 1
	lmax := 0
	rmax := 0
	res := 0
	for l < r {
		lmax = max(height[l], lmax)
		rmax = max(height[r], rmax)
		
		if lmax < rmax {
			res += lmax - height[l]
			l++
		} else {
			res += rmax - height[r]
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
// @lc code=end

