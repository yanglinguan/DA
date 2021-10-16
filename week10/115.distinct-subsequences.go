func numDistinct(s string, t string) int {
  n := len(s)
  m := len(t)
  dp := make([][]int, n+1)
  for i := range dp {
    dp[i] = make([]int, m+1)
  }
  
  for i := 0; i <= n; i++ {
    dp[i][0] = 1
  }
  
  for i := 1; i <= m; i++ {
    dp[0][i] = 0
  }
  
  for i := 1; i <= n; i++ {
    for j := 1; j <= m; j++ {
      if s[i-1] == t[j-1] {
        dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
      } else {
        dp[i][j] = dp[i-1][j]
      }
    } 
  }
  return dp[n][m]
}