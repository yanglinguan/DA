func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i - stack[len(stack)-1])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParenthesesDP(s string) int {
	ans := 0
	dp := make([]int, len(s))
	
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i - dp[i-1] > 1 {
						dp[i] += dp[i-dp[i-1]-2]
				}
			}
			if ans < dp[i] {
				ans = dp[i]
			}
		}
	}
	
	return ans
}