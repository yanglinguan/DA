func reverseOnlyLetters(s string) string {
  rList := []rune(s)
  l, r := 0, len(s)-1
  for l < r {
    for l < r && !unicode.IsLetter(rList[l]) {
      l++
    }
    for l < r && !unicode.IsLetter(rList[r]) {
      r--
    }
    
		rList[l], rList[r] = rList[r], rList[l]
		l++
		r--
  }
  return string(rList)
}