/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, n := range nums {
		for _, s := range res {
			cp := make([]int, len(s))
			copy(cp, s)
			cp = append(cp, n)
			res = append(res, cp)
		}
	}
	return res
}
// @lc code=end

