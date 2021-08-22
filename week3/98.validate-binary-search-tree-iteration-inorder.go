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

// inorder iteration
func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	
	var prev *TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && cur.Val <= prev.Val {
			return false
		}
		prev = cur
		cur = cur.Right
	}

	return true
}
// @lc code=end

