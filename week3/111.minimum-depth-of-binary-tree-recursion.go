/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
 func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left != 0 && right != 0 {
		return 1 + min(left, right)
	} else if left == 0 {
		return 1 + right
	} else {
		return 1 + left
	}
}

func min(a, b int) int {
    if a > b { return b }
    return a
}
// @lc code=end

