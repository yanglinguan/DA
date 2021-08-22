/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, 0, []int{}, &res)
	return res
}

func backtrack(nums []int, start int, cur []int, res *[][]int) {
	if start == len(nums) {
		cp := make([]int, len(cur))
		copy(cp, cur)
		*res = append(*res, cp)
		return
	}

	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		nums[i], nums[start] = nums[start], nums[i]
		backtrack(nums, start+1, cur, res)
		nums[i], nums[start] = nums[start], nums[i]
		cur = cur[:len(cur)-1]
	}
}
// @lc code=end

