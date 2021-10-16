func longestPalindrome(s string) string {
  n := len(s)
  dp := make([][]bool, n)
  ans := s[0:1]
  // base case
  for i := range dp {
    dp[i] = make([]bool, n)
    dp[i][i] = true
    if i < len(s)-1 && s[i] == s[i+1] {
      // size = 2
      dp[i][i+1] = true
      ans = s[i:i+2]
    }
  }
  
  for size := 3; size <= n; size++ {
    left := 0
    right := left + size - 1
    for right < n {
      dp[left][right] = s[left] == s[right] && dp[left+1][right-1]
      if dp[left][right] && right - left + 1 > len(ans) {
        ans = s[left:right+1]
      }
      right++
      left++
    }
  }
  return ans
}