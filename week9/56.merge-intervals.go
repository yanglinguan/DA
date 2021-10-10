func merge(intervals [][]int) [][]int {
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  res := [][]int{intervals[0]}
  for i := 1; i < len(intervals); i++ {
    last := res[len(res)-1]
    if last[1] >= intervals[i][0] {
      res[len(res)-1][1] = max(last[1], intervals[i][1])
    } else {
      res = append(res, intervals[i])
    }
  }
  return res
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}