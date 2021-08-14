/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// 1. sort s and t, and compare them. Time: O(nlogn), Space: O(1)
// 2. count the char in s and store in the map, then compare with t. Time: O(n+m), Space: O(n)

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cmap := make(map[rune]int)
	for _, c := range s {
		cmap[c]++
	}

	for _, c := range t {
		cmap[c]--
		if cmap[c] < 0 {
			return false
		}
	}
  return true
}
// @lc code=end

