func numIslands(grid [][]byte) int {
  n := len(grid)
  m := len(grid[0])
  count := 0
  for i, row := range grid {
    for j, c := range row {
      if c == '1' {
        dfs(grid, i, j, n, m)
        count++
      }
    }
  }
  
  return count
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
func dfs(grid [][]byte, i, j, n, m int) {
  if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] != '1' {
    return
  }
  grid[i][j] = '0'
  
  for _, d := range dirs {
    dfs(grid, i+d[0], j+d[1], n, m)
  }
  
}