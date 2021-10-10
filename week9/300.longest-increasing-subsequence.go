func lengthOfLIS(nums []int) int {
  var lis []int
  for _, n := range nums {
    idx := findFirstGreat(lis, n)
    if idx == len(lis) {
      lis = append(lis, n)
    } else {
      lis[idx] = n
    }
  }
  return len(lis)
}

func findFirstGreat(nums []int, n int) int {
  l := 0
  r := len(nums) - 1
  for l <= r {
    mid := l + (r-l)/2
    if nums[mid] >= n {
      if mid == 0 || nums[mid-1] < n {
        return mid
      }
      r = mid - 1
    } else {
      l = mid + 1
    }
  }
  return len(nums)
}

