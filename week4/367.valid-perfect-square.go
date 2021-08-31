/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 { return true }
	l := 0
	r := num/2
	for l <= r {
		mid := l + (r -l) / 2
		tmp := mid * mid
		if tmp == num {
			return true
		} else if tmp < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
// @lc code=end

