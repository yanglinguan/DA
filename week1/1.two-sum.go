/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
  need := make(map[int]int)
	for i, n := range nums {
		if _, exist := need[n]; exist {
			return []int{need[n], i}
		}
		need[target-n] = i
	}
	return []int{}
}
// @lc code=end

