/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// ideas:
// 1. find each subset of size 2, two loops. Time complexity: O(n^2). Space: O(1)
// 2. use map to store required numbers (target - num) and the index. Time complexity: O(n), Space: O(n)

// @lc code=start
func twoSum(nums []int, target int) []int {
	// store the required number, the value is idx
	need := make(map[int]int)

	for i, n := range nums {
		if _, exist := need[n]; exist {
			return []int{need[n], i}
		}
		need[target - n] = i
	}
  return []int{}
}
// @lc code=end

