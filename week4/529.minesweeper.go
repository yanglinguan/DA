/*
 * @lc app=leetcode id=529 lang=golang
 *
 * [529] Minesweeper
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		// game over
		board[click[0]][click[1]] = 'X'
		return board
	}
	n := len(board)
	m := len(board[0])
	dfs(board, click[0], click[1], n, m)
	return board
}

var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
func dfs(board [][]byte, i, j, n, m int) {
	if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != 'E' {
		return
	}
	mines := checkMine(board, i, j, n, m)
	if mines == 0 {
		board[i][j] = 'B'
		// continue search
		for _, d := range dirs {
			dfs(board, i+d[0], j+d[1], n, m)
		}
	} else {
		board[i][j] = byte(mines + '0')
		// stop search
	}
}

// find number of surrounding mines 
func checkMine(board [][]byte, i, j, n, m int) int {
	mine := 0
	for _, d := range dirs {
		di := i + d[0]
		dj := j + d[1]
		if di < 0 || di >= n || dj < 0 || dj >= m {
			continue
		}
		if board[di][dj] == 'M' {
			mine++
		}
	}
	return mine
}
// @lc code=end

