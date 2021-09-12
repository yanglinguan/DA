func maxSumSubmatrix(matrix [][]int, k int) int {
	n := len(matrix)
	m := len(matrix[0])
	ans := math.MinInt32
	for left := 0; left < n; left++ {
		sum := make([]int, m)
		for right := left; right < n; right++ {
			for i := 0; i < m; i++ {
				sum[i] += matrix[right][i]
			}
			
			tmp := []int{0}
			curSum := 0 
			for _, s := range sum {
				curSum += s
				idx := findFirstGreat(tmp, curSum-k)
				if idx != -1 {
						ans = max(ans, curSum - tmp[idx])
				}
				insert(&tmp, curSum)
			}
		}
	}
	return ans
}

func findFirstGreat(nums []int, target int) int {
	l := 0
	r := len(nums) - 1 
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
					return mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func insert(nums *[]int, n int) {
	idx := findFirstGreat(*nums, n)
	if idx == -1 {
		*nums = append(*nums, n)
		return
	}
	
	*nums = append(*nums, 0)
	copy((*nums)[idx+1:], (*nums)[idx:])
	(*nums)[idx] = n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
