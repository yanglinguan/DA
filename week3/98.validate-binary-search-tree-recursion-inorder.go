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

// inorder recursion
func isValidBST(root *TreeNode) bool {
	last := math.MinInt64
	return helper(root, &last)
}

func helper(root *TreeNode, last *int) bool {
	if root == nil {
		return true
	}

	left := helper(root.Left, last)
	if !left || root.Val <= *last {
		return false
	}
	*last = root.Val
	return helper(root.Right, last)
}
// @lc code=end

