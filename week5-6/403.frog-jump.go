func canCross(stones []int) bool {
	n := len(stones)
	// dp[i]: a list of steps that can jump to stone i
	dp := make([]map[int]bool, n)
	for i := range dp {
			dp[i] = make(map[int]bool)
	}
	dp[0][0] = true
	
	for i := 1; i < n; i++ {
		for j := i-1; j >= 0; j-- {
			step := stones[i] - stones[j]
			if dp[j][step] || dp[j][step-1] || dp[j][step+1] {
				dp[i][step] = true
			}
		}
	}
	return len(dp[n-1]) > 0
}