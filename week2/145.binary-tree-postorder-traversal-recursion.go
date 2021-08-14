/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
			return []int{}
		}
		left := postorderTraversal(root.Left)
		right := postorderTraversal(root.Right)
		left = append(left, right...)
		left = append(left, root.Val)
		return left
}
// @lc code=end

