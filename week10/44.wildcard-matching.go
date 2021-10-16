func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	
	i := 0
	j := 0
	last_star := -1
	last_match := 0
	for i < n {
			if j < m && (s[i] == p[j] || p[j] == '?') {
					i++
					j++
			} else if j < m && p[j] == '*' {
					last_star = j
					last_match = i
					j++
			} else if last_star != -1 {
					j = last_star + 1
					last_match++
					i = last_match
			} else {
					return false
			}
	}
	
	for j < m && p[j] == '*' {
			j++
	}
	return j == m
}