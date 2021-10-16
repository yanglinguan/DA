func findAnagrams(s string, p string) []int {
  countMap := map[byte]int{}
  for i := range p {
    countMap[p[i]]++
  }
  
  left, right := 0, 0
  count := len(countMap)
  result := []int{}
  for right < len(s) {
    c := s[right]
    countMap[c]--
    if countMap[c] == 0 {
      count--
    }
    right++
    
    for count == 0 {
      c := s[left]
      countMap[c]++
      if countMap[c] > 0 {
        count++
      }
      
      if right - left == len(p) {
        result = append(result, left)
      }
      left++
    }
  }
  return result
}