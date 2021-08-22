/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := [][]string{}
	backtrack(n, 0, make([]int, n), map[int]bool{}, map[int]bool{}, map[int]bool{}, &res)
	return res
}

func backtrack(n int, i int, queens []int, col, diff, sum map[int]bool, res *[][]string) {
	if i == n {
		solve := []string{}
		for _, q := range queens {
			row := ""
			for i := 0; i < n; i++ {
				if i == q {
					row += "Q"
				} else {
					row += "."
				}
			}
			solve = append(solve, row)
		}
		*res = append(*res, solve)
		return
	}
	// loop each col
	for j := 0; j < n; j++ {
		if col[j] || diff[i-j] || sum[i+j] {
			// we cannot put queen in j's column
			continue
		}
		col[j] = true
		diff[i-j] = true
		sum[i+j] = true
		queens[i] = j
		backtrack(n, i+1, queens, col, diff, sum, res)
		col[j] = false
		diff[i-j] = false
		sum[i+j] = false
	}
}
// @lc code=end

