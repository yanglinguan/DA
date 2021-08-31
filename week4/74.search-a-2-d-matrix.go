/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	l := 0
	r := m * n - 1

	for l <= r {
		mid := l + (r -l) / 2
		i := mid / m
		j := mid % m
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
// @lc code=end

