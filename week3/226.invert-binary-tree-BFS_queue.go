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

 // BFS using queue
func invertTree(root *TreeNode) *TreeNode {
	if root == nil { return nil }
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		tmp := top.Left
		top.Left = top.Right
		top.Right = tmp

		if top.Left != nil {
			queue = append(queue, top.Left)
		}

		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return root
}
// @lc code=end

