/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
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

 // DFS: using stack
func invertTree(root *TreeNode) *TreeNode {
	if root == nil { return nil }
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// swap left and right nodes
		tmp := top.Left
		top.Left = top.Right
		top.Right = tmp
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return root
}
// @lc code=end

