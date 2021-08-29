/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap := map[string]bool{}
	for _, w := range wordList {
			wordMap[w] = true
	}
	
	if !wordMap[endWord] {
			return [][]string{}
	}
	
	delete(wordMap, beginWord)
	
	queue := [][]string{{beginWord}}
	step := 0
	res := [][]string{}
	// the result all be add at the same level
	// so if len(res) > 0, we find all the result
	for len(queue) > 0 && len(res) == 0 {
		size := len(queue)
		step++
		// record the visited word in the current level
		// since we need to find the all the min path
		// so we need a visited map at each level
		// otherwise we may miss pathes
		visited := map[string]bool{}
		for i := 0; i < size; i++ {
			path := queue[i]
			last := path[len(path)-1]
			for idx := range last {
				// find next strings with one char difference
				cList := []rune(last)
				for c := 'a'; c <= 'z'; c++ {
					if last[idx] == byte(c) {continue}
					cList[idx] = c
					newS := string(cList)
					if wordMap[newS] {
						newPath := make([]string, len(path))
						copy(newPath, path)
						newPath = append(newPath, newS)
						if newS == endWord {
								res = append(res, newPath)
						} else {
								queue = append(queue, newPath)   
								visited[newS] = true
						}
					}
				}
			}
		}
		queue = queue[size:]
		// remove visited word for next level
		for w := range visited {
			delete(wordMap, w)
		}
	}
	return res
}
// @lc code=end

