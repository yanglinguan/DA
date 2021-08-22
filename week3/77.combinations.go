/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	backtrack(n, 1, k, []int{}, &res)
	return res
}

func backtrack(n int, i int, k int, cur []int, res *[][]int) {
	if len(cur) == k {
		cp := make([]int, k)
		copy(cp, cur)
		*res = append(*res, cp)
		return
	}
	if i > n { return }
	// add i into the combination
	cur = append(cur, i)
	backtrack(n, i+1, k, cur, res)
	cur = cur[:len(cur)-1]
	// does not choose i in the combination
	backtrack(n, i+1, k, cur, res)
}
// @lc code=end

