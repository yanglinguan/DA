/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			count++
			i++
			j++
		} else {
			j++
		}
	}
	return count
}
// @lc code=end

