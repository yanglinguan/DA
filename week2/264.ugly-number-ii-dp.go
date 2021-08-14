/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */
// dp
// @lc code=start
func nthUglyNumber(n int) int {
		t2, t3, t5 := 0, 0, 0
		dp := make([]int, n)
		dp[0] = 1
		for i := 1; i < n; i++ {
			dp[i] = min(dp[t2]*2, dp[t3]*3, dp[t5]*5)
			if dp[i] == dp[t2]*2 { t2++ }
			if dp[i] == dp[t3]*3 { t3++ }
			if dp[i] == dp[t5]*5 { t5++ }
		}
		return dp[n-1]
}

func min(nums ...int) int {
	res := math.MaxInt32
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}
// @lc code=end

