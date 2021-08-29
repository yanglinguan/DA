/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	queue := []string{start}
	visited := map[string]bool{}
	bankSet := map[string]bool{}
	for _, g := range bank {
		bankSet[g] = true
	}
	gen := []rune{'A', 'C', 'G', 'T'}
	step := 0
	for len(queue) > 0 {
		size := len(queue)

		// process current level
		for i := 0; i < size; i++ {
			s := queue[i]
			if s == end {
				return step
			}
			visited[s] = true
			for i := range s {
				cList := []rune(s)
				// find a mutation
				for _, g := range gen {
					if s[i] == byte(g) { continue }
					cList[i] = g
					newGen := string(cList)
					if bankSet[newGen] {
						queue = append(queue, newGen)
						delete(bankSet, newGen)
					}
				}
			}
		}
		step++
		queue = queue[size:]
	}
	return -1
}
// @lc code=end

