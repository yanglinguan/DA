/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// use map to store the sorted string as key 

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// key: sorted string, value: idx in the res array
	smap := make(map[string]int)
	res := make([][]string, 0)

	for _, s := range strs {
		ss := sortString(s)
		if _, exist := smap[ss]; !exist {
			smap[ss] = len(res)
			res = append(res, []string{})
		}
		res[smap[ss]] = append(res[smap[ss]], s)
	}

	return res
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
// @lc code=end

