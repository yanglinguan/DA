/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp[i][j]: the number of paths when reach (i, j)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsOneDimension(m int, n int) int {
	// instead of using matrix, we can use one dimesion array
	// updating this array
	dp := make([]int, n)
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

// @lc code=end

