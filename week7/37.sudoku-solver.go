func solveSudoku(board [][]byte)  {
	rows := make([][]int, 9)
	cols := make([][]int, 9)
	for i := range rows {
			rows[i] = make([]int, 9)
			cols[i] = make([]int, 9)
	}
	blocks := make([][][]int, 3)
	for i := range blocks {
			blocks[i] = make([][]int, 3)
			for j := range blocks[i] {
					blocks[i][j] = make([]int, 9)
			}
	}
	
	for i, row := range board {
			curRow := rows[i]
			for j, c := range row {
					if c == '.' {
							continue
					}
					curCol := cols[j]
					curBlock := blocks[i/3][j/3]
					n := int(c - '0') - 1
					curRow[n] = 1
					curCol[n] = 1
					curBlock[n] = 1
			}
	}
	solve(board, rows, cols, blocks)
}

func solve(board [][]byte, rows [][]int, cols [][]int, blocks [][][]int) bool {
	for i := range board {
			curRow := rows[i]
			for j := range board[i] {
					if board[i][j] != '.' {
							continue
					}
					curCol := cols[j]
					curBlock := blocks[i/3][j/3]
					for n := 0; n < 9; n++ {
							if curRow[n] == 1 || curCol[n] == 1 || curBlock[n] == 1 {
									continue
							}
							board[i][j] = byte(n+1+'0')
							curRow[n] = 1
							curCol[n] = 1
							curBlock[n] = 1
							if solve(board, rows, cols, blocks) {
								 return true 
							}
							curRow[n] = 0
							curCol[n] = 0
							curBlock[n] = 0
							board[i][j] = '.'
					}
					return false
			}
	}
	return true
}