/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / helper(x, -n)
	}
	return helper(x, n)
}

func helper(x float64, n int) float64 {
	if n == 0 { return 1 }
	tmp := helper(x, n/2)
	if n % 2 == 0 {
		return tmp * tmp
	} else {
		return tmp * tmp * x
	}
}
// @lc code=end

