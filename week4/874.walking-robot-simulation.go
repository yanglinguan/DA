/*
 * @lc app=leetcode id=874 lang=golang
 *
 * [874] Walking Robot Simulation
 */

// @lc code=start

// north, east, south, west
var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
func robotSim(commands []int, obstacles [][]int) int {
	obMap := map[[2]int]bool{}
	for _, o := range obstacles {
		obMap[[2]int{o[0], o[1]}] = true
	}
	curDir := 0
	res := 0
	x, y := 0, 0
	for _, c := range commands {
		if c == -1 {
			curDir = (curDir + 1) % 4
		} else if c == -2 {
			curDir = (curDir - 1 + 4) % 4
		} else {
			for ; c > 0; c-- {
				d := dirs[curDir]
				if !obMap[[2]int{x+d[0], y+d[1]}] {
					x += d[0]
					y += d[1]
				}
			}
			res = max(res, x*x+y*y)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
// @lc code=end

