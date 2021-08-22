/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, 0, []int{}, &res)
  return res
}

func backtrack(nums []int, i int, cur []int, res *[][]int) {
	if i == len(nums) {
		cp := make([]int, len(cur))
		copy(cp, cur)
		*res = append(*res, cp)
		return
	}
	// does not choose nums[i] into the subset
	backtrack(nums, i+1, cur, res)
	// choose nums[i]
	cur = append(cur, nums[i])
	backtrack(nums, i+1, cur, res)
	cur = cur[:len(cur)-1]
}
// @lc code=end

