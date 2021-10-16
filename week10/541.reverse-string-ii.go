func reverseStr(s string, k int) string {
  rList := []rune(s)
  
  for start := 0; start < len(s); start += 2*k {
    end := start + k
    if end > len(s) { end = len(s) }
    reverse(rList, start, end-1)
  }
  return string(rList)
}

func reverse(rList []rune, l, r int) {
  for l < r {
    rList[l], rList[r] = rList[r], rList[l]
    l++
    r--
  }
}