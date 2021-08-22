/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums, 0, []int{}, used, &res)
	return res
}

func backtrack(nums []int, i int, cur []int, used []bool, res *[][]int) {
	if len(cur) == len(nums) {
		cp := make([]int, len(cur))
		copy(cp, cur)
		*res = append(*res, cp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] { continue }
		if i > 0 && nums[i] == nums[i-1] && used[i-1]  {
			continue
		}
		cur = append(cur, nums[i])
		used[i] = true
		backtrack(nums, i+1, cur, used, res)
		used[i] = false
		cur = cur[:len(cur)-1]
	}
}
// @lc code=end

