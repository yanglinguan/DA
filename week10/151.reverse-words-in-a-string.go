func reverseWords(s string) string {
  rList := []byte{}
  left := 0
  right := len(s)-1
	// remove leading and trailing space
  for left < right && s[left] == ' ' {
    left++
  }
  for left < right && s[right] == ' ' {
    right--
  }
	// remove additional space in between
  for i := left; i <= right; i++ {
    if i > left && s[i] == ' ' && s[i-1] == ' ' {
      continue
    }
    rList = append(rList, s[i])
  }
  // reveser entire string
  reverse(rList, 0, len(rList)-1)
  
	// reverser each word
  for start := 0; start <= len(rList); {
    end := start
    for end < len(rList) && rList[end] != ' ' {
      end++
    }
    reverse(rList, start, end-1)
    start = end+1
  }
  return string(rList)
}

func reverse(rList []byte, l, r int) {
  for l < r {
    rList[l], rList[r] = rList[r], rList[l]
    l++
    r--
  }
}