
/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
 // recursion
func inorderTraversal(root *TreeNode) []int {
  if root == nil {
		return []int{}
	}
	left := inorderTraversal(root.Left)
	left = append(left, root.Val)
	right := inorderTraversal(root.Right)
	left = append(left, right...)
	return left
}
// @lc code=end

