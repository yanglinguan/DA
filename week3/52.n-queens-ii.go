/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
func totalNQueens(n int) int {
	res := 0
	backtrack(n, 0, map[int]bool{}, map[int]bool{}, map[int]bool{}, &res)
	return res
}

func backtrack(n int, i int, col, diff, sum map[int]bool, res *int) {
	if i == n {
		*res++
		return
	}

	for j := 0; j < n; j++ {
		if col[j] || diff[i-j] || sum[i+j] {
			continue
		}
		col[j] = true
		diff[i-j] = true
		sum[i+j] = true
		backtrack(n, i+1, col, diff, sum, res)
		col[j] = false
		diff[i-j] = false
		sum[i+j] = false
	}
}
// @lc code=end

