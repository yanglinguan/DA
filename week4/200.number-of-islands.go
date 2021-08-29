/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 { return 0 }
	m := len(grid[0])
	
	islands := 0
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				// mark all connected '1's to '0'
				dfs(grid, i, j, n, m)
				islands++
			}
		}
	}
	return islands
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

for dfs(grid [][]byte, i, j, n, m int) {
	if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	for _, d := range dirs {
		dfs(grid, i+d[0], j+d[1], n, m)
	}
}
// @lc code=end

