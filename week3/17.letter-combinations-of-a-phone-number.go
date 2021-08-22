/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 { return []string{} }
	lmap := map[byte]string{}
	lmap['2'] = "abc"
	lmap['3'] = "def"
	lmap['4'] = "ghi"
	lmap['5'] = "jkl"
	lmap['6'] = "mno"
	lmap['7'] = "pqrs"
	lmap['8'] = "tuv"
	lmap['9'] = "wxyz"

	res := []string{}
	helper(digits, 0, "", &res, lmap)
	return res
}

func helper(digits string, i int, cur string, res *[]string, lmap map[byte]string) {
	if i == len(digits) {
		*res = append(*res, cur)
		return
	}
	letters := lmap[digits[i]]
	for _, c := range letters {
		cur += string(c)
		helper(digits, i+1, cur, res, lmap)
		cur = cur[:len(cur)-1]
	}
}
// @lc code=end

