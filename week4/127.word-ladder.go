/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
// time
func ladderLength(beginWord string, endWord string, wordList []string) int {
	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{}
	visited := map[string]bool{}
	wordSet := map[string]bool{}
	for _, w := range wordList {
			wordSet[w] = true
	}
	if wordSet[endWord] {
			endSet[endWord] = true
	} else {
			return 0
	}
	step := 2
	for len(beginSet) > 0 {
		// stores next level strings
		tmp := map[string]bool{}
		for s := range beginSet {
			for i := range s {
				cList := []rune(s)
				// change the i's char 
				for c := 'a'; c <= 'z'; c++ {
					if s[i] == byte(c) { continue }
					cList[i] = c
					newS := string(cList)
					if endSet[newS] {
						return step
					}
					if !visited[newS] && wordSet[newS] {
						visited[newS] = true
						tmp[newS] = true
					}
				}
			}
		}
		step++
		// for the next iteration we choose the one with smaller size
		if len(tmp) < len(endSet) {
			beginSet = tmp
		} else {
			beginSet = endSet
			endSet = tmp
		}
	}

	return 0
}
// @lc code=end

