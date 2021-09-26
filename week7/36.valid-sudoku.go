func isValidSudoku(board [][]byte) bool {
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
					if curRow[n] == 1 || curCol[n] == 1 || curBlock[n] == 1 {
							return false
					}
					curRow[n] = 1
					curCol[n] = 1
					curBlock[n] = 1
			}
	}
	return true
}