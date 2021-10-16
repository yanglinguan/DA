func reverseWords(s string) string {
  rList := []rune(s)
  n := len(s)
  for start := 0; start < n; {
    end := start 
    for end < n && rList[end] != ' ' {
      end++
    }
    reverse(rList, start, end-1)
    start = end+1
  }
  return string(rList)
}

func reverse(rList []rune, l, r int) {
  for ; l < r; l, r = l+1, r-1 {
    rList[l], rList[r] = rList[r], rList[l]
  } 
}