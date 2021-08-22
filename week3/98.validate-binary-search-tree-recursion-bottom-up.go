/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// bottom-up
func isValidBST(root *TreeNode) bool {
	v, _, _ := helper(root)
	return v
}

func helper(root *TreeNode) (bool, int, int) {
	if root == nil {
		return true, math.MaxInt64, math.MinInt64
	}

	left, lmin, lmax := helper(root.Left)
	right, rmin, rmax := helper(root.Right)

	if left && right && root.Val > lmax && root.Val < rmin {
		return true, min(lmin, root.Val), max(rmax, root.Val)
	}

	return false, 0, 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
// @lc code=end

