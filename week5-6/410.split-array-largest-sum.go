func splitArray(nums []int, m int) int {
	n := len(nums)
	
	// dp[i][j]: when the length of nums is i and split j parts
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	
	presum := getPresum(nums)
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], presum[i] - presum[k]))
			}
		}
	}
	
	return dp[n][m]
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func min(a, b int) int {
	if a > b { return b }
	return a
}

func getPresum(nums []int) []int {
	n := len(nums)
	presum := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	return presum
}