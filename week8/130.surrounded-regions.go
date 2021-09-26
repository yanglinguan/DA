func solve(board [][]byte)  {
  n := len(board)
  m := len(board[0])
  
  for col := 0; col < m; col++ {
    if board[0][col] == 'O' {
      dfs(board, 0, col, n, m)
    }
    if board[n-1][col] == 'O' {
      dfs(board, n-1, col, n, m)
    }
  }
  for row := 0; row < n; row++ {
    if board[row][0] == 'O' {
      dfs(board, row, 0, n, m)
    }
    if board[row][m-1] == 'O' {
      dfs(board, row, m-1, n, m)
    }
  }
  
  for _, row := range board {
    for j, c := range row {
      if c == '1' {
        row[j] = 'O'
      } else if c == 'O' {
        row[j] = 'X'
      }
    }
  }
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
func dfs(board [][]byte, i, j, n, m int) {
  if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != 'O' {
    return
  }
  
  board[i][j] = '1'
  for _, d := range dirs {
    dfs(board, i+d[0], j+d[1], n, m)
  }
}