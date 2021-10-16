func myAtoi(s string) int {
  idx := 0
  // read leading ' '
  for idx < len(s) && s[idx] == ' ' {
    idx++
  }
  if idx == len(s) { return 0 }
  
  // check sign 
  sign := int64(1)
  if s[idx] == '+' {
    sign = 1
    idx++
  } else if s[idx] == '-' {
    sign = -1
    idx++
  } 
  // convert to int
  result := int64(0)
  for idx < len(s) && unicode.IsDigit(rune(s[idx])) {
    tmp := result * 10 + int64(s[idx]-'0')
    if tmp > math.MaxInt32 {
      if sign < 0 { 
        return math.MinInt32 
      }
      return math.MaxInt32
    }
    result = tmp
    idx++
  }
  
  return int(result * sign)
}