/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int)  {
  n := len(nums)
	k %= n
	count := 0
	for i := 0; count < n; i++ {
		start := i
		cur := i
		prev := nums[cur]
		for {
			next := (cur + k) % n
			tmp := nums[next]
			nums[next] = prev
			cur = next
			prev = tmp
			count++
			if cur == start { break }
		}
	}
}
// @lc code=end

