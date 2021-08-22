/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	// initially, we can add n '(' and 0 ')'
	backtrack(n, 0, "", &res)
	return res
}

// left: number of '(' can be added
// right: number of ')' can be added
func backtrack(left, right int, cur string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
	}

	// add '('
	if left > 0 {
		// when add a '(', we can add ')'
		// therefore, right++
		backtrack(left-1, right+1, cur+"(", res)
	}
	// add ')'
	if right > 0 {
		backtrack(left, right-1, cur+")", res)
	}
}
// @lc code=end

