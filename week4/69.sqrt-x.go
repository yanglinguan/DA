/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	return int(sqrt2(float64(x)))
}

// math
func sqrt2(x float64) float64 {
	r := float64(x)
	for math.Abs(x / r - r) > 1e-9 {
		r = (r + x / r) / 2
	}
	return r
}

// binary search
func sqrt(x float64) float64 {
	l := float64(0)
	r := x
	for math.Abs(l - r) > 1e-6 {
		mid := l + (r - l) / 2
		tmp := x / mid
		if math.Abs(tmp - mid) <= 1e-6 {
			return mid
		} else if mid < tmp {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}
// @lc code=end

