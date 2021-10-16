func isIsomorphic(s string, t string) bool {
  ms := make([]int, 256)
  mt := make([]int, 256)
  
  for i := range s {
    if ms[int(s[i])] != mt[int(t[i])] {
      return false
    }
    ms[int(s[i])] = i + 1
    mt[int(t[i])] = i + 1
  }
  
  return true
}
